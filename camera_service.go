package camera

import (
	"encoding/csv"
	"fmt"
	"os"
	"postgres/sql"
	"strconv"
)

type Camera struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ParseCameras reads a CSV file and returns a slice of Camera structs
func ParseCameras(filename string) ([]Camera, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var cameras []Camera
	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header
	if err != nil {
		return nil, fmt.Errorf("error reading header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading record: %w", err)
		}
		if len(record) < 3 {
			return nil, fmt.Errorf("invalid record length: %v", record)
		}
		latitude, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing latitude: %w", err)
		}
		longitude, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing longitude: %w", err)
		}
		cameras = append(cameras, Camera{
			ID:        record[0],
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return cameras, nil
}

// GetCameraByID retrieves a camera by its ID from a slice of cameras
func GetCameraByID(id string) (Camera, error) {
	for _, camera := range cameras {
		if camera.ID == id {
			return camera, nil
		}
	}
	return Camera{}, fmt.Errorf("camera with ID %s not found", id)
}

// InsertCameras inserts the parsed cameras into the database
func InsertCameras(db *sql.DB, cameras []Camera) error {
	for _, cam := range cameras {
		_, err := db.Exec(`INSERT INTO cameras (id, latitude, longitude) 
        VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING`,
			cam.ID, cam.Latitude, cam.Longitude)
		if err != nil {
			return fmt.Errorf("failed to insert camera %s: %v", cam.ID, err)
		}
	}
	return nil
}
