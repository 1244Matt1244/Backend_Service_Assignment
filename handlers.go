package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// ListMTGCardsHandler handles the request to list all MTG cards
func ListMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	// Example: Fetch cards from a service and return as JSON
	cards, err := fetchAllMTGCards() // Replace with actual fetching logic
	if err != nil {
		http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// GetMTGCardHandler handles the request to get a single MTG card by ID
func GetMTGCardHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	card, err := getMTGCardByID(id) // Replace with actual fetching logic
	if err != nil {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(card); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// ImportMTGCardsHandler handles the request to import MTG cards
func ImportMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement file upload logic or similar import functionality
}

// ListCamerasHandler handles the request to list all cameras
func ListCamerasHandler(w http.ResponseWriter, r *http.Request) {
	cameras, err := fetchAllCameras() // Replace with actual fetching logic
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
	camera, err := getCameraByID(id) // Replace with actual fetching logic
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
	err := deleteCameraByID(id) // Replace with actual deletion logic
	if err != nil {
		http.Error(w, "Failed to delete camera", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // No content to return on successful deletion
}

// ListCamerasInRadiusHandler handles the request to list cameras within a certain radius
func ListCamerasInRadiusHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to list cameras within a certain radius
}

// ListCards handles pagination and filters
func ListCards(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	color := r.URL.Query().Get("color")
	if color != "" && !isValidColor(color) {
		http.Error(w, "Invalid color", http.StatusBadRequest)
		return
	}

	// Fetch and return the paginated and filtered cards
	cards, err := fetchFilteredMTGCards(page, color) // Replace with actual fetching logic
	if err != nil {
		http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func isValidColor(color string) bool {
	validColors := map[string]bool{"white": true, "black": true, "blue": true, "red": true, "green": true}
	return validColors[color]
}
