package query

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/volatiletech/null"
)

//go:generate mockgen -source author.go -destination mock/mock_author.go
type Author interface {
	// List returns authors with filtering and pagination.
	List(ctx context.Context, filter AuthorListFilter) (output []*model.Author, err error)
	// GetByID returns the author with the given ID.
	GetByID(ctx context.Context, query AuthorGetQuery, orFail bool) (output *model.Author, err error)
}

type AuthorListFilter struct {
	Name null.String // 名前での部分一致検索
}

type AuthorGetQuery struct {
	ID null.Int64
}
