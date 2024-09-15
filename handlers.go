package main

import (
	"encoding/json"
	"mtg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListMTGCardsHandler handles the request to list all MTG cards
func ListMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	cards, err := mtg.FetchAllMTGCards()
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
	vars := mux.Vars(r)
	id := vars["id"]

	card, err := mtg.GetMTGCardByID(id)
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
	apiURL := "https://api.magicthegathering.io/v1/cards"
	err := mtg.FetchCardsFromAPI(apiURL, mtg.GetDB())
	if err != nil {
		http.Error(w, "Failed to import cards", http.StatusInternalServerError)
		return
	}
	JSONResponse(w, map[string]string{"status": "import successful"}, http.StatusOK)
}

// ListCamerasHandler handles the request to list all cameras
func ListCamerasHandler(w http.ResponseWriter, r *http.Request) {
	cameras, err := mtg.FetchAllCameras()
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
	vars := mux.Vars(r)
	id := vars["id"]

	camera, err := mtg.GetCameraByID(id)
	if err != nil {
		http.Error(w, "Camera not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(camera); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// DeleteCameraHandler handles the request to delete a camera by ID
func DeleteCameraHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := mtg.DeleteCameraByID(id)
	if err != nil {
		http.Error(w, "Failed to delete camera", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListCamerasInRadiusHandler handles the request to list cameras within a certain radius
func ListCamerasInRadiusHandler(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")
	radiusStr := r.URL.Query().Get("radius")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	cameras, err := mtg.FetchCamerasInRadius(lat, lon, radius)
	if err != nil {
		http.Error(w, "Failed to fetch cameras", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cameras); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// ListCards handles pagination and filters for MTG cards
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

	cards, err := mtg.FetchFilteredMTGCards(page, color)
	if err != nil {
		http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// isValidColor checks if the color is valid for MTG cards
func isValidColor(color string) bool {
	validColors := map[string]bool{
		"white": true, "black": true, "blue": true, "red": true, "green": true,
	}
	return validColors[color]
}

// JSONResponse is a helper function to send JSON responses
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
