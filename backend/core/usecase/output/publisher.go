package output

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

// CreatePublisher is output for creating a new publisher
func NewCreatePublisher(publisher *model.Publisher) *CreatePublisher {
	return &CreatePublisher{
		Publisher: publisher,
	}
}

type CreatePublisher struct {
	Publisher *model.Publisher
}

// GetPublisher is output for getting a publisher
func NewGetPublisher(publisher *model.Publisher) *GetPublisher {
	return &GetPublisher{
		Publisher: publisher,
	}
}

type GetPublisher struct {
	Publisher *model.Publisher
}

// ListPublishers is output for listing publishers
func NewListPublishers(publishers []*model.Publisher) *ListPublishers {
	return &ListPublishers{
		Publishers: publishers,
	}
}

type ListPublishers struct {
	Publishers []*model.Publisher
}