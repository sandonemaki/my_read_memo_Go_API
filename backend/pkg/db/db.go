package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/stephenafamo/scan"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/util"
)

// Client :
type Client struct {
	dbClient *sql.DB
}

// NewDB initialize database
func NewDB(sqlDB *sql.DB) *Client {
	time.Local = time.FixedZone("JST", 9*60*60)
	client := NewClient(sqlDB)
	return &client
}

// NewClient :
func NewClient(db *sql.DB) Client {
	// dbはポインタ型ではなく、値型であるため、直接値を返す
	return Client{db}
}

// SQLHandler : BOB版のExecutorインターフェース
type SQLHandler interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (scan.Rows, error) // Bob ORMが期待するscan.Rows
}

// Get : コンテキストからトランザクションまたはDBクライアントを取得
func (c Client) Get(ctx context.Context) SQLHandler {
	if tx := util.GetDBTx(ctx); tx != nil {
		return &TxWrapper{tx: tx} // トランザクション用のラッパー
	}
	return c // 通常のDB接続の場合
}

// ExecContextメソッドは、SQLデータベースに対して更新系のクエリを実行するためのメソッド
// 主に書き込み系の操作で使用
func (db Client) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.dbClient.ExecContext(ctx, query, args...)
}

// QueryContextメソッドは、SQLデータベースに対してSELECT系のクエリを実行するためのメソッド
// 主に読み取り系の操作で使用
func (db Client) QueryContext(ctx context.Context, query string, args ...any) (scan.Rows, error) {
	rows, err := db.dbClient.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// TxWrapper : トランザクション用のラッパー
type TxWrapper struct {
	tx *sql.Tx
}

func (w *TxWrapper) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return w.tx.ExecContext(ctx, query, args...)
}

func (w *TxWrapper) QueryContext(ctx context.Context, query string, args ...any) (scan.Rows, error) {
	rows, err := w.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
