package database

import (
	"fmt"
	"log/slog"
	"time"
	"todoflow-api/internal/config"
	"todoflow-api/internal/logger"
	"todoflow-api/internal/redis_db"

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

func CloseConRedis() {
	RDB.Close()
	log.Info("Closing connection with Redis...")
}

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
	if conf.Postgres.Status == 1 {
		return nil, nil
	}
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
