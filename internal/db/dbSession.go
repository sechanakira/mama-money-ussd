package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"time"
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

func InitUssdSession(sessionId string, msisdn string) {
	_, err := db.Exec("INSERT INTO ussd_session (session_id,msisdn,next_stage,session_start_time) VALUES ($1,$2,$3,$4)",
		sessionId, msisdn, "MENU_2", time.Now())
	if err != nil {
	}
}

func UpdateUssdSession(sessionId string, values map[string]string) {
	nextStage := values["nextStage"]
	countryName := values["countryName"]
	amount := values["amount"]
	foreignCurrencyCode := values[""]
	_, err := db.Exec("UPDATE ussd_session SET next_stage = $1, country_name = $2, amount = $3, foreign_currency_code =$4 WHERE session_id = $5",
		nextStage, countryName, amount, foreignCurrencyCode, sessionId)
	if err != nil {
	}
}
