package output

import (
	"database/sql"
	"time"
)

type User struct {
	ULID      string
	UID       string
	Nickname  string
	DeletedAt sql.Null[time.Time]
	UpdatedAt time.Time
	CreatedAt time.Time
}
