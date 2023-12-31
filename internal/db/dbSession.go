package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
)

const (
	DB_USER_NAME         = "DB_USER_NAME"
	DB_PASSWORD          = "DB_PASSWORD"
	DB_URL               = "DB_URL"
	DB_PORT              = "DB_PORT"
	DB_NAME              = "DB_NAME"
	DRIVER_NAME          = "postgres"
	MIGRATION_FILES_DEST = "file:///migration"
)

var db *sql.DB
var url string

func init() {
	url = dbUrl()
	db = initDb()
	migrateDb()
}

func migrateDb() {
	m, err := migrate.New(MIGRATION_FILES_DEST, url)
	if err != nil {
		m.Up()
	} else {
		m.Down()
	}
}

func initDb() *sql.DB {
	db, err := sql.Open(DRIVER_NAME, url)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv(DB_USER_NAME),
		os.Getenv(DB_PASSWORD),
		os.Getenv(DB_URL),
		os.Getenv(DB_PORT),
		os.Getenv(DB_NAME))
}
