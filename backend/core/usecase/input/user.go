package input

import (
	"github.com/pkg/errors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/errof"
)

func NewCreateUser(UID string, name string) CreateUser {
	return CreateUser{
		UID:      UID,
		Nickname: name,
	}
}

type CreateUser struct {
	UID      string `validate:"required"`
	Nickname string `validate:"required"`
}

func (u *CreateUser) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewGetCurrentUserDetail(UID string) GetCurrentUserDetail {
	return GetCurrentUserDetail{
		UID: UID,
	}
}

type GetCurrentUserDetail struct {
	UID string `validate:"required"`
}

func (u *GetCurrentUserDetail) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewUpdateUser(UID string, name string) UpdateUser {
	return UpdateUser{
		Nickname: name,
	}
}

type UpdateUser struct {
	ULID     string `validate:"required"`
	Nickname string `validate:"required"`
}

func (u *UpdateUser) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewDeleteUser(ULID string) DeleteUser {
	return DeleteUser{
		ULID: ULID,
	}
}

type DeleteUser struct {
	ULID string `validate:"required"`
}

func (u *DeleteUser) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}
