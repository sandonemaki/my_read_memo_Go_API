package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

//go:generate mockgen -source publisher.go -destination mock/mock_publisher.go
type Publisher interface {
	// Create creates a new publisher.
	Create(ctx context.Context, publisher *model.Publisher) (publisherID int64, err error)
	// Update updates an existing publisher.
	Update(ctx context.Context, publisher *model.Publisher) (publisherID int64, err error)
}
