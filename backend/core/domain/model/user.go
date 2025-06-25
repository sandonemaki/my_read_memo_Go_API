package model

import (
	"database/sql"
	"time"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
)

type User dbmodels.User

func NewUser(
	ulid string,
	nickname string,
	deletedAt sql.Null[time.Time],
) *User {

	return &User{
		Ulid:      ulid,
		Nickname:  nickname,
		DeletedAt: deletedAt,
	}
}
