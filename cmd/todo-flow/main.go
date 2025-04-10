package main

import (
	"log/slog"

	"todoflow-api/internal/config"

	"todoflow-api/internal/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.MustLoad()

	log := logger.Init(conf.Env)
	log = log.With(slog.String("env", conf.Env))

	log.Info("Initializing server", slog.String("Address", conf.HttpServer.Address))
	log.Debug("Logger debug mode enabled")

	router := gin.Default()

	router.Run(conf.HttpServer.Address)
}
