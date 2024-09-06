// db.go
package db

import (
	"database/sql"
)

var dbInstance *sql.DB

func GetDBConnection() (*sql.DB, error) {
	if dbInstance == nil {
		connStr := "user=postgres dbname=postgres sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}
		dbInstance = db
	}
	return dbInstance, nil
}
