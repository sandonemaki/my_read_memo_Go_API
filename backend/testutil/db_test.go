package testutil

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

// テスト用のデータベース接続文字列
const testDSN = "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"

func TestDatabaseConnection(t *testing.T) {
	// データベース接続テスト
	t.Run("データベース接続テスト", func(t *testing.T) {
		sqlDB, err := sql.Open("postgres", testDSN)
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
		sqlDB, err := sql.Open("postgres", testDSN)
		if err != nil {
			t.Fatalf("データベース接続エラー: %v", err)
		}
		defer sqlDB.Close()

		// db.Clientを作成
		dbClient := db.NewClient(sqlDB)

		// dbClientが正しく初期化されているか確認
		if dbClient == (db.Client{}) { // 空のClientと比較
			t.Fatal("db.Clientの作成に失敗しました")
		}
		t.Log("✅ db.Client作成成功")
	})
}
