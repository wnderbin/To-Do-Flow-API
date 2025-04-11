package main

import (
	"log/slog"
	"net/http"
	"os"

	"todoflow-api/internal/config"
	"todoflow-api/internal/database"
	"todoflow-api/internal/logger"
	"todoflow-api/internal/migrator"

	"github.com/gin-gonic/gin"
)

func api_init() *config.Config {
	conf := config.MustLoad()

	log := logger.Init(conf.Env)
	log = log.With(slog.String("env", conf.Env))

	log.Info("Initializing server", slog.String("Address", conf.HttpServer.Address))
	log.Debug("Logger debug mode enabled")

	db, err := database.SQLiteDBInit(conf)
	if err != nil {
		log.Error("filed to connect to database:", "error", err)
		os.Exit(1)
	}
	log.Info("connected to database")

	sqldb, _ := db.DB()

	if err := migrator.ApplySQLiteMigrations(sqldb); err != nil {
		log.Error("migrations failed", "error", err)
		os.Exit(1)
	}

	return conf
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "1" {
		api_init()
	} else {
		conf := api_init()

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
