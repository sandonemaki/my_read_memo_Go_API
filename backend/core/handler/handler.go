package handler

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/firebase"
)

type Core struct {
	logger      *slog.Logger
	firebase    firebase.Glue
	userUsecase usecase.User
}

func NewCore(
	logger *slog.Logger,
	firebase firebase.Glue,
	userUsecase usecase.User,
) *Core {
	return &Core{
		logger:      logger,
		firebase:    firebase,
		userUsecase: userUsecase,
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
			logger.InfoContext(r.Context(), "dump request", "method", r.Method, "url", r.URL.String(), "header", r.Header, "body", string(body))
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
