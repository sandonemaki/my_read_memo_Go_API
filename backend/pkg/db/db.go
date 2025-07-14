package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/stephenafamo/scan"
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

// SQLHandlerはbobの場合必要ない

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
