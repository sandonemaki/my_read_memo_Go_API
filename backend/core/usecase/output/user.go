package output

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
)

type GetUser struct {
	User *model.User
}

func NewGetUser(user *model.User) *GetUser {
	return &GetUser{
		User: user,
	}
}

type CreateUser struct {
	User *model.User
}

func NewCreateUser(user *model.User) *CreateUser {
	return &CreateUser{
		User: user,
	}
}

type UpdateUser struct {
	User *model.User
}

func NewUpdateUser(user *model.User) *UpdateUser {
	return &UpdateUser{
		User: user,
	}
}
