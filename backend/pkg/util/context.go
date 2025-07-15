package util

import (
	"context"
	"database/sql"
)

// https://deeeet.com/writing/2017/02/23/go-context-value/
type contextKey string

var (
	dbTxContextKey contextKey = "dbTx"
)

// 型アサーション（Type Assertion）
// tx := value.(*sql.Tx)
// - interface{}から*sql.Tx型に変換
// - 危険： 型が違うとパニック（プログラム停止）が発生
func GetDBTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(dbTxContextKey).(*sql.Tx); ok {
		return tx
	}
	return nil
}
