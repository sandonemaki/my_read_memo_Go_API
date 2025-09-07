package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

type masterBook struct {
	masterBookQuery query.MasterBook
	masterBookRepo  repository.MasterBook
	authorQuery     query.Author
	publisherQuery  query.Publisher
}

func NewMasterBook(masterBookQuery query.MasterBook, masterBookRepo repository.MasterBook, authorQuery query.Author, publisherQuery query.Publisher) *masterBook {
	return &masterBook{
		masterBookQuery: masterBookQuery,
		masterBookRepo:  masterBookRepo,
		authorQuery:     authorQuery,
		publisherQuery:  publisherQuery,
	}
}

func (u *masterBook) Create(ctx context.Context, in input.CreateMasterBook) (out *output.CreateMasterBook, err error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	// 1. 著者の処理（名前→ID）
	author, err := u.authorQuery.GetByName(ctx, in.AuthorName, false)
	if err != nil {
		return nil, err
	}
	// 2. 出版社の処理（名前→ID）
	publisher, err := u.publisherQuery.GetByName(ctx, in.PublisherName, false)
	if err != nil {
		return nil, err
	}

	masterBook := &model.MasterBook{
		Title:       in.Title,
		Isbn:        in.ISBN,
		AuthorID:    author.ID,    // 取得したID
		PublisherID: publisher.ID, // 取得したID
		CoverS3URL:  in.CoverS3URL,
		TotalPage:   int32(in.TotalPages),
		PublishedAt: in.GetPublishedAt(),
	}
	masterBookID, err := u.masterBookRepo.Create(ctx, masterBook)
	if err != nil {
		return nil, err
	}
	masterBook.ID = masterBookID
	return output.NewCreateMasterBook(masterBook, author, publisher), nil
}
