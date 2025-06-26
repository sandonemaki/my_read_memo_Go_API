package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

type User interface {
	Create(ctx context.Context, user *model.User) (err error)
}
