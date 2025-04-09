package logger

import (
	"log/slog"
	"os"
)

const (
	LocalEnv = "local"
	DevEnv   = "dev"
	ProdEnv  = "prod"
)

func Init(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case LocalEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case DevEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case ProdEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
