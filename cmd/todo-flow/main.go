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
		router.GET("/exit", func(c *gin.Context) {
			database.CloseConRedis()
			database.ClosePostgres()
			os.Exit(0)
		})
		router.GET("/main", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome to main page!\nYou can find the API documentation in text form in the api/docs directory.")
		})
		router.GET("/get_user", func(c *gin.Context) {
			// /get_user?user_id=_
			// _ - your data
			user_id := c.Query("user_id")

			user, status := database.GetUser(c.Request.Context(), user_id)
			if status {
				c.JSON(http.StatusAccepted, user)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
			}
		})
		router.GET("/get_user_by_username", func(c *gin.Context) {
			// /get_user_by_username?username=_&password=X
			// _ - your data
			username := c.Query("username")
			password := c.Query("password")

			user, status := database.GetUserByUsername(c.Request.Context(), username, password)
			if status {
				c.JSON(http.StatusAccepted, user)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent user")
			}
		})
		router.GET("/get_note", func(c *gin.Context) {
			// /get_note?note_id=_&user_id=_
			// _ - your data
			note_id := c.Query("note_id")
			user_id := c.Query("user_id")

			note, status := database.GetToDoNote(c.Request.Context(), note_id, user_id)
			if status {
				c.JSON(http.StatusAccepted, note)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
			}
		})
		router.GET("/get_notes", func(c *gin.Context) {
			// /get_notes?user_id=_
			// _ - your data
			user_id := c.Query("user_id")

			notes, status := database.GetToDoNotes(c.Request.Context(), user_id)
			if status {
				for _, note := range notes {
					c.JSON(http.StatusAccepted, note)
				}
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
			}
		})

		// post-urls
		router.POST("/create_user", func(c *gin.Context) {
			// /create_user?name=_&username=_&password=_&email=_
			// _ - your data
			name := c.Query("name")
			username := c.Query("username")
			password := c.Query("password")
			email := c.Query("email")

			user, status := database.CreateUser(name, username, password, email)
			if status {
				c.JSON(http.StatusCreated, user)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nMost likely, there is already a user with the specified username, try changing it")
			}
		})
		router.POST("/create_note", func(c *gin.Context) {
			// /create_note?note=_&user_id=_
			// _ - your data
			note := c.Query("note")
			user_id := c.Query("user_id")

			note_obj, status := database.CreateToDo(note, user_id)
			if status {
				c.JSON(http.StatusCreated, note_obj)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error")
			}
		})

		// put-urls
		router.PUT("/update_user", func(c *gin.Context) {
			// /update_user?uuid=_&name=_&username=_&password=_&email=_
			// _ - your data
			uuid := c.Query("uuid")
			name := c.Query("name")
			username := c.Query("username")
			password := c.Query("password")
			email := c.Query("email")

			user, status := database.UpdateUser(uuid, name, username, password, email)
			if status {
				c.JSON(http.StatusAccepted, user)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a user ID and password that does not exist")
			}
		})
		router.PUT("/update_note", func(c *gin.Context) {
			// /update_note?note_id=_&user_id=_&note=_
			// _ - your data

			note_id := c.Query("note_id")
			user_id := c.Query("user_id")
			new_note := c.Query("note")

			note, status := database.UpdateUserNote(note_id, user_id, new_note)
			if status {
				c.JSON(http.StatusAccepted, note)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou probably specified non-existent identifiers")
			}
		})

		// patch-urls
		router.PATCH("/update_user_password", func(c *gin.Context) {
			// /update_user_password?uuid=_&password=_
			// _ - your data
			uuid := c.Query("uuid")
			password := c.Query("password")

			user, status := database.UpdateUserPassword(uuid, password)
			if status {
				c.JSON(http.StatusAccepted, user)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
			}
		})

		// delete-urls
		router.DELETE("/delete_user", func(c *gin.Context) {
			// /delete_user?user_id=_
			// _ - your data
			user_id := c.Query("user_id")

			user, status := database.DeleteUser(user_id)
			if status {
				c.JSON(http.StatusAccepted, user)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
			}
		})
		router.DELETE("/delete_note", func(c *gin.Context) {
			// /delete_note?note_id=_&user_id_
			// _ - your data
			note_id := c.Query("note_id")
			user_id := c.Query("user_id")

			note, status := database.DeleteNote(note_id, user_id)
			if status {
				c.JSON(http.StatusAccepted, note)
			} else {
				c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
			}
		})

		router.Run(conf.HttpServer.Address)
	}
}
