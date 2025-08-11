package output

import "github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"

func NewCreateAuthor(author *model.Author) *CreateAuthor {
	return &CreateAuthor{
		Author: author,
	}
}

type CreateAuthor struct {
	Author *model.Author
}

func NewGetAuthor(author *model.Author) *GetAuthor {
	return &GetAuthor{
		Author: author,
	}
}

type GetAuthor struct {
	Author *model.Author
}

func NewListAuthors(authors []*model.Author) *ListAuthors {
	return &ListAuthors{
		Authors: authors,
	}
}

type ListAuthors struct {
	Authors []*model.Author
}
