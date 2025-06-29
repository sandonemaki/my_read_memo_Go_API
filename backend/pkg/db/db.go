package db

import (
	"context"
	"database/sql"

	"github.com/stephenafamo/scan"
)

// Client :
type Client struct {
	dbClient *sql.DB
}

// NewClient :
func NewClient(db *sql.DB) Client {
	// dbはポインタ型ではなく、値型であるため、直接値を返す
	return Client{db}
}

// SQLHandler :
// type SQLHandler interface {
// 	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
// 	QueryContext(ctx context.Context, query string, args ...any) (scan.Rows, error)
// }

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
