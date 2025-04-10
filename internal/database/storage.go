package database

import (
	"fmt"
	"todoflow-api/internal/config"
	"todoflow-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteDBInit(conf *config.Config) (*gorm.DB, error) {
	dsn := conf.SQLite_path
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite_db: %w", err)
	}

	db.AutoMigrate(&models.User{}, &models.ToDoList{})
	return db, nil
}
