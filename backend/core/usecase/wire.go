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

func NewPublisherDI(*sql.DB) (_ Publisher) {
	wire.Build(
		db.NewDB,
		query.NewPublisher,
		repository.NewPublisher,
		NewPublisher,
	)
	return
}

// func NewAuthorDI(*sql.DB) (_ Author) {
// 	wire.Build(
// 		db.NewDB,
// 		query.NewAuthor,
// 		repository.NewAuthor,
// 		NewAuthor,
// 	)
// 	return
// }
