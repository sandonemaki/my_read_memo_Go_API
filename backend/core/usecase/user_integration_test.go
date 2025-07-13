package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"os"
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

// getTestDSN はテスト用のデータベース接続文字列を取得します（セキュリティ対応）
func getTestDSN() string {
	if dsn := os.Getenv("TEST_DATABASE_URL"); dsn != "" {
		return dsn
	}
	// フォールバック（ローカル開発用）
	return "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"
}

// setupTestDB はテスト用のデータベース接続を提供します
func setupTestDB(t *testing.T) *sql.DB {
	sqlDB, err := sql.Open("postgres", getTestDSN())
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
	
	// テスト終了後のクリーンアップ
	defer cleanupTestData(t, sqlDB)

	// vectors構造の導入
	vectors := map[string]struct {
		params   input.CreateUser
		expected *output.CreateUser
		wantErr  error
		options  cmp.Options
	}{
		"OK": {
			params: input.CreateUser{
				Ulid:        "01HWEB0000000000000000001", // 新しい固定ULID（重複回避）
				UID:         "integration_test_uid_fixed_ulid_001", // 新しい固定UID（重複回避）
				DisplayName: "固定ULID統合テストユーザー001",
			},
			expected: &output.CreateUser{
				User: &model.User{
					Ulid:        "01HWEB0000000000000000001", // 新しい固定ULID
					UID:         "integration_test_uid_fixed_ulid_001",
					DisplayName: "固定ULID統合テストユーザー001",
					// CreatedAt, UpdatedAtは動的値なので期待値には含めない
				},
			},
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.CreateUser{}, "User.CreatedAt", "User.UpdatedAt"), // Ulidを除外から削除
			},
		},
		"UID空文字列エラー": {
			params: input.CreateUser{
				UID:         "", // 空文字列でバリデーションエラー
				DisplayName: "テストユーザー",
			},
			expected: nil, // エラーケースなので期待値なし
			wantErr:  fmt.Errorf("validation failed"), // バリデーションエラーを期待
			options:  cmp.Options{},
		},
	}

	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// 依存関係のセットアップ（元の実装を維持）
			dbClient := db.NewClient(sqlDB) // use the *sql.DB instance, not tx
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

// cleanupTestData はテスト用データを削除します
func cleanupTestData(t *testing.T, sqlDB *sql.DB) {
	_, err := sqlDB.Exec("DELETE FROM users WHERE uid LIKE '%test%' OR uid LIKE '%integration%'")
	if err != nil {
		t.Logf("クリーンアップエラー（無視されます）: %v", err)
	} else {
		t.Log("✅ テストデータクリーンアップ完了")
	}
}

// setupTestUsers はテスト用のユーザーデータを事前作成します
func setupTestUsers(t *testing.T, sqlDB *sql.DB) {
	dbClient := db.NewClient(sqlDB)
	userQuery := queryImpl.NewUser(&dbClient)
	userRepo := repositoryImpl.NewUser(&dbClient)
	userUsecase := NewUser(userQuery, userRepo)
	ctx := context.Background()

	// テスト用ユーザーを事前作成
	testUsers := []input.CreateUser{
		{
			UID:         "pre_created_test_uid_001",
			DisplayName: "事前作成ユーザー001",
		},
		{
			UID:         "pre_created_test_uid_002", 
			DisplayName: "事前作成ユーザー002",
		},
	}

	for _, user := range testUsers {
		_, err := userUsecase.Create(ctx, user)
		if err != nil {
			t.Fatalf("事前ユーザー作成エラー: %v", err)
		}
	}
}

func TestUserUsecase_GetCurrentUser(t *testing.T) {
	// データベース接続
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	// 事前にテストデータを作成
	setupTestUsers(t, sqlDB)

	// vectors構造（Getのテストのみに集中）
	vectors := map[string]struct {
		params   input.GetCurrentUserDetail   // 取得時のパラメータ
		expected *output.GetUser              // 期待する結果
		wantErr  error
		options  cmp.Options
	}{
		"OK": {
			params: input.GetCurrentUserDetail{
				UID: "pre_created_test_uid_001", // 事前作成済みのUID
			},
			expected: &output.GetUser{
				User: &model.User{
					UID:         "pre_created_test_uid_001",
					DisplayName: "事前作成ユーザー001",
					// Ulid, CreatedAt, UpdatedAtは動的値なので期待値には含めない
				},
			},
			wantErr: nil,
			options: cmp.Options{
				cmpopts.IgnoreFields(output.GetUser{}, "User.Ulid", "User.CreatedAt", "User.UpdatedAt"),
			},
		},
		"存在しないUIDエラー": {
			params: input.GetCurrentUserDetail{
				UID: "non_existent_uid", // 存在しないUID
			},
			expected: nil,
			wantErr:  fmt.Errorf("user not found"), // More specific error expectation
			options:  cmp.Options{},
		},
		"UID空文字列エラー": {
			params: input.GetCurrentUserDetail{
				UID: "", // 空文字列でバリデーションエラー
			},
			expected: nil,
			wantErr:  fmt.Errorf("validation failed"), // Validation error expectation
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

			// GetCurrentUserのテスト実行（事前データありき）
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

