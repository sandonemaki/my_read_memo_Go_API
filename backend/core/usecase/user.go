package usecase

import (
	"context"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

type User interface {
	GetCurrentUser(ctx context.Context, input input.GetCurrentUserDetail) (result *output.GetUser, err error)
	Create(ctx context.Context, input input.CreateUser) (result *output.CreateUser, err error)
	UpdateCurrentUser(ctx context.Context, input input.UpdateUser) (result *output.UpdateUser, err error)
	Delete(ctx context.Context, input input.DeleteUser) (err error)
}
