package handler

import (
	"log/slog"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/firebase"
)

type Core struct {
	logger      *slog.Logger
	firebase    firebase.Glue
	userUsecase usecase.User
}

func NewCore(
	logger *slog.Logger,
	firebase firebase.Glue,
	userUsecase usecase.User,
) *Core {
	return &Core{
		logger:      logger,
		firebase:    firebase,
		userUsecase: userUsecase,
	}
}
