package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

//go:generate mockgen -source publisher.go -destination mock/mock_publisher.go
type Publisher interface {
	Create(ctx context.Context, in input.CreatePublisher) (out *output.CreatePublisher, err error)
	GetByID(ctx context.Context, publisherID int64) (out *output.GetPublisher, err error)
	List(ctx context.Context) (out *output.ListPublishers, err error)
	SearchByName(ctx context.Context, name string) (out *output.ListPublishers, err error)
}