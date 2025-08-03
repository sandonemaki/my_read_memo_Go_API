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

type author struct {
	dbClient *db.Client
}

func NewAuthor(dbClient *db.Client) query.Author {
	return &author{dbClient}
}

func (r *author) GetByID(ctx context.Context, query query.AuthorGetQuery, orFail bool) (output *model.Author, err error) {
	mods := []bob.Mod[*dialect.SelectQuery]{}

	// GetByIDメソッドなので、IDのみで検索
	if query.ID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Authors.ID.EQ(query.ID.Int64))
	}

	// Bob ORMでクエリを構築して実行
	dbAuthor, err := dbmodels.Authors.Query(mods...).One(ctx, r.dbClient)

	// エラーハンドリング
	if err != nil {
		// データが見つからない場合のエラーかチェック
		if errors.Is(err, sql.ErrNoRows) {
			// orFailパラメータで動作を制御
			if orFail {
				// orFail=true: データが必須なのでエラーを返す
				return nil, err
			}
			// orFail=false: データがなくてもOK、nilを返す（エラーではない）
			return nil, nil
		}
		// その他のDBエラー（接続エラー等）はそのまま返す
		return nil, err
	}

	// 型変換: dbmodels.Author → model.Author
	// 現在は型エイリアスなので直接キャスト可能だが、
	// 将来OpenAPI生成型を使う場合は明示的なフィールドマッピングが必要
	return (*model.Author)(dbAuthor), nil
}

func (r *author) List(ctx context.Context, filter query.AuthorListFilter) (output []*model.Author, err error) {
	// Step 1: Modifierの配列を初期化（WHERE句などの条件を格納）
	mods := []bob.Mod[*dialect.SelectQuery]{}

	// Step 2: フィルター条件があれば追加（今は空のまま）
	// TODO: filter.Nameがある場合の処理を追加

	// Step 3: 全件取得のクエリを実行
	// .All() は複数件取得するメソッド
	dbAuthors, err := dbmodels.Authors.Query(mods...).All(ctx, r.dbClient)
	if err != nil {
		return nil, err
	}

	// Step 4: 型変換: []*dbmodels.Author → []*model.Author
	// スライスは要素ごとに変換する必要がある
	authors := make([]*model.Author, len(dbAuthors)) // 結果用のスライスを作成
	for i, dbAuthor := range dbAuthors {
		// 各要素を型変換
		authors[i] = (*model.Author)(dbAuthor)
	}

	return authors, nil
}
