package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

func testMain() {
	// データベース接続
	dsn := "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}
	defer sqlDB.Close()

	// 接続テスト
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("データベースPingエラー:", err)
	}
	fmt.Println("✓ データベース接続成功！")

	// db.Clientを作成
	dbClient := db.NewClient(sqlDB)
	fmt.Println("✓ db.Client作成成功")

	// リポジトリを作成
	userRepo := repository.NewUser(&dbClient)
	fmt.Println("✓ リポジトリ作成成功！")

	// テストユーザーを作成
	ctx := context.Background()
	testUser := &model.User{
		Ulid:      fmt.Sprintf("test_%d", time.Now().Unix()),
		UID:       fmt.Sprintf("uid_%d", time.Now().Unix()),
		Nickname:  "テストユーザー",
		DeletedAt: sql.Null[time.Time]{}, // NULL値
	}

	// Create操作をテスト
	err = userRepo.Create(ctx, testUser)
	if err != nil {
		log.Fatal("ユーザー作成エラー:", err)
	}

	fmt.Println("✓ ユーザー作成成功！")
	fmt.Printf("作成されたユーザー: ULID=%s, UID=%s, Nickname=%s\n",
		testUser.Ulid, testUser.UID, testUser.Nickname)
}
