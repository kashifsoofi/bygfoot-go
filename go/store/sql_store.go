package store

import (
	"database/sql"
	"embed"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

const exitDbOpen = 101

//go:embed migrations/*.sql
var migrations embed.FS

type SqlStore struct {
	dbx    *sqlx.DB
	db     *sql.DB
	region RegionStore
}

func initConnection(driverName, dataSource string) *SqlStore {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.WithFields(log.Fields{
			"driverName": driverName,
			"dataSource": dataSource,
			"Error":      err.Error(),
		}).Fatal("failed to open db")
		os.Exit(exitDbOpen)
	}

	sqlStore := &SqlStore{
		dbx: sqlx.NewDb(db, driverName),
		db:  db,
	}
	sqlStore.dbx.MapperFunc(func(s string) string { return s })

	return sqlStore
}

func NewSqlStore(driverName, dataSource string) Store {
	sqlStore := initConnection(driverName, dataSource)

	sqlStore.region = NewSqlRegionStore(sqlStore)

	runMigrations(sqlStore.db, driverName)

	return sqlStore
}

func (ss *SqlStore) Close() {
	log.Info("closing db")
	ss.dbx.Close()
}

func (ss *SqlStore) Region() RegionStore {
	return ss.region
}

func runMigrations(db *sql.DB, driverName string) {
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("failed to create migration driver")
		os.Exit(1)
	}

	// https://github.com/golang-migrate/migrate/issues/471
	// temp solution would update to use embed source once merged
	srcDriver, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("failed to create source driver")
		os.Exit(1)
	}

	m, err := migrate.NewWithInstance("migrations", srcDriver, driverName, dbDriver)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("failed to create migrate instance")
		os.Exit(1)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("failed to run database migration")
		os.Exit(1)
	}
}
