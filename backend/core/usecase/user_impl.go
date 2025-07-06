package usecase

import (
	"context"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
	"github.com/volatiletech/null"
)

type User struct {
	userQuery query.User
	userRepo  repository.User
}

// NewUser creates a new User usecase instance
func NewUser(
	userQuery query.User,
	userRepo repository.User) *User {
	return &User{
		userQuery: userQuery,
		userRepo:  userRepo,
	}
}

func (u *User) Create(ctx context.Context, p input.CreateUser) (result *output.CreateUser, err error) {
	// get user by ulid
	user := &model.User{
		UID:      p.UID,
		Nickname: p.Nickname,
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &output.CreateUser{
		User: user,
	}, nil
}

func (u *User) GetMe(ctx context.Context, input input.GetCurrentUserDetail) (result *output.GetUser, err error) {
	var user *model.User
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user, err = u.userQuery.GetByULID(ctx, query.UserGetQuery{
		ULID: null.StringFrom(input.UID),
	})
	if err != nil {
		return nil, err
	}

	return &output.GetUser{
		User: user,
	}, nil
}

func (u *User) UpdateNickname(ctx context.Context, input input.UpdateUser) (result *output.UpdateUser, err error) {
	// ulidを取得する
	user, err := u.userQuery.GetByULID(ctx, query.UserGetQuery{
		ULID: null.StringFrom(input.ULID),
	})
	if err != nil {
		return nil, err
	}

	// ユーザー情報を更新
	user.Nickname = input.Nickname

	// データベースを更新
	var userUlid string
	userUlid, err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	// 更新後のユーザー情報を取得
	updatedUser, err := u.userQuery.GetByULID(ctx, query.UserGetQuery{
		ULID: null.StringFrom(userUlid),
	})
	if err != nil {
		return nil, err
	}

	return &output.UpdateUser{
		User: updatedUser,
	}, nil
}

func (u *User) Delete(ctx context.Context, input input.DeleteUser) (err error) {
	// ulidを取得する
	return u.userRepo.Delete(ctx, input.ULID)
}
