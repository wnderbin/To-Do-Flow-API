package main

import (
	"log/slog"
	"os"

	"todoflow-api/internal/config"
	"todoflow-api/internal/database"
	"todoflow-api/internal/logger"
	"todoflow-api/internal/middleware"
	"todoflow-api/internal/migrator"

	"github.com/gin-gonic/gin"
)

func api_init() *config.Config {
	conf := config.MustLoad()

	log := logger.Init(conf.Env)
	log = log.With(slog.String("env", conf.Env))

	log.Info("Initializing server", slog.String("Address", conf.HttpServer.Address))
	log.Debug("Logger debug mode enabled")

	if conf.Postgres.Status == 0 {
		sqldb, _ := database.Postgres_db.DB()

		if err := migrator.ApplyPostgresMigrations(sqldb); err != nil {
			log.Error("migrations failed", "error", err)
			os.Exit(1)
		}

		return conf
	} else {
		return conf
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "1" { //
		api_init()
	} else {
		conf := api_init()
		router := gin.Default()

		// get-urls
		router.GET("/exit", middleware.ExitMiddleware)
		router.GET("/main", middleware.MainPageMiddleware)
		router.GET("/get_user", middleware.GetUserMiddleware)
		router.GET("/get_user_by_username", middleware.GetUserByUsernameMiddleware)
		router.GET("/get_note", middleware.GetToDoNoteMiddleWare)
		router.GET("/get_notes", middleware.GetToDoNotesMiddleware)

		// post-urls
		router.POST("/create_user", middleware.CreateUserMiddleware)
		router.POST("/create_note", middleware.CreateNoteMiddleware)

		// put-urls
		router.PUT("/update_user", middleware.UpdateUserMiddleware)
		router.PUT("/update_note", middleware.UpdateNoteMiddleware)

		// patch-urls
		router.PATCH("/update_user_password", middleware.UpdateUserPasswordMiddleware)

		// delete-urls
		router.DELETE("/delete_user", middleware.DeleteUserMiddleware)
		router.DELETE("/delete_note", middleware.DeleteNoteMiddleware)

		router.Run(conf.HttpServer.Address)
	}
}
