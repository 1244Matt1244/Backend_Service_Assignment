package main

import (
	"database/sql"

	_ "github.com/lib/pq" // Assuming you're using PostgreSQL
)

// Database connection variable
var db *sql.DB

// InitializeDB sets up the database connection
func InitializeDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}
