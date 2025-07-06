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

type user struct {
	dbClient *db.Client
}

func NewUser(dbClient *db.Client) query.User {
	return &user{dbClient}
}

func (r *user) GetByULID(ctx context.Context, query query.UserGetQuery) (user *model.User, err error) {

	mods := []bob.Mod[*dialect.SelectQuery]{}

	if query.ULID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Users.Ulid.EQ(query.ULID.String))
	}
	if query.UID.Valid {
		mods = append(mods, dbmodels.SelectWhere.Users.UID.EQ(query.UID.String))
	}

	dbUser, err := dbmodels.Users.Query(mods...).One(ctx, r.dbClient)
	if err != nil {
		return nil, err
	}

	return (*model.User)(dbUser), nil
}
