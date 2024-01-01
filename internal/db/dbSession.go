package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

const (
	DB_USER_NAME = "DB_USER_NAME"
	DB_PASSWORD  = "DB_PASSWORD"
	DB_URL       = "DB_URL"
	DB_PORT      = "DB_PORT"
	DB_NAME      = "DB_NAME"
	DRIVER_NAME  = "postgres"
)

var db *sql.DB
var url string
var migrations string

func init() {
	url = dbUrl()
	db = initDb()
	migrations = migrationFilesDest()
	migrateDb()
}

func migrateDb() {
	m, err := migrate.New(migrations, url)
	if err == nil {
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
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
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
	foreignCurrencyCode := values["foreignCurrencyCode"]
	_, err := db.Exec("UPDATE ussd_session SET next_stage = $1, country_name = $2, amount = $3, foreign_currency_code =$4 WHERE session_id = $5",
		nextStage, countryName, amount, foreignCurrencyCode, sessionId)
	if err != nil {
	}
}

func FindSession(sessionId string) (*UssdSession, error) {
	rows, err := db.Query("SELECT * FROM ussd_session WHERE session_id = $1", sessionId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var s = UssdSession{}
	var foundRows bool

	for rows.Next() {
		foundRows = true
		rows.Scan(
			&s.SessionId,
			&s.Msisdn,
			&s.NextStage,
			&s.CountryName,
			&s.Amount,
			&s.ForeignCurrencyCode,
			&s.SessionStartTime)
	}

	if !foundRows {
		return nil, nil
	}

	return &s, nil
}

func DeleteSession(sessionId string) {
	db.Exec("DELETE FROM ussd_session WHERE session_id = $1", sessionId)
}

func migrationFilesDest() string {
	basePath, err := os.Getwd()
	if err != nil {
	}
	return fmt.Sprintf("file://%s/internal/db/migration", basePath)
}

type UssdSession struct {
	SessionId           string
	Msisdn              string
	NextStage           string
	CountryName         string
	Amount              float32
	ForeignCurrencyCode string
	SessionStartTime    time.Time
}
