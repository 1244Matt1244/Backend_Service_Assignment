// db.go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

// ConnectDB - Initializes a single DB connection instance
func ConnectDB() (*sql.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %v", err)
	}

	dbInstance = db
	return db, nil
}
