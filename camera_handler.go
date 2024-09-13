package camera

import (
	"camera"
	"encoding/json"
	"net/http"
)

// ListCamerasHandler handles the request to list all cameras
func ListCamerasHandler(w http.ResponseWriter, r *http.Request) {
	cameras, err := camera.FetchCameras()
	if err != nil {
		http.Error(w, "Failed to fetch cameras", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cameras); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// GetCameraByIDHandler handles the request to get a single camera by ID
func GetCameraByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	camera, err := camera.GetCameraByID(id)
	if err != nil {
		http.Error(w, "Camera not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(camera); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// AddCameraHandler handles the request to add a new camera
func AddCameraHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to add a new camera
}

// DeleteCameraHandler handles the request to delete a camera by ID
func DeleteCameraHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := camera.DeleteCameraByID(id)
	if err != nil {
		http.Error(w, "Failed to delete camera", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
