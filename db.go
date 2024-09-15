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
	// Fetch the connection string from environment variables
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Fallback to default connection string if not set
		connStr = "postgresql://user:Prague1993@localhost:5432/myappdb?sslmode=disable"
	}

	var err error
	// Initialize the connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging the database: %v", err)
	}

	log.Println("Database connection established")
	return nil
}

// GetDB returns the database instance.
func GetDB() *sql.DB {
	if db == nil {
		log.Println("Database connection has not been initialized")
	}
	return db
}
