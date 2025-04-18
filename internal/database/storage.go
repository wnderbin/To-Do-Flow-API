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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conf        *config.Config = config.MustLoad()
	log         slog.Logger    = *logger.Init(conf.Env)
	RDB         *redis.Client  = redis_db.Init(conf)
	db          *gorm.DB       = openPostgres()
	Postgres_db                = db
)

// redis

func CloseConRedis() {
	RDB.Close()
	log.Info("Closing connection with Redis...")
}

// sql

func ClosePostgres() error {
	sqldb, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}
	log.Info("Closing connection with Postgres...")
	return sqldb.Close()
}

func openPostgres() *gorm.DB {
	db, err := PostgresDBInit(conf)
	if err != nil {
		log.Error(err.Error())
		return db
	}
	return db
}

func PostgresDBInit(conf *config.Config) (*gorm.DB, error) {
	dsn := conf.Postgres.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

// Crud

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

// cRud

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

// crUd

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

// cruD

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
