package output

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

func NewCreateMasterBook(masterBook *model.MasterBook, author *model.Author, publisher *model.Publisher) *CreateMasterBook {
	return &CreateMasterBook{
		MasterBook: masterBook,
		Author:     author,
		Publisher:  publisher,
	}
}

type CreateMasterBook struct {
	MasterBook *model.MasterBook
	Author     *model.Author
	Publisher  *model.Publisher
}
