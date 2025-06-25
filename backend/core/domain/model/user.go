package model

import (
	"database/sql"
	"time"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/dbmodels"
	"github.com/volatiletech/null/v8"
)

type User dbmodels.User

func NewUser(
	ulid string,
	nickname string,
	deletedAt null.Time,
) *User {
	var sqlDeletedAt sql.Null[time.Time]
	if deletedAt.Valid {
		sqlDeletedAt = sql.Null[time.Time]{
			V:     deletedAt.Time,
			Valid: true,
		}
	} else {
		sqlDeletedAt = sql.Null[time.Time]{
			V: time.Time{},
		}
	}
	return &User{
		Ulid:      ulid,
		Nickname:  nickname,
		DeletedAt: sqlDeletedAt,
	}
}
