package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() error {
	connStr := "postgresql://user:password@db:5432/myappdb?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}
