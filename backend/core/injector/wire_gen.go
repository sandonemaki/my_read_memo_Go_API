// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/handler"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/query"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/infra/repository"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/db"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/firebase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/logger"
	"log/slog"
)

// Injectors from wire.go:

func InitializeCoreHandler(loggerConfig config.Logger, postgresConfig config.Postgres, applicationName string) (*handler.Core, error) {
	slogHandler := logger.NewLogger(loggerConfig)
	slogLogger := slog.New(slogHandler)
	glue, err := firebase.NewFirebaseGlue()
	if err != nil {
		return nil, err
	}
	sqlDB := db.NewPSQL(postgresConfig, slogLogger, applicationName)
	client := db.NewDB(sqlDB)
	user := query.NewUser(client)
	repositoryUser := repository.NewUser(client)
	usecaseUser := usecase.NewUser(user, repositoryUser)
	core := handler.NewCore(slogLogger, glue, usecaseUser)
	return core, nil
}
