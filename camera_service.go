package camera

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// Camera struct represents a traffic camera with its id, name, latitude, and longitude
type Camera struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InsertCamerasFromCSV reads a CSV file and inserts camera data into the database
func InsertCamerasFromCSV(db *sql.DB, filepath string) error {
	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the CSV header and discard it (skip the first line)
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Read the remaining rows from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	// Loop over each record in the CSV and insert it into the database
	for _, record := range records {
		name := record[0]
		lat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Invalid latitude value for camera %s: %v", name, err)
			continue
		}
		lon, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Printf("Invalid longitude value for camera %s: %v", name, err)
			continue
		}

		// Prepare the SQL query for inserting the camera
		query := `INSERT INTO traffic_cameras (name, latitude, longitude, location)
				  VALUES ($1, $2, $3, ST_SetSRID(ST_MakePoint($3, $2), 4326))`

		_, err = db.Exec(query, name, lat, lon)
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}
	log.Println("Cameras inserted successfully.")
	return nil
}

// FindCamerasWithinRadius retrieves cameras within a specified radius from a given point
func FindCamerasWithinRadius(db *sql.DB, lat, lon, radius float64) ([]Camera, error) {
	query := `SELECT id, name, latitude, longitude 
			  FROM traffic_cameras 
			  WHERE ST_DWithin(
				ST_SetSRID(ST_MakePoint($1, $2), 4326),
				ST_SetSRID(ST_MakePoint(longitude, latitude), 4326), $3)`

	rows, err := db.Query(query, lon, lat, radius)
	if err != nil {
		return nil, fmt.Errorf("error querying cameras: %w", err)
	}
	defer rows.Close()

	var cameras []Camera
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Name, &camera.Latitude, &camera.Longitude); err != nil {
			return nil, fmt.Errorf("error scanning camera: %w", err)
		}
		cameras = append(cameras, camera)
	}

	if cameras == nil {
		return []Camera{}, nil
	}

	return cameras, nil
}

// GetCameraByID retrieves a camera from the database by its ID
func GetCameraByID(db *sql.DB, id string) (*Camera, error) {
	query := `SELECT id, name, latitude, longitude FROM traffic_cameras WHERE id = $1`
	row := db.QueryRow(query, id)

	var camera Camera
	err := row.Scan(&camera.ID, &camera.Name, &camera.Latitude, &camera.Longitude)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("camera not found")
	} else if err != nil {
		return nil, fmt.Errorf("error fetching camera: %w", err)
	}

	return &camera, nil
}
