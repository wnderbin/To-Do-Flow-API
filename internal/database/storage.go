package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"todoflow-api/internal/config"
	"todoflow-api/internal/logger"
	"todoflow-api/internal/redis_db"
	"todoflow-api/models"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	conf *config.Config = config.MustLoad()
	log  slog.Logger    = *logger.Init(conf.Env)
	RDB  *redis.Client  = redis_db.Init(conf)
)

func CloseConRedis() {
	RDB.Close()
	log.Info("Closing connection with Redis...")
}

// sql

func SQLiteDBInit(conf *config.Config) (*gorm.DB, error) {
	dsn := conf.SQLite_path
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite_db: %w", err)
	}

	return db, nil
}

func SQLiteDBClose(db *gorm.DB) error {
	sqlite_db, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlite_db: %w", err)
	}

	if err = sqlite_db.Close(); err != nil {
		return fmt.Errorf("failed to close sqlite_db: %w", err)
	}

	return nil
}

// Crud

func CreateUser(name string, username string, password string, email string) (models.User, bool) {
	var user models.User

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	user_uuid := uuid.NewString()

	err = db.Create(&models.User{
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

	user, status := GetUser(user_uuid)
	if !status {
		log.Error("bad status")
		return user, false
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return user, false
	}
	return user, true
}

func CreateToDo(ToDo_Note string, User_id string) (models.ToDoList, bool) {
	var note models.ToDoList
	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	note_uuid := uuid.NewString()

	db.Create(&models.ToDoList{
		Id:         note_uuid,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Todonote:   ToDo_Note,
		User_id:    User_id,
	})

	note, status := GetToDoNote(note_uuid, User_id)
	if !status {
		log.Error("bad status")
		return note, false
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return note, false
	}
	return note, true
}

// cRud

func GetUser(uuid_str string) (models.User, bool) { // redis works (verified)
	var user models.User

	cached_user, err := RDB.Get(context.Background(), uuid_str).Result()
	if err == redis.Nil {
		db, err := SQLiteDBInit(conf)
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

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

		err = RDB.Set(context.Background(), uuid_str, userJSON, 5*time.Minute).Err()
		if err != nil {
			log.Error(err.Error())
			return user, false
		}

		if err = SQLiteDBClose(db); err != nil {
			log.Error(err.Error())
			return user, false
		}

		return user, true
	} else if err != nil {
		log.Error(err.Error())
		return user, false
	}

	if err = json.Unmarshal([]byte(cached_user), &user); err != nil {
		log.Error(err.Error())
		return user, false
	}
	return user, true
}

func GetUserByUsername(username string, password string) (models.User, bool) {
	var user models.User

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	err = db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return user, false
	}

	return user, true
}

func GetToDoNotes(user_uuid string) ([]models.ToDoList, bool) {
	var notes []models.ToDoList

	ParsedUUID, err := uuid.Parse(user_uuid)
	if err != nil {
		log.Error(err.Error())
		return notes, false
	}

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return notes, false
	}

	err = db.Where("user_id = ?", ParsedUUID).Find(&notes).Error
	if err != nil {
		log.Error(err.Error())
		return notes, false
	}

	user, status := GetUser(ParsedUUID.String())
	if !status {
		log.Error("bad status")
		return notes, false
	}

	for i := range notes {
		notes[i].User = user
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return notes, false
	}

	return notes, true
}

func GetToDoNote(note_uuid_str string, user_uuid string) (models.ToDoList, bool) {
	var note models.ToDoList

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	err = db.Where("id = ?", note_uuid_str).Find(&note).Error
	if err != nil {
		log.Error(err.Error())
	}

	user, status := GetUser(user_uuid)
	if !status {
		log.Error("bad status")
		return note, false
	}
	note.User = user

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return note, false
	}

	return note, true
}

// crUd

func UpdateUser(uuid string, name string, username string, password string, email string) (models.User, bool) {
	var user models.User

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	err = db.Where("id = ? AND password = ?", uuid, password).First(&user).Error
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

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return user, false
	}

	return user, true
}

func UpdateUserPassword(uuid string, password string) (models.User, bool) {
	var user models.User

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	err = db.Where("id = ?", uuid).First(&user).Error
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

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return user, false
	}

	return user, true
}

func UpdateUserNote(note_uuid, user_uuid, new_note string) (models.ToDoList, bool) {
	var note models.ToDoList

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	note, status := GetToDoNote(note_uuid, user_uuid)
	if !status {
		log.Error("bad status")
		return note, false
	}

	note.Todonote = new_note
	note.Updated_at = time.Now()

	err = db.Save(&note).Error
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return note, false
	}

	return note, true
}

// cruD

func DeleteUser(user_id string) (models.User, bool) {
	var user models.User
	var return_user models.User

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return return_user, false
	}

	err = db.Where("id = ?", user_id).Find(&return_user).Error
	if err != nil {
		log.Error(err.Error())
		return return_user, false
	}
	err = db.Where("id = ?", user_id).Delete(&user).Error
	if err != nil {
		log.Error(err.Error())
		return return_user, false
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return return_user, false
	}

	return return_user, true
}

func DeleteNote(note_id string, user_id string) (models.ToDoList, bool) {
	var note models.ToDoList
	var return_note models.ToDoList

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return return_note, false
	}

	return_note, status := GetToDoNote(note_id, user_id)
	if !status {
		log.Error("bad status")
		return return_note, false
	}

	err = db.Where("id = ?", note_id).Delete(&note).Error
	if err != nil {
		log.Error(err.Error())
		return return_note, false
	}

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return return_note, false
	}

	return return_note, true
}
