package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

type MasterBook interface {
	Create(ctx context.Context, in input.CreateMasterBook) (out *output.CreateMasterBook, err error)
}
