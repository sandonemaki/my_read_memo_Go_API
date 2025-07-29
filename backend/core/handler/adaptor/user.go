package adaptor

import (
	"time"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/oapi"
)

func NewUser(user *model.User) oapi.User {
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.V
	}

	return oapi.User{
		Ulid:        user.Ulid,
		Uid:         user.UID,
		DisplayName: user.DisplayName,
		DeletedAt:   deletedAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
