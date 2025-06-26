package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

type User struct {
	userRepo repository.User
}

func NewUser(userRepo repository.User) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) Create(ctx context.Context, p input.CreateUser) (result *output.User, err error) {
	// get user by ulid
	user := &model.User{
		UID:      p.UID,
		Nickname: p.Nickname,
	}

	return &output.User{
		ULID:      user.Ulid,
		UID:       user.UID,
		Nickname:  user.Nickname,
		DeletedAt: user.DeletedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}, nil
}
