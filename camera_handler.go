package main

import (
	"encoding/json"
	"net/http"
)

// ListCameras handles the HTTP request to list cameras
func ListCameras(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, location, url FROM traffic_cameras")
	if err != nil {
		http.Error(w, "Error fetching traffic cameras", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cameras []Camera
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Location, &camera.URL); err != nil {
			http.Error(w, "Error scanning camera", http.StatusInternalServerError)
			return
		}
		cameras = append(cameras, camera)
	}
	json.NewEncoder(w).Encode(cameras)
}
