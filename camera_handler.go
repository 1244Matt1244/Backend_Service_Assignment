package camera

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver import
)

// Camera struct should not be redeclared here, just use the one from camera.go

// FetchCameras queries the database to retrieve all cameras
func FetchCameras(conn *sql.DB) ([]Camera, error) {
	rows, err := conn.Query("SELECT id, latitude, longitude FROM traffic_cameras")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cameras []Camera
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Latitude, &camera.Longitude); err != nil {
			return nil, err
		}
		cameras = append(cameras, camera)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cameras, nil
}

// CameraHandler handles HTTP requests for cameras
func CameraHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=username dbname=yourdb sslmode=disable" // Example connection string
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	cameras, err := FetchCameras(dbConn)
	if err != nil {
		http.Error(w, "Could not fetch cameras", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cameras); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
