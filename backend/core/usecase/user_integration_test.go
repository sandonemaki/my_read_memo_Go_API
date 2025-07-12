package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/lib/pq"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	queryImpl "github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	repositoryImpl "github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

// テスト用のデータベース接続文字列
const testDSN = "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"

// setupTestDB はテスト用のデータベース接続を提供します
func setupTestDB(t *testing.T) *sql.DB {
	sqlDB, err := sql.Open("postgres", testDSN)
	if err != nil {
		t.Fatalf("データベース接続エラー: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("データベースPingエラー: %v", err)
	}
	return sqlDB
}

func TestCreateUser(t *testing.T) {
	// データベース接続
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	// vectors構造の導入
	vectors := map[string]struct {
		params   input.CreateUser
		expected *output.CreateUser
		wantErr  error
		options  cmp.Options
	}{
		"OK": {
			params: input.CreateUser{
				UID:         "integration_test_uid_001",
				DisplayName: "統合テストユーザー001",
			},
			expected: &output.CreateUser{
				User: &model.User{
					UID:         "integration_test_uid_001",
					DisplayName: "統合テストユーザー001",
					// Ulid, CreatedAt, UpdatedAtは動的値なので期待値には含めない
				},
			},
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.CreateUser{}, "User.Ulid", "User.CreatedAt", "User.UpdatedAt"),
			},
		},
		"UID空文字列エラー": {
			params: input.CreateUser{
				UID:         "", // 空文字列でバリデーションエラー
				DisplayName: "テストユーザー",
			},
			expected: nil, // エラーケースなので期待値なし
			wantErr:  fmt.Errorf("expected error"), // エラーを期待（見本通り）
			options:  cmp.Options{},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// 依存関係のセットアップ
			dbClient := db.NewClient(sqlDB)
			userQuery := queryImpl.NewUser(&dbClient)
			userRepo := repositoryImpl.NewUser(&dbClient)

			// ユースケースの作成
			var userUsecase User = NewUser(userQuery, userRepo)

			// テスト実行
			ctx := context.Background()
			result, err := userUsecase.Create(ctx, v.params)

			// エラー検証（エラー発生の有無のみチェック）
			if err != nil {
				if v.wantErr == nil {
					t.Fatalf("unexpected error: %v", err)
				}
				// エラーが期待されていて、実際にエラーが発生した場合は成功
				t.Logf("期待通りエラーが発生: %v", err)
			} else if v.wantErr != nil {
				t.Fatalf("expected an error, got none")
			}

			// 成功ケースの結果検証
			if err == nil {
				// go-cmpを使った構造体比較
				if !cmp.Equal(result, v.expected, v.options...) {
					t.Errorf("unexpected result: %s", cmp.Diff(v.expected, result, v.options...))
				}

				t.Logf("✅ ユーザー作成成功！ ULID: %s, UID: %s, DisplayName: %s", 
					result.User.Ulid, result.User.UID, result.User.DisplayName)
			}
		})
	}
}

func TestUserUsecase_GetCurrentUser(t *testing.T) {
	// データベース接続
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	// vectors構造の導入
	vectors := map[string]struct {
		setupUser   *input.CreateUser             // 事前作成用のユーザー（nilの場合はセットアップなし）
		params      input.GetCurrentUserDetail   // 取得時のパラメータ
		expected    *output.GetUser              // 期待する結果
		wantErr     error
		options     cmp.Options
	}{
		"OK": {
			setupUser: &input.CreateUser{
				UID:         "integration_get_test_uid_001",
				DisplayName: "取得テストユーザー001",
			},
			params: input.GetCurrentUserDetail{
				UID: "integration_get_test_uid_001",
			},
			expected: &output.GetUser{
				User: &model.User{
					UID:         "integration_get_test_uid_001",
					DisplayName: "取得テストユーザー001",
					// Ulid, CreatedAt, UpdatedAtは動的値なので期待値には含めない
				},
			},
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.GetUser{}, "User.Ulid", "User.CreatedAt", "User.UpdatedAt"),
			},
		},
		"UID空文字列エラー": {
			setupUser: nil, // セットアップなし
			params: input.GetCurrentUserDetail{
				UID: "", // 空文字列でバリデーションエラー
			},
			expected: nil,
			wantErr:  fmt.Errorf("expected error"),
			options:  cmp.Options{},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// 依存関係のセットアップ
			dbClient := db.NewClient(sqlDB)
			userQuery := queryImpl.NewUser(&dbClient)
			userRepo := repositoryImpl.NewUser(&dbClient)
			
			// ユースケースの作成
			var userUsecase User = NewUser(userQuery, userRepo)
			ctx := context.Background()

			// 事前データ作成（必要な場合のみ）
			if v.setupUser != nil {
				_, err := userUsecase.Create(ctx, *v.setupUser)
				if err != nil {
					t.Fatalf("事前ユーザー作成エラー: %v", err)
				}
			}

			// GetCurrentUserのテスト実行
			result, err := userUsecase.GetCurrentUser(ctx, v.params)

			// エラー検証（エラー発生の有無のみチェック）
			if err != nil {
				if v.wantErr == nil {
					t.Fatalf("unexpected error: %v", err)
				}
				// エラーが期待されていて、実際にエラーが発生した場合は成功
				t.Logf("期待通りエラーが発生: %v", err)
			} else if v.wantErr != nil {
				t.Fatalf("expected an error, got none")
			}

			// 成功ケースの結果検証
			if err == nil {
				// go-cmpを使った構造体比較
				if !cmp.Equal(result, v.expected, v.options...) {
					t.Errorf("unexpected result: %s", cmp.Diff(v.expected, result, v.options...))
				}

				t.Logf("✅ ユーザー取得成功！ ULID: %s, UID: %s, DisplayName: %s", 
					result.User.Ulid, result.User.UID, result.User.DisplayName)
			}
		})
	}
}

