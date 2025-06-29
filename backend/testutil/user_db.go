package testutil

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

func TestDatabaseConnection(t *testing.T) {
	// データベース接続テスト
	t.Run("データベース接続テスト", func(t *testing.T) {
		// データベース接続
		dsn := "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"
		sqlDB, err := sql.Open("postgres", dsn)
		if err != nil {
			t.Fatalf("データベース接続エラー: %v", err)
		}
		defer sqlDB.Close()

		// 接続テスト
		if err := sqlDB.Ping(); err != nil {
			t.Fatalf("データベースPingエラー: %v", err)
		}
		t.Log("✅ データベース接続成功！")
	})

	t.Run("db.Client作成テスト", func(t *testing.T) {
		// データベース接続
		dsn := "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"
		sqlDB, err := sql.Open("postgres", dsn)
		if err != nil {
			t.Fatalf("データベース接続エラー: %v", err)
		}
		defer sqlDB.Close()

		// db.Clientを作成
		dbClient := db.NewClient(sqlDB)
		if dbClient.ExecContext == nil {
			t.Fatal("db.Clientの作成に失敗しました")
		}
		t.Log("✅ db.Client作成成功")
	})
}
