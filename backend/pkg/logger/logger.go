package logger

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
)

func NewLogger(c config.Logger) slog.Handler {
	level := slog.LevelInfo
	if c.Debug {
		level = slog.LevelDebug
	}
	return NewOriginalHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
}

type OriginalHandler struct {
	jsonHandler *slog.JSONHandler
}

func NewOriginalHandler(w io.Writer, opts *slog.HandlerOptions) *OriginalHandler {
	return &OriginalHandler{
		jsonHandler: slog.NewJSONHandler(w, opts),
	}
}

func (h *OriginalHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.jsonHandler.Enabled(ctx, level)
}

func (h *OriginalHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.jsonHandler.WithAttrs(attrs)
}

func (h *OriginalHandler) WithGroup(name string) slog.Handler {
	return h.jsonHandler.WithGroup(name)
}

func (h *OriginalHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.jsonHandler.Handle(ctx, r)
}
