//go:build wireinject

package usecase

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

func NewUserDI(*sql.DB) (_ User) {
	wire.Build(
		db.NewDB,
		query.NewUser,
		repository.NewUser,
		NewUser,
	)
	return
}
