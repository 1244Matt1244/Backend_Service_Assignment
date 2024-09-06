// camera_handler.go
import (
	"database/sql"
	"fmt"
	"myapp/internal/db"
)

// Define the Camera struct to represent a camera entity.
type Camera struct {
	ID       int    // Assuming ID is an integer, adjust the type as needed.
	Location string // Assuming Location is a string, adjust the type as needed.
	URL      string // Assuming URL is a string, adjust the type as needed.
}

func FetchCameras() ([]Camera, error) {
	// Get the database connection.
	conn, err := db.GetDBConnection()
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %v", err)
	}
	defer conn.Close() // Close the connection after the function returns.

	// Query the database for all cameras.
	rows, err := conn.Query("SELECT * FROM traffic_cameras")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close the result set after the function returns.

	// Create a slice to hold the cameras.
	var cameras []Camera

	// Scan the rows and append to the cameras slice.
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Location, &camera.URL); err != nil {
			return nil, err
		}
		cameras = append(cameras, camera)
	}

	// Check for any errors that occurred during the iteration.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cameras, nil
}

