package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/volatiletech/null"
)

type User interface {
	Get(ctx context.Context, query UserGetQuery) (*model.User, error)
	Create(ctx context.Context, user *model.User) (err error)
}

type UserGetQuery struct {
	ULID null.String // 検索条件にULIDを含めるか？
	UID  null.String // 検索条件にUIDを含めるか？
}
