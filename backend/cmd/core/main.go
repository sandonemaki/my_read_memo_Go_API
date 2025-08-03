package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/injector"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func main() {
	c := config.Prepare()

	r := chi.NewRouter()
	// wire呼び出し: 必要な部品（ログ、DB）を自動で組み立て
	core, err := injector.InitializeCoreHandler(c.Logger, c.Postgres, "core")
	if err != nil {
		slog.Error("failed to initialize core handler", slog.Any("err", err))
		os.Exit(1)
	}

	// Wireで作成されたloggerを使用（重複解消）
	core.Logger.Info("starting server", "port", c.HTTP.Port)

	// CORS設定
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: c.HTTP.Cors,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true,
	}))

	serverCtx := context.Background()
	//  認証チェック
	r.Use(core.GetAuthMiddleware(serverCtx, core.Logger))

	// OpenAPI仕様ファイルに基づいて、自動的にルート（URL pattern）を登録
	strictHandler := oapi.NewStrictHandler(core, nil)
	oapi.HandlerFromMux(strictHandler, r)

	server := http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf(":%d", c.HTTP.Port),
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       2 * time.Second,
	}

	// 1. サーバーをgorutineで起動
	go func() {
		core.Logger.Info("server is running", slog.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			core.Logger.Error("server closed", slog.Any("err", err))
			os.Exit(1)
		}
	}()
	// 2. シグナル受診チャンネル作成
	quit := make(chan os.Signal, 1)

	// 3. Ctrl+CやSIGTERMを受け取るための設定
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 4. シグナルを待機（ここでブロック）
	<-quit
	core.Logger.Info("shutting down server...")

	// 5. サーバーをGraceful Shutdown開始
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 6. 処理中のリクエストが完了するまで待つ
	if err := server.Shutdown(ctx); err != nil {
		core.Logger.Error("server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}
	core.Logger.Info("server exited")
}
