package handlers

import (
	"Backend_Service_Assignment/cmd/camera"
	"Backend_Service_Assignment/cmd/mtg"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HealthCheck returns a simple message indicating server status
// @Summary Health Check
// @Description Check if the server is running
// @Tags health
// @Produce plain
// @Success 200 {string} string "Server is running"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

// CameraHandler handles the request to get a single camera by ID
// @Summary Get Camera by ID
// @Description Retrieve a specific camera by its ID
// @Tags cameras
// @Accept json
// @Produce json
// @Param id path string true "Camera ID"
// @Success 200 {object} camera.Camera
// @Failure 404 {string} string "Camera not found"
// @Router /cameras/{id} [get]
func CameraHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		cameraData, err := camera.GetCameraByID(db, id)
		if err != nil {
			http.Error(w, "Camera not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cameraData); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}
}

// ListCamerasHandler handles the request to list cameras within a specific radius
// @Summary List cameras by radius
// @Description List cameras within a specified radius of a given location
// @Tags cameras
// @Accept json
// @Produce json
// @Param latitude query number true "Latitude"
// @Param longitude query number true "Longitude"
// @Param radius query number true "Radius in meters"
// @Success 200 {array} camera.Camera
// @Failure 400 {string} string "Invalid parameters"
// @Failure 500 {string} string "Error retrieving cameras"
// @Router /cameras [get]
func ListCamerasHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		cameras, err := camera.FindCamerasWithinRadius(db, lat, lon, radius)
		if err != nil {
			http.Error(w, "Failed to fetch cameras", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cameras); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}
}

// MTGCardHandler handles the request to get a single MTG card by ID
// @Summary Get MTG card by ID
// @Description Retrieve a specific MTG card by its ID
// @Tags mtg
// @Accept json
// @Produce json
// @Param id path string true "MTG Card ID"
// @Success 200 {object} mtg.MTGCard
// @Failure 404 {string} string "Card not found"
// @Router /cards/{id} [get]
func MTGCardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		card, err := mtg.GetCardByID(db, id)
		if err != nil {
			http.Error(w, "Card not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(card); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}
}

// ListMTGCardsHandler handles the request to list all MTG cards with optional filters
// @Summary List MTG cards
// @Description List MTG cards with optional filters and pagination
// @Tags mtg
// @Accept json
// @Produce json
// @Param color query string false "Card Color"
// @Param rarity query string false "Card Rarity"
// @Param type query string false "Card Type"
// @Param name query string false "Card Name"
// @Param page query int false "Page number"
// @Success 200 {array} mtg.MTGCard
// @Failure 500 {string} string "Failed to fetch cards"
// @Router /list-cards [get]
func ListMTGCardsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := r.URL.Query().Get("color")
		rarity := r.URL.Query().Get("rarity")
		cardType := r.URL.Query().Get("type")
		name := r.URL.Query().Get("name")

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}

		filters := mtg.CardFilters{
			Color:  color,
			Rarity: rarity,
			Type:   cardType,
			Name:   name,
		}

		cards, total, currentPage, err := mtg.SearchCards(db, filters, page)
		if err != nil {
			http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"total": total,
			"page":  currentPage,
			"items": cards,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}
}

// ImportMTGCardsHandler imports cards from the MTG API into the database
// @Summary Import MTG cards from API
// @Description Imports 500 MTG cards into the PostgreSQL database
// @Tags mtg
// @Produce json
// @Success 200 {string} string "Cards imported successfully"
// @Failure 500 {string} string "Error importing cards"
// @Router /import-cards [post]
func ImportMTGCardsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mtg.ImportCardsFromAPI(db)
		if err != nil {
			http.Error(w, "Error importing cards", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cards imported successfully"))
	}
}
