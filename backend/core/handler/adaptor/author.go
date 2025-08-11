package adaptor

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func NewAuthor(author *model.Author) oapi.Author {
	return oapi.Author{
		Id:        author.ID,
		Name:      author.Name,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func NewAuthors(authors []*model.Author) []oapi.Author {
	result := make([]oapi.Author, len(authors))
	for i, author := range authors {
		result[i] = NewAuthor(author)
	}
	return result
}