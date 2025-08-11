package adaptor

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func NewPublisher(publisher *model.Publisher) oapi.Publisher {
	return oapi.Publisher{
		Id:        publisher.ID,
		Name:      publisher.Name,
		CreatedAt: publisher.CreatedAt,
		UpdatedAt: publisher.UpdatedAt,
	}
}
func NewPublishers(publishers []*model.Publisher) []oapi.Publisher {
	result := make([]oapi.Publisher, len(publishers))
	for i, publisher := range publishers {
		result[i] = NewPublisher(publisher)
	}
	return result
}
