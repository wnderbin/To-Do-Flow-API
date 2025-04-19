package database

import (
	"context"
	"time"
	"todoflow-api/models"
)

func UpdateUser(uuid string, name string, username string, password string, email string) (models.User, bool) {
	var user models.User

	err := db.Where("id = ? AND password = ?", uuid, password).First(&user).Error
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	user.Name = name
	user.Username = username
	user.Password = password
	user.Email = email
	user.Updated_at = time.Now()

	err = db.Save(&user).Error
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	return user, true
}

func UpdateUserPassword(uuid string, password string) (models.User, bool) {
	var user models.User

	err := db.Where("id = ?", uuid).First(&user).Error
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	user.Password = password
	user.Updated_at = time.Now()

	err = db.Save(&user).Error
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	return user, true
}

func UpdateUserNote(note_uuid, user_uuid, new_note string) (models.ToDoList, bool) {
	var note models.ToDoList

	note, status := GetToDoNote(context.Background(), note_uuid, user_uuid)
	if !status {
		log.Error("bad status")
		return note, false
	}

	note.Todonote = new_note
	note.Updated_at = time.Now()

	err := db.Save(&note).Error
	if err != nil {
		log.Error(err.Error())
		return note, false
	}
	return note, true
}
