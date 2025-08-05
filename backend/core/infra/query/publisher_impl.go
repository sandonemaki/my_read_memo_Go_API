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

type publisher struct {
	dbClient *db.Client
}

func NewPublisher(dbClient *db.Client) query.Publisher {
	return &publisher{dbClient}
}

func (p *publisher) GetByID(ctx context.Context, query query.PublisherGetQuery, orFail bool) (output *model.Publisher, err error) {
	mods := []bob.Mod[*dialect.SelectQuery]{}

	if query.ID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Publishers.ID.EQ(query.ID.Int64))
	}

	dbPublisher, err := dbmodels.Publishers.Query(mods...).One(ctx, p.dbClient.Get(ctx))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if orFail {
				return nil, err
			}
			return nil, nil
		}
		return nil, err
	}
	return (*model.Publisher)(dbPublisher), nil
}

func (p *publisher) List(ctx context.Context, filter query.PublisherListFilter) (output []*model.Publisher, err error) {

	mods := []bob.Mod[*dialect.SelectQuery]{}

	dbPublishers, err := dbmodels.Publishers.Query(mods...).All(ctx, p.dbClient.Get(ctx))
	if err != nil {
		return nil, err
	}
	publishers := make([]*model.Publisher, len(dbPublishers))
	for i, dbPublisher := range dbPublishers {
		publishers[i] = (*model.Publisher)(dbPublisher)
	}
	return publishers, nil

}
