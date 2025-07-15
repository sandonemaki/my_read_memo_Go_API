package logger

import (
	"log/slog"
	"os"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
)

func NewLogger(c config.Logger) slog.Handler {
	level := slog.LevelInfo
	if c.Debug {
		level = slog.LevelDebug
	}
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
}