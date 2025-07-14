//go:build wireinject

package injector

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
)

func NewUserUseCase(sqlDB *sql.DB) (_ usecase.User) {
	wire.Build(
		db.NewDB,
		query.NewUser,
		repository.NewUser,
		usecase.NewUser,
	)
	return
}
