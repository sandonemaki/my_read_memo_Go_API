package query

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/volatiletech/null"
)

//go:generate mockgen -source publisher.go -destination mock/mock_publisher.go
type Publisher interface {
	List(ctx context.Context) (output []*model.Publisher, err error)
	// GetByID returns the publisher with the given ID.
	GetByID(ctx context.Context, query PublisherGetQuery, orFail bool) (output *model.Publisher, err error)
}

type PublisherGetQuery struct {
	ID null.Int64
}
