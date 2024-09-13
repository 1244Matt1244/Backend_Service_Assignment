package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Importing the PostgreSQL driver
)

var db *sql.DB

// ConnectDB initializes the database connection.
func ConnectDB() error {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgresql://user:Prague1993@localhost:5432/myappdb?sslmode=disable"
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Database connection established")
	return nil
}

// GetDB returns the database instance.
func GetDB() *sql.DB {
	return db
}
