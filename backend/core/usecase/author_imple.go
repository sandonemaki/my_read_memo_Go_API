package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

type author struct {
	authorQuery query.Author      // データ取得用（SELECT）
	authorRepo  repository.Author // データ変更用（INSERT/UPDATE）
}

func NewAuthor(
	authorQuery query.Author,
	authorRepo repository.Author,
) Author { // Author インターフェースを返す（実装詳細を隠す）
	return &author{
		authorQuery: authorQuery,
		authorRepo:  authorRepo,
	}
}

func (a *author) Create(ctx context.Context, in input.CreateAuthor) (*output.CreateAuthor, error) {
	
	if err := in.Validate(); err != nil {
		return nil, err
	}

	author := &model.Author{
		Name: in.Name
	}

	authorID, err := a.autherRepo.Create(ctx, author)
	if err != nil {
		return nil, err
	}
	author.ID = authorID
	return output.NewCreateAuthor(author), nil
}

func (a *author) GetByID(ctx context.Context, authorID int64) (*output.GetAuthor, error) {
	author, err := a.authorQuery.GetByID(ctx, authorID)
	if err != nil {
		return nil, err
	}
	return output.NewGetAuthor(author), nil
}
