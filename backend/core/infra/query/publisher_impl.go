package query

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

// publisher は出版社クエリの実装構造体
// private: パッケージ外から直接インスタンス化を防ぐ
type publisher struct {
	// dbClient はデータベース接続を管理
	// private: 外部から直接アクセスさせない（カプセル化）
	dbClient *db.Client
}

// NewPublisher は出版社クエリのインスタンスを生成する
// Public: 外部パッケージから呼び出し可能（コンストラクタ）
// 引数: dbClient - データベース接続クライアント
// 戻り値: query.Publisher インターフェース（実装の詳細を隠蔽）
func NewPublisher(dbClient *db.Client) query.Publisher {
	return &publisher{dbClient}
}

// GetByID は指定されたIDの出版社を取得する
// Public（インターフェース実装）: query.Publisherインターフェースの実装
// 引数:
//   - ctx: コンテキスト（タイムアウトやキャンセル制御）
//   - query: 検索条件（IDを含む）
//   - orFail: trueの場合、見つからない時にエラーを返す
//
// 戻り値:
//   - output: 見つかった出版社（見つからない場合はnilまたはエラー）
//   - err: エラー情報
func (p *publisher) GetByID(ctx context.Context, query query.PublisherGetQuery, orFail bool) (output *model.Publisher, err error) {
	// Bob ORMのクエリビルダー用のモディファイア配列
	// private（ローカル変数）: メソッド内でのみ使用
	mods := []bob.Mod[*dialect.SelectQuery]{}

	// IDが有効な場合、WHERE条件を追加
	// null.Int64のValidフィールドで有効性を確認
	if query.ID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Publishers.ID.EQ(query.ID.Int64))
	}

	// データベースから1件取得
	// Bob ORMのQueryメソッドでSQLを生成・実行
	dbPublisher, err := dbmodels.Publishers.Query(mods...).One(ctx, p.dbClient.Get(ctx))

	// エラーハンドリング
	if err != nil {
		// レコードが見つからない場合の処理
		if errors.Is(err, sql.ErrNoRows) {
			// orFailがtrueの場合はエラーを返す
			if orFail {
				return nil, err
			}
			// orFailがfalseの場合はnilを返す（エラーとしない）
			return nil, nil
		}
		// その他のエラーはそのまま返す
		return nil, err
	}
	// dbmodels.Publisher型をdomain層のmodel.Publisher型にキャスト
	// 型変換により、infra層の詳細をdomain層の型で隠蔽
	return (*model.Publisher)(dbPublisher), nil
}

// List は出版社の一覧を取得する
// Public（インターフェース実装）: query.Publisherインターフェースの実装
// 引数:
//   - ctx: コンテキスト
//   - filter: フィルター条件（現在は未使用だが将来の拡張用）
//
// 戻り値:
//   - output: 出版社のスライス
//   - err: エラー情報
func (p *publisher) List(ctx context.Context) (output []*model.Publisher, err error) {
	// クエリモディファイア（現在は条件なしで全件取得）
	// private（ローカル変数）: メソッド内でのみ使用
	mods := []bob.Mod[*dialect.SelectQuery]{}

	// データベースから全件取得
	// Bob ORMのAllメソッドで複数レコードを取得
	dbPublishers, err := dbmodels.Publishers.Query(mods...).All(ctx, p.dbClient.Get(ctx))
	if err != nil {
		return nil, err
	}
	// dbmodels.Publisher型の配列をmodel.Publisher型の配列に変換
	// private（ローカル変数）: 戻り値の準備用
	publishers := make([]*model.Publisher, len(dbPublishers))
	for i, dbPublisher := range dbPublishers {
		// 各要素を型キャストして格納
		publishers[i] = (*model.Publisher)(dbPublisher)
	}
	return publishers, nil

}
