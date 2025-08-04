package repository

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

//go:generate mockgen -source author.go -destination mock/mock_author.go
type Author interface {
	// Create creates a new author.
	Create(ctx context.Context, author *model.Author) (authorID int64, err error)
	// Update updates an existing author.
	// TODO: OpenAPIにPUTエンドポイントを追加後に有効化
	// Update(ctx context.Context, author *model.Author) (authorID int64, err error)
}
