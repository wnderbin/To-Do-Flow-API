package database

import (
	"context"
	"todoflow-api/models"
)

func DeleteUser(user_id string) (models.User, bool) {
	var user models.User
	var note models.ToDoList
	var return_user models.User

	err := db.Where("id = ?", user_id).Find(&return_user).Error
	if err != nil {
		log.Error(err.Error())
		return return_user, false
	}
	err = db.Where("user_id = ?", user_id).Delete(&note).Error
	if err != nil {
		log.Error(err.Error())
		return return_user, false
	}
	err = db.Where("id = ?", user_id).Delete(&user).Error
	if err != nil {
		log.Error(err.Error())
		return return_user, false
	}

	return return_user, true
}

func DeleteNote(note_id string, user_id string) (models.ToDoList, bool) {
	var note models.ToDoList
	var return_note models.ToDoList

	return_note, status := GetToDoNote(context.Background(), note_id, user_id)
	if !status {
		log.Error("bad status")
		return return_note, false
	}

	err := db.Where("id = ?", note_id).Delete(&note).Error
	if err != nil {
		log.Error(err.Error())
		return return_note, false
	}

	return return_note, true
}
