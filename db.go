// db.go
package db

import (
	"database/sql"
	"fmt"
)

var dbInstance *sql.DB

// GetDBConnection returns a singleton instance of the database connection
func GetDBConnection() (*sql.DB, error) {
	if dbInstance == nil {
		connStr := "user=your_user dbname=your_db sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to db: %v", err)
		}
		dbInstance = db
	}
	return dbInstance, nil
}
