package database

import (
	"fmt"
	"log/slog"
	"time"

	"todoflow-api/internal/config"
	"todoflow-api/internal/logger"
	"todoflow-api/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	conf *config.Config = config.MustLoad()
	log  slog.Logger    = *logger.Init(conf.Env)
)

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
	//var users []models.User
	var user models.User

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	/*err = db.Where("username = ?", username).Find(&users).Error
	if err != nil {
		log.Error(err.Error())
	}

	if len(users) == 0 {*/

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

func GetUser(uuid_str string) (models.User, bool) {
	var user models.User

	ParsedUUID, err := uuid.Parse(uuid_str)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	err = db.Where("id = ?", ParsedUUID).First(&user).Error
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

	ParsedUUID, err := uuid.Parse(note_uuid_str)
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	err = db.Where("id = ?", ParsedUUID).Find(&note).Error
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

// cruD
