package input

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/errof"
)

func NewCreateMasterBook(title string, isbn string, authorName string, publisherName string, coverS3URL string, totalPages int, publishedAt time.Time) CreateMasterBook {
	return CreateMasterBook{
		Title:         title,
		ISBN:          isbn,
		AuthorName:    authorName,
		PublisherName: publisherName,
		CoverS3URL:    coverS3URL,
		TotalPages:    totalPages,
		PublishedAt:   publishedAt,
	}
}

type CreateMasterBook struct {
	Title         string    `validate:"required,min=1,max=200"`
	ISBN          string    `validate:"omitempty,isbn"`
	AuthorName    string    `validate:"required,min=1,max=100"`
	PublisherName string    `validate:"required,min=1,max=100"`
	CoverS3URL    string    `validate:"omitempty,url"`
	TotalPages    int       `validate:"required,min=1"`
	PublishedAt   time.Time `validate:"required"`
}

func (p *CreateMasterBook) Validate() error {
	if err := validate.Struct(p); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func (c *CreateMasterBook) GetPublishedAt() sql.Null[time.Time] {
	return sql.Null[time.Time]{
		V:     c.PublishedAt,
		Valid: true,
	}
}
