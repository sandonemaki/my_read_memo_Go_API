package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

type User interface {
	Create(ctx context.Context, user *model.User) (err error)
	Update(ctx context.Context, user *model.User) (userUID string, err error)
	Delete(ctx context.Context, userUID string) (err error)
}
