package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

type MasterBook interface {
	Create(ctx context.Context, book *model.MasterBook) (bookID int64, err error)
	Update(ctx context.Context, book *model.MasterBook) (bookID int64, err error)
}
