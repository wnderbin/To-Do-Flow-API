package main

import (
	"log/slog"
	"todoflow-api/internal/config"
	"todoflow-api/internal/logger"
)

func main() {
	conf := config.MustLoad()

	log := logger.Init(conf.Env)
	log = log.With(slog.String("Env", conf.Env))

	log.Info("Initializing server", slog.String("Address", conf.HttpServer.Address))
	log.Debug("Logger debug mode enabled")
}
