package camera

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver import
)

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
func handleCameraInsert(w http.ResponseWriter, r *http.Request) {
	err := InsertCameras(db, cameras)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting cameras: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cameras successfully inserted"))
}
