package usecase

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
	"github.com/volatiletech/null"
)

type user struct {
	userQuery query.User
	userRepo  repository.User
}

// NewUser creates a new User usecase instance
func NewUser(
	userQuery query.User,
	userRepo repository.User) User {
	return &user{
		userQuery: userQuery,
		userRepo:  userRepo,
	}
}

func (u *user) Create(ctx context.Context, p input.CreateUser) (result *output.CreateUser, err error) {
	// Generate ULID - テスト用に一時的に固定値
	ulidValue := ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
	
	user := &model.User{
		Ulid:        ulidValue.String(),
		UID:         p.UID,
		DisplayName: p.DisplayName,
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// INSERT後にGETして実際のデータベースの値を取得（見本パターン）
	createdUser, err := u.userQuery.GetByUID(ctx, query.UserGetQuery{
		UID: null.StringFrom(user.UID),
	})
	if err != nil {
		return nil, err
	}

	return &output.CreateUser{
		User: createdUser, // GETした結果を返す（mockの固定値）
	}, nil
}

func (u *user) GetCurrentUser(ctx context.Context, input input.GetCurrentUserDetail) (result *output.GetUser, err error) {
	var user *model.User
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user, err = u.userQuery.GetByUID(ctx, query.UserGetQuery{
		UID: null.StringFrom(input.UID),
	})
	if err != nil {
		return nil, err
	}

	return &output.GetUser{
		User: user,
	}, nil
}

func (u *user) UpdateCurrentUser(ctx context.Context, input input.UpdateUser) (result *output.UpdateUser, err error) {
	var user *model.User
	user, err = u.userQuery.GetByUID(ctx, query.UserGetQuery{
		UID: null.StringFrom(input.UID),
	})
	if err != nil {
		return nil, err
	}

	// ユーザー情報を更新
	user.DisplayName = input.DisplayName

	// データベースを更新
	_, err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	// 更新後のユーザー情報を取得
	updatedUser, err := u.userQuery.GetByUID(ctx, query.UserGetQuery{
		UID: null.StringFrom(input.UID),
	})
	if err != nil {
		return nil, err
	}

	return &output.UpdateUser{
		User: updatedUser,
	}, nil
}

func (u *user) Delete(ctx context.Context, input input.DeleteUser) (err error) {
	if err := u.userRepo.Delete(ctx, input.UID); err != nil {
		return err
	}

	return nil
}
