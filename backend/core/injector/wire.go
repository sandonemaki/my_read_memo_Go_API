//go:build wireinject

package injector

import (
	"log/slog"

	"github.com/google/wire"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/handler"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/firebase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/logger"
)

func InitializeCoreHandler(config.Logger, config.Postgres, string) (*handler.Core, error) {
	wire.Build(
		db.NewDB,
		db.NewPSQL,
		logger.NewLogger,
		slog.New,
		firebase.NewFirebaseGlue,
		query.NewUser,
		repository.NewUser,
		usecase.NewUser,
		handler.NewCore,
	)
	return nil, nil
}
