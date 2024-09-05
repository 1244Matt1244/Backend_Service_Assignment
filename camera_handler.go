// camera_handler.go
package camera

import (
	"fmt"
	"myapp/db" // Import your db package
)

func FetchCameras() ([]Camera, error) {
	conn, err := db.GetDBConnection() // Reuse the singleton connection
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %v", err)
	}

	rows, err := conn.Query("SELECT * FROM traffic_cameras")
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

	return cameras, nil
}
