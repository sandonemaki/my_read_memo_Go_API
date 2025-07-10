package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	queryImpl "github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	repositoryImpl "github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
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

func TestUserUsecase_Create(t *testing.T) {
	// データベース接続
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	t.Run("ユーザー作成テスト", func(t *testing.T) {
		// 依存関係のセットアップ
		dbClient := db.NewClient(sqlDB)
		userQuery := queryImpl.NewUser(&dbClient)
		userRepo := repositoryImpl.NewUser(&dbClient)

		// ユースケースの作成
		userUsecase := NewUser(userQuery, userRepo)

		// テスト用入力データ
		timestamp := time.Now().Unix()
		inputData := input.CreateUser{
			UID:         fmt.Sprintf("test_uid_%d", timestamp),
			DisplayName: fmt.Sprintf("テストユーザー_%d", timestamp),
		}

		// テスト実行
		ctx := context.Background()
		result, err := userUsecase.Create(ctx, inputData)

		// 結果検証
		if err != nil {
			t.Fatalf("ユーザー作成エラー: %v", err)
		}

		if result == nil {
			t.Fatal("結果がnilです")
		}

		if result.User == nil {
			t.Fatal("ユーザーがnilです")
		}

		if result.User.UID != inputData.UID {
			t.Errorf("UID が一致しません。期待値: %s, 実際: %s", inputData.UID, result.User.UID)
		}

		if result.User.DisplayName != inputData.DisplayName {
			t.Errorf("DisplayName が一致しません。期待値: %s, 実際: %s", inputData.DisplayName, result.User.DisplayName)
		}

		t.Logf("✅ ユーザー作成成功！ ULID: %s, UID: %s, DisplayName: %s", 
			result.User.Ulid, result.User.UID, result.User.DisplayName)
	})
}

func TestUserUsecase_GetMe(t *testing.T) {
	// データベース接続
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	t.Run("ユーザー取得テスト", func(t *testing.T) {
		// 依存関係のセットアップ
		dbClient := db.NewClient(sqlDB)
		userQuery := queryImpl.NewUser(&dbClient)
		userRepo := repositoryImpl.NewUser(&dbClient)
		
		// ユースケースの作成
		userUsecase := NewUser(userQuery, userRepo)

		// 事前にテストユーザーを作成
		timestamp := time.Now().Unix()
		createInput := input.CreateUser{
			UID:         fmt.Sprintf("test_uid_%d", timestamp),
			DisplayName: fmt.Sprintf("テストユーザー_%d", timestamp),
		}

		ctx := context.Background()
		createResult, err := userUsecase.Create(ctx, createInput)
		if err != nil {
			t.Fatalf("事前ユーザー作成エラー: %v", err)
		}

		// GetMeのテスト
		getMeInput := input.GetCurrentUserDetail{
			UID: createResult.User.UID,
		}

		result, err := userUsecase.GetMe(ctx, getMeInput)
		if err != nil {
			t.Fatalf("ユーザー取得エラー: %v", err)
		}

		if result == nil {
			t.Fatal("結果がnilです")
		}

		if result.User == nil {
			t.Fatal("ユーザーがnilです")
		}

		if result.User.UID != createResult.User.UID {
			t.Errorf("UID が一致しません。期待値: %s, 実際: %s", createResult.User.UID, result.User.UID)
		}

		t.Logf("✅ ユーザー取得成功！ ULID: %s, UID: %s, DisplayName: %s", 
			result.User.Ulid, result.User.UID, result.User.DisplayName)
	})
}

