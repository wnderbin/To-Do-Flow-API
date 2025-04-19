package middleware

import (
	"net/http"
	"os"

	"todoflow-api/internal/database"

	"github.com/gin-gonic/gin"
)

// --- BASE URL'S ---

func ExitMiddleware(c *gin.Context) {
	database.CloseConRedis()
	database.ClosePostgres()
	os.Exit(0)
}

func MainPageMiddleware(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to main page!\nYou can find the API documentation in text form in the api/docs directory.")
}

// --- GET URL'S ---

func GetUserMiddleware(c *gin.Context) {
	user_id := c.Query("user_id")

	user, status := database.GetUser(c.Request.Context(), user_id)
	if status {
		c.JSON(http.StatusAccepted, user)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
	}
}

func GetUserByUsernameMiddleware(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, status := database.GetUserByUsername(c.Request.Context(), username, password)
	if status {
		c.JSON(http.StatusAccepted, user)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent user")
	}
}

func GetToDoNoteMiddleWare(c *gin.Context) {
	note_id := c.Query("note_id")
	user_id := c.Query("user_id")

	note, status := database.GetToDoNote(c.Request.Context(), note_id, user_id)
	if status {
		c.JSON(http.StatusAccepted, note)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
	}
}

func GetToDoNotesMiddleware(c *gin.Context) {
	user_id := c.Query("user_id")

	notes, status := database.GetToDoNotes(c.Request.Context(), user_id)
	if status {
		for _, note := range notes {
			c.JSON(http.StatusAccepted, note)
		}
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
	}
}

// --- POST URL'S ---

func CreateUserMiddleware(c *gin.Context) {
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
}

func CreateNoteMiddleware(c *gin.Context) {
	note := c.Query("note")
	user_id := c.Query("user_id")

	note_obj, status := database.CreateToDo(note, user_id)
	if status {
		c.JSON(http.StatusCreated, note_obj)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
}

// --- PUT URL'S ---

func UpdateUserMiddleware(c *gin.Context) {
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
}

func UpdateNoteMiddleware(c *gin.Context) {
	note_id := c.Query("note_id")
	user_id := c.Query("user_id")
	new_note := c.Query("note")

	note, status := database.UpdateUserNote(note_id, user_id, new_note)
	if status {
		c.JSON(http.StatusAccepted, note)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou probably specified non-existent identifiers")
	}
}

// PATCH URL'S

func UpdateUserPasswordMiddleware(c *gin.Context) {
	uuid := c.Query("uuid")
	password := c.Query("password")

	user, status := database.UpdateUserPassword(uuid, password)
	if status {
		c.JSON(http.StatusAccepted, user)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
	}
}

// DELETE URL'S

func DeleteUserMiddleware(c *gin.Context) {
	user_id := c.Query("user_id")

	user, status := database.DeleteUser(user_id)
	if status {
		c.JSON(http.StatusAccepted, user)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
	}
}

func DeleteNoteMiddleware(c *gin.Context) {
	note_id := c.Query("note_id")
	user_id := c.Query("user_id")

	note, status := database.DeleteNote(note_id, user_id)
	if status {
		c.JSON(http.StatusAccepted, note)
	} else {
		c.String(http.StatusInternalServerError, "Internal Server Error\nYou may have specified a non-existent id")
	}
}
