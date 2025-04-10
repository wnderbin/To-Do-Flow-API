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

	db.AutoMigrate(&models.User{}, &models.ToDoList{})
	return db, nil
}

func SQLiteDBClose(db *gorm.DB) error {
	sqlite_db, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlite_db: %w", err)
	}

	err = sqlite_db.Close()
	if err != nil {
		return fmt.Errorf("failed to close sqlite_db: %w", err)
	}

	return nil
}

// Crud

func CreateUser(name string, username string, password string, email string) bool {
	var users []models.User
	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return false
	}

	db.Where("username = ?", username).Find(&users)
	if len(users) == 0 {

		db.Create(&models.User{
			Name:       name,
			Username:   username,
			Password:   password,
			Email:      email,
			Created_at: time.Now(),
		})
		if err = SQLiteDBClose(db); err != nil {
			log.Error(err.Error())
			return false
		}
		return true
	} else {
		SQLiteDBClose(db)
		return false
	}
}

func CreateToDo(ToDo_Note string, User_id int) bool {
	db, err := SQLiteDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return false
	}

	db.Create(&models.ToDoList{
		Created_at: time.Now(),
		Updated_at: time.Now(),
		ToDoNote:   ToDo_Note,
		User_id:    User_id,
	})

	if err = SQLiteDBClose(db); err != nil {
		log.Error(err.Error())
		return false
	}
	return true
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
	SQLiteDBClose(db)
	if err != nil {
		log.Error(err.Error())
		return user, false
	}

	return user, true
}

func GetToDoNotes(user_uuid_str string) ([]models.ToDoList, bool) {
	var notes []models.ToDoList

	ParsedUUID, err := uuid.Parse(user_uuid_str)
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
	SQLiteDBClose(db)
	if err != nil {
		log.Error(err.Error())
		return notes, false
	}

	return notes, false
}

func GetToDoNote(note_uuid_str string) (models.ToDoList, bool) {
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

	err = db.Where("id = ?", ParsedUUID).First(&note).Error
	SQLiteDBClose(db)
	if err != nil {
		log.Error(err.Error())
		return note, false
	}

	return note, true
}

// crUd

// cruD
