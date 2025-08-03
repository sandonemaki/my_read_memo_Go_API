package model

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

type Author dbmodels.Author

func NewAuthor(
	name string,
) *Author {
	return &Author{
		Name: name,
	}
}
