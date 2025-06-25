package model

import (
	"database/sql"
	"time"
)

type User struct {
	Ulid      string
	Nickname  string
	DeletedAt sql.Null[time.Time]
}

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
