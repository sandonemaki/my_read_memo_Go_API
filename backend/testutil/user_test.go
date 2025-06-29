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

// テスト用のデータベース接続を提供するヘルパー関数
func setupTestDB(t *testing.T) *sql.DB {
	dsn := "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("データベース接続エラー: %v", err)
	}
	return sqlDB
}

func TestUserCreation(t *testing.T) {
	// データベース接続をセットアップ
	sqlDB := setupTestDB(t)
	defer sqlDB.Close()

	t.Run("ユーザー作成テスト", func(t *testing.T) {
		// db.Clientを作成
		dbClient := db.NewClient(sqlDB)

		// リポジトリを作成
		userRepo := repository.NewUser(&dbClient)
		t.Log("✅ リポジトリ作成成功！")

		// テストユーザーを作成
		ctx := context.Background()
		timestamp := time.Now().Unix()
		testUser := &model.User{
			Ulid:      fmt.Sprintf("test_%d", timestamp),
			UID:       fmt.Sprintf("uid_%d", timestamp),
			Nickname:  fmt.Sprintf("テストユーザー_%d", timestamp),
			DeletedAt: sql.Null[time.Time]{}, // NULL値
		}

		t.Logf("作成予定のユーザー: ULID=%s, UID=%s, Nickname=%s",
			testUser.Ulid, testUser.UID, testUser.Nickname)

		// Create操作をテスト
		err := userRepo.Create(ctx, testUser)
		if err != nil {
			t.Fatalf("ユーザー作成エラー: %v", err)
		}

		t.Log("✅ ユーザー作成成功！")
		t.Logf("🎉 作成されたユーザー: ULID=%s, UID=%s, Nickname=%s",
			testUser.Ulid, testUser.UID, testUser.Nickname)
	})
}
