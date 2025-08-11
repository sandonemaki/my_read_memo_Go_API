package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/auth/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/auth/infra/firebase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

// コンパイル時に Core 構造体が StrictServerInterface を完全に実装しているかチェック
// Core構造体のポインタ型にnilをキャスト
var _ oapi.StrictServerInterface = (*Core)(nil)

type Core struct {
	Unimplemented                 // 手動で作成したUnimplementedを埋め込み
	Logger           *slog.Logger // 大文字に変更して外部アクセス可能
	firebaseAuthGlue firebase.FirebaseAuthGlue
	userUsecase      usecase.User
	publisherUsecase usecase.Publisher
	authorUsecase    usecase.Author
}

func NewCore(
	logger *slog.Logger,
	firebaseAuthGlue firebase.FirebaseAuthGlue,
	userUsecase usecase.User,
	publisherUsecase usecase.Publisher,
	authorUsecase usecase.Author,
) *Core {
	return &Core{
		Logger:           logger, // フィールド名も大文字に変更
		firebaseAuthGlue: firebaseAuthGlue,
		userUsecase:      userUsecase,
		publisherUsecase: publisherUsecase,
		authorUsecase:    authorUsecase,
	}
}

func (c *Core) GetAuthMiddleware(ctx context.Context, logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Authentication logic here
			ctx := r.Context()

			// TODO: healthという認証が必要ないパスを作成する

			var buf bytes.Buffer
			tee := io.TeeReader(r.Body, &buf)
			r.Body = io.NopCloser(&buf) // r.Bodyを復元

			body, _ := io.ReadAll(tee)
			logger.InfoContext(r.Context(), "dump request",
				"method", r.Method,
				"url", r.URL.String(),
				"header", r.Header,
				"body", string(body))
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

// トランザクション処理のラッパー関数（前処理・後処理を提供）
func WithTx(ctx context.Context, logger *slog.Logger, txfunc func(ctx context.Context) error) (err error) {
	defer func() {
		// panicが発生した場合のリカバリ
		if r := recover(); r != nil {
			// panicをerrorに変換
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("panic recovered: %v", r)
			}
		}
		// エラーがあればログの出力
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("error: %+v", err))
		}
	}()
	// 渡されたトランザクション処理を実行
	return txfunc(ctx)
}

// NOTE: 新規ユーザーの場合2回のAPI呼び出し（以下）を1回にするため
func (c *Core) LoginOrSignup(ctx context.Context, idToken string) (*model.Credential, error) {
	// 1. Firebase認証
	credential, err := c.firebaseAuthGlue.CheckLoginJWT(ctx, idToken)
	if err != nil {
		return nil, err
	}

	// 2. 事前チェック: ユーザー存在確認（効率的）
	_, err = c.userUsecase.GetCurrentUser(ctx, input.GetCurrentUserDetail{
		UID: credential.UID,
	})

	if err != nil {
		// 3. 存在しない場合のみ作成（必要な時だけ）
		userInput := input.NewCreateUser(credential.UID, credential.GetSafeDisplayName())
		_, err = c.userUsecase.Create(ctx, userInput)
		if err != nil {
			return nil, err
		}
	}
	// 4. 既存ユーザーの場合は何もしない
	return credential, nil
}
