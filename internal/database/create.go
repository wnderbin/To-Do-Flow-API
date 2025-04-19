package database

import (
	"context"
	"time"
	"todoflow-api/models"

	"github.com/google/uuid"
)

func CreateUser(name string, username string, password string, email string) (models.User, bool) {
	var user models.User

	user_uuid := uuid.NewString()

	err := db.Create(&models.User{
		Id:         user_uuid,
		Name:       name,
		Username:   username,
		Password:   password,
		Email:      email,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}).Error
	if err != nil {
		return user, false
	}

	user, status := GetUser(context.Background(), user_uuid)
	if !status {
		log.Error("bad status")
		return user, false
	}
	return user, true
}

func CreateToDo(ToDo_Note string, User_id string) (models.ToDoList, bool) {
	var note models.ToDoList

	note_uuid := uuid.NewString()

	db.Create(&models.ToDoList{
		Id:         note_uuid,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Todonote:   ToDo_Note,
		User_id:    User_id,
	})

	note, status := GetToDoNote(context.Background(), note_uuid, User_id)
	if !status {
		log.Error("bad status")
		return note, false
	}
	return note, true
}
