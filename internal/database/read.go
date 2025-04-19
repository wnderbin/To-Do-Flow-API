package database

import (
	"context"
	"encoding/json"
	"time"
	"todoflow-api/models"

	"github.com/redis/go-redis/v9"
)

func GetUser(ctx context.Context, uuid_str string) (models.User, bool) { // redis works (verified)
	var user models.User

	cached_user, err := RDB.Get(ctx, uuid_str).Result()
	log.Info(cached_user)
	if err == nil {
		if err = json.Unmarshal([]byte(cached_user), &user); err == nil {
			return user, true
		} else {
			log.Error(err.Error())
			return user, false
		}
	} else if err == redis.Nil {

		err = db.Where("id = ?", uuid_str).First(&user).Error
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		err = RDB.Set(ctx, uuid_str, userJSON, 5*time.Minute).Err()
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		return user, true
	}
	return user, false
}

func GetUserByUsername(ctx context.Context, username string, password string) (models.User, bool) { // redis works (verified)
	var user models.User

	cached_user, err := RDB.Get(ctx, username).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(cached_user), &user); err == nil {
			if user.Password == password {
				return user, true
			} else {
				return user, false
			}
		}
	} else if err == redis.Nil {

		err = db.Where("username = ? AND password = ?", username, password).First(&user).Error
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		err = RDB.Set(ctx, username, userJSON, 5*time.Minute).Err()
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		return user, true
	}
	return user, false
}

func GetToDoNotes(ctx context.Context, user_uuid string) ([]models.ToDoList, bool) { // redis works (verified)
	var notes []models.ToDoList

	cached_notes, err := RDB.Get(ctx, user_uuid).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(cached_notes), &notes); err == nil {
			return notes, true
		} else {
			log.Error(err.Error())
			return notes, false
		}
	} else if err == redis.Nil {

		err = db.Where("user_id = ?", user_uuid).Find(&notes).Error
		if err != nil {
			log.Error(err.Error())
			return notes, false
		}

		user, status := GetUser(context.Background(), user_uuid)
		if !status {
			log.Error("bad status")
			return notes, false
		}

		for i := range notes {
			notes[i].User = user
		}

		notesJSON, err := json.Marshal(notes)
		if err != nil {
			log.Error(err.Error())
			return notes, false
		}
		err = RDB.Set(ctx, user_uuid, notesJSON, 5*time.Minute).Err()
		if err != nil {
			log.Error(err.Error())
			return notes, false
		}

		return notes, true
	}

	return notes, false
}

func GetToDoNote(ctx context.Context, note_uuid_str string, user_uuid string) (models.ToDoList, bool) { // redis works (verified)
	var note models.ToDoList

	cached_note, err := RDB.Get(ctx, note_uuid_str).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(cached_note), &note); err == nil {
			if note.User_id == user_uuid {
				return note, true
			} else {
				return note, false
			}
		} else {
			log.Error(err.Error())
			return note, false
		}
	} else if err == redis.Nil {

		err = db.Where("id = ?", note_uuid_str).Find(&note).Error
		if err != nil {
			log.Error(err.Error())
		}

		user, status := GetUser(context.Background(), user_uuid)
		if !status {
			log.Error("bad status")
			return note, false
		}
		note.User = user

		noteJSON, _ := json.Marshal(note)
		err = RDB.Set(ctx, note_uuid_str, noteJSON, 5*time.Minute).Err()
		if err != nil {
			log.Error(err.Error())
			return note, false
		}

		return note, true
	}
	return note, false
}
