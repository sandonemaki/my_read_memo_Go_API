package model

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

type Publisher dbmodels.Publisher

func NewPublisher(name string) *Publisher {
	return &Publisher{
		Name: name,
	}
}