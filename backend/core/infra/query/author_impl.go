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

func (r *author) GetByID(ctx context.Context, query query.AuthorGetQuery, orFail bool, options ...db.Query) (output *model.Author, err error) {
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
	return (*model.Author)(dbAuthor), nil
}

// List returns authors with filtering and pagination.
func (r *author) List(ctx context.Context, filter query.AuthorListFilter, options ...db.Query) (output []*model.Author, err error) {
	// TODO: 一緒に実装しましょう
	return nil, nil
}