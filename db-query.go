package camera

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost dbname=postgres user=postgres password=Prague1993 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection successful")
}

// GetCamerasWithinRadius retrieves cameras within a specified radius
func GetCamerasWithinRadius(longitude, latitude, radius float64) ([]Camera, error) {
	query := `
        SELECT id, latitude, longitude 
        FROM traffic_cameras 
        WHERE ST_DWithin(
            ST_SetSRID(ST_MakePoint($1, $2), 4326),
            ST_SetSRID(ST_MakePoint(longitude, latitude), 4326),
            $3
        );
    `

	rows, err := db.Query(query, longitude, latitude, radius)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var cameras []Camera
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Latitude, &camera.Longitude); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		cameras = append(cameras, camera)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return cameras, nil
}

func main() {
	longitude := 15.99765
	latitude := 45.79791
	radius := 1000.0

	cameras, err := GetCamerasWithinRadius(longitude, latitude, radius)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cameras within radius:", cameras)
}
