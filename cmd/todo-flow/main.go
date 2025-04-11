package main

import (
	"log/slog"
	"net/http"
	"os"

	"todoflow-api/internal/config"
	"todoflow-api/internal/database"
	"todoflow-api/internal/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "1" {
		conf := config.MustLoad()

		log := logger.Init(conf.Env)
		log = log.With(slog.String("env", conf.Env))

		log.Info("Initializing server", slog.String("Address", conf.HttpServer.Address))
		log.Debug("Logger debug mode enabled")
	} else {
		conf := config.MustLoad()

		log := logger.Init(conf.Env)
		log = log.With(slog.String("env", conf.Env))

		log.Info("Initializing server", slog.String("Address", conf.HttpServer.Address))
		log.Debug("Logger debug mode enabled")

		router := gin.Default()

		router.GET("/main", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome to main page")
		})
		router.GET("/create_user", func(c *gin.Context) {
			// /create_user?name=_&username=_&password=_&email=_
			// _ - your data
			name := c.Query("name")
			username := c.Query("username")
			password := c.Query("password")
			email := c.Query("email")

			status := database.CreateUser(name, username, password, email)
			if status {
				c.String(http.StatusOK, "Success")
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error")
			}
		})

		router.Run(conf.HttpServer.Address)
	}
}
