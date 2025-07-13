package input

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/errof"
)

func NewCreateUser(UID string, name string) CreateUser {
	return CreateUser{
		UID:         UID,
		DisplayName: name,
	}
}

type CreateUser struct {
	Ulid        string `validate:""`
	UID         string `validate:"required"`
	DisplayName string `validate:"required"`
}

// SetDefaults はデフォルト値を設定します（見本に従う）
func (i *CreateUser) SetDefaults() {
	if i.Ulid == "" {
		i.Ulid = ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).String()
	}
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
		UID:         UID,
		DisplayName: name,
	}
}

type UpdateUser struct {
	UID         string `validate:"required"`
	DisplayName string `validate:"required"`
}

func (u *UpdateUser) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewDeleteUser(UID string) DeleteUser {
	return DeleteUser{
		UID: UID,
	}
}

type DeleteUser struct {
	UID string `validate:"required"`
}

func (u *DeleteUser) Validate() error {
	if err := validate.Struct(u); err != nil {
		return errors.Wrap(errof.ErrInvalidRequest, err.Error())
	}
	return nil
}
