package query

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/volatiletech/null"
)

//go:generate mockgen -source author.go -destination mock/mock_author.go
type Author interface {
	List(ctx context.Context) (output []*model.Author, err error)
	// GetByID returns the author with the given ID.
	GetByID(ctx context.Context, query AuthorGetQuery, orFail bool) (output *model.Author, err error)
}

type AuthorGetQuery struct {
	ID null.Int64
}
