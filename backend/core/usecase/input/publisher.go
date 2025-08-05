package input

import (
	"github.com/pkg/errors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/errof"
)

// CreatePublisher is input for creating a new publisher
func NewCreatePublisher(name string) CreatePublisher {
	return CreatePublisher{
		Name: name,
	}
}

type CreatePublisher struct {
	Name string `validate:"required,min=1,max=100"`
}

func (p *CreatePublisher) Validate() error {
	if err := validate.Struct(p); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

// GetPublisherByID is input for getting a publisher by ID
func NewGetPublisherByID(id int64) GetPublisherByID {
	return GetPublisherByID{
		ID: id,
	}
}

type GetPublisherByID struct {
	ID int64 `validate:"required,min=1"`
}

func (p *GetPublisherByID) Validate() error {
	if err := validate.Struct(p); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

// SearchPublisher is input for searching publishers by name
func NewSearchPublisher(name string) SearchPublisher {
	return SearchPublisher{
		Name: name,
	}
}

type SearchPublisher struct {
	Name string `validate:"required,min=1,max=100"`
}

func (p *SearchPublisher) Validate() error {
	if err := validate.Struct(p); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}