package migrator

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ApplySQLiteMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("init sqlite migrations: %w", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite",
		driver,
	)
	if err != nil {
		return fmt.Errorf("create migrations instance: %w", err)
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
