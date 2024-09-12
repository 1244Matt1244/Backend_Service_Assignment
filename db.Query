package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Assume db is already connected and initialized
var db *sql.DB

func GetCamerasWithinRadius(longitude, latitude float64, radius float64) ([]Camera, error) {
	query := `
    SELECT id, location, url
    FROM traffic_cameras
    WHERE ST_DWithin(
        ST_SetSRID(ST_MakePoint($1, $2), 4326),
        ST_SetSRID(ST_MakePoint(longitude, latitude), 4326),
        $3
    );
    `

	rows, err := db.Query(query, longitude, latitude, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cameras []Camera
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Location, &camera.URL); err != nil {
			return nil, err
		}
		cameras = append(cameras, camera)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cameras, nil
}

func main() {
	// Initialize your database connection here
	// db, err := sql.Open("postgres", "your connection string")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	longitude := 15.99765
	latitude := 45.79791
	radius := 1000.0 // Radius in meters

	cameras, err := GetCamerasWithinRadius(longitude, latitude, radius)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cameras within radius:", cameras)
}
