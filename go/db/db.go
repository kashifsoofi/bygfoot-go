package db

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dbPath string) (*sql.DB, error) {
	sqliteDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite db %w", err)
	}

	return sqliteDb, nil
}

//go:embed migrations/*.sql
var migrations embed.FS

func RunMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("creating sqlite3 db driver failed %w", err)
	}

	// https://github.com/golang-migrate/migrate/issues/471
	// temp solution would update to use embed source once merged
	source, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return fmt.Errorf("error initialising source %w", err)
	}

	m, err := migrate.NewWithInstance("migrations", source, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("initialising db migration failed %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("database migration failed %w", err)
	}

	return nil
}
