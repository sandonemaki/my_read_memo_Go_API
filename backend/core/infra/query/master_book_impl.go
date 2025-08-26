package query

import (
	"context"
	"errors"

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

func NewMasterBook(dbClient *db.Client) query.MasterBook {
	return &masterBook{dbClient}
}

func (m *masterBook) Search(ctx context.Context, query query.MasterBookSearchQuery) (output []*model.MasterBook, err error) {
	mods := []bob.Mod[*dialect.SelectQuery]{}

	if query.AuthorName.Valid {
		// 1. 著者名検索の場合、authorsテーブルとJOIN
		mods = append(mods, dbmodels.SelectJoins.MasterBooks.InnerJoin.Author)
		// SQLインジェクション対策
		pattern := "%" + db.EscapeLikePattern(query.AuthorName.String) + "%"

		// authors.nameで検索（部分一致、大文字小文字区別なし）
		mods = append(mods, dbmodels.SelectWhere.Authors.Name.ILike(pattern))
	}

	// 出版社名検索
	if query.PublisherName.Valid {
		mods = append(mods, dbmodels.SelectJoins.MasterBooks.InnerJoin.Publisher)
		pattern := "%" + db.EscapeLikePattern(query.PublisherName.String) + "%"

		mods = append(mods, dbmodels.SelectWhere.Publishers.Name.ILike(pattern))
	}

	// タイトル検索
	if query.Title.Valid {
		pattern := "%" + db.EscapeLikePattern(query.Title.String) + "%"

		mods = append(mods, dbmodels.SelectWhere.MasterBooks.Title.ILike(pattern))
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

func (m *masterBook) GetByID(ctx context.Context, query query.MasterBookGetQuery, orFail bool) (*model.MasterBook, error) {
	// TODO:実装
	return nil, errors.New("masterBook.GetByID not implemented")
}

func (m *masterBook) List(ctx context.Context) ([]*model.MasterBook, error) {
	// TODO:実装
	return nil, errors.New("masterBook.GetByID not implemented")
}
