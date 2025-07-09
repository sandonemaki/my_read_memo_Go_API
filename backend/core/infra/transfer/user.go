package transfer

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

func ToUserEntity(user *model.User) *dbmodels.User {
	return &dbmodels.User{
		Ulid:        user.Ulid,
		UID:         user.UID,
		DisplayName: user.DisplayName,
		DeletedAt:   user.DeletedAt,
	}
}
