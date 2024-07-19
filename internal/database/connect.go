package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time-tracker/config"
)

var DB *sql.DB

func ConnectDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config("POSTGRES_HOST"),
		config.Config("POSTGRES_PORT"),
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_DB"),
	)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
