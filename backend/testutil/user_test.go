package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

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

func TestUserCreation(t *testing.T) {
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	t.Run("ユーザー作成テスト", func(t *testing.T) {
		// db.Clientを作成
		dbClient := db.NewClient(sqlDB)

		// リポジトリを作成
		userRepo := repository.NewUser(&dbClient)

		// テストユーザーを作成
		ctx := context.Background()
		timestamp := time.Now().Unix()
		testUser := &model.User{
			Ulid:        fmt.Sprintf("test_%d", timestamp),
			UID:         fmt.Sprintf("uid_%d", timestamp),
			DisplayName: fmt.Sprintf("テストユーザー_%d", timestamp),
			DeletedAt:   sql.Null[time.Time]{},
		}

		t.Logf("作成予定のユーザー: ULID=%s, UID=%s, DisplayName=%s",
			testUser.Ulid, testUser.UID, testUser.DisplayName)

		// Create操作をテスト
		err := userRepo.Create(ctx, testUser)
		if err != nil {
			t.Fatalf("ユーザー作成エラー: %v", err)
		}

		t.Log("✅ ユーザー作成成功！")
		t.Logf("🎉 作成されたユーザー: ULID=%s, UID=%s, DisplayName=%s",
			testUser.Ulid, testUser.UID, testUser.DisplayName)
	})
}
