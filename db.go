package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Assuming you're using PostgreSQL
)

// Connect to database with error handling
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "your_connection_string")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, fmt.Errorf("could not connect to the database")
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %v", err)
	}

	return db, nil
}
