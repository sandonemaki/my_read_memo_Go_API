package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

type Author interface {
	Create(ctx context.Context, in input.CreateAuthor) (out *output.CreateAuthor, err error)
	GetByID(ctx context.Context, authorID int64) (out *output.GetAuthor, err error)
	List(ctx context.Context) (out *output.ListAuthors, err error)
	SearchByName(ctx context.Context, name string) (out *output.ListAuthors, err error)
}
