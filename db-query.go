package camera

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// ConnectDB initializes the database connection
func ConnectDB() error {
	var err error
	db, err = sql.Open("postgres", "host=localhost dbname=postgres user=postgres password=Prague1993 sslmode=disable")
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	fmt.Println("Database connection successful")
	return nil
}

// GetCamerasWithinRadius retrieves cameras within a specified radius
func GetCamerasWithinRadius(longitude, latitude, radius float64) ([]Camera, error) {
	// Note: The radius in ST_DWithin should be converted to degrees if it's in meters or kilometers
	query := `
        SELECT id, latitude, longitude 
        FROM traffic_cameras 
        WHERE ST_DWithin(
            ST_SetSRID(ST_MakePoint($1, $2), 4326),
            ST_SetSRID(ST_MakePoint(longitude, latitude), 4326),
            $3 / 1000  -- Assuming radius is in meters, convert to kilometers
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
	// Initialize database connection
	if err := ConnectDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Example coordinates and radius (in meters)
	longitude := 15.99765
	latitude := 45.79791
	radius := 1000.0 // 1000 meters = 1 km

	// Query cameras within the radius
	cameras, err := GetCamerasWithinRadius(longitude, latitude, radius)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cameras within radius:", cameras)
}
