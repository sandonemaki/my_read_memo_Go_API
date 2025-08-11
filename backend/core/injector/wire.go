//go:build wireinject

package injector

import (
	"log/slog"

	"github.com/google/wire"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/handler"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/auth/infra/firebase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/logger"
)

func InitializeCoreHandler(loggerConfig config.Logger, postgresConfig config.Postgres, applicationName string) (*handler.Core, error) {
	wire.Build(
		db.NewDB,
		db.NewPSQL,
		logger.NewLogger,
		slog.New,
		firebase.NewFirebaseAuthGlue,
		query.NewUser,
		repository.NewUser,
		repository.NewPublisher,
		// repository.NewAuthor,
		query.NewPublisher,
		// query.NewAuthor,
		usecase.NewUser,
		usecase.NewPublisher,
		// usecase.NewAuthor,
		handler.NewCore,
	)
	return nil, nil
}
