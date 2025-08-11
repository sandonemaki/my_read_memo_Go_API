package query

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

type masterBook struct {
	dbClient *db.Client
}

func NewMMasterBook(dbClient *db.Client) query.MasterBook {
	return &masterBook{dbClient}
}

func (m *masterBook) Search(ctx context.Context, query query.MasterBookSearchQuery) (output []*model.MasterBook, err error) {
	mods := []bob.Mod[*dialect.SelectQuery]{}

	if query.AuthorName.Valid {
		// 1. 著者名検索の場合、authorsテーブルとJOIN
		mods = append(mods, dbmodels.SelectJoins.MasterBooks.InnerJoin.Author)
		// authors.nameで検索（部分一致、大文字小文字区別なし）
		mods = append(mods, dbmodels.SelectWhere.Authors.Name.ILike("%"+query.AuthorName.String+"%"))
	}

	// 出版社名検索
	if query.PublisherName.Valid {
		mods = append(mods, dbmodels.SelectJoins.MasterBooks.InnerJoin.Publisher)
		mods = append(mods, dbmodels.SelectWhere.Publishers.Name.ILike("%"+query.PublisherName.String+"%"))
	}

	// タイトル検索
	if query.Title.Valid {
		mods = append(mods, dbmodels.SelectWhere.MasterBooks.Title.ILike("%"+query.Title.String+"%"))
	}

	dbMasterBooks, err := dbmodels.MasterBooks.Query(mods...).All(ctx, m.dbClient.Get(ctx))
	if err != nil {
		return nil, err
	}

	masterBooks := make([]*model.MasterBook, len(dbMasterBooks))
	for i, dbMasterBook := range dbMasterBooks {
		masterBooks[i] = (*model.MasterBook)(dbMasterBook)
	}
	return masterBooks, nil
}
