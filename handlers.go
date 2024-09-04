package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListMTGCardsHandler handles requests to list all MTG cards
func ListMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	var cards []MTGCard
	rows, err := db.Query("SELECT * FROM mtg_cards")
	if err != nil {
		http.Error(w, "Error fetching cards from the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var card MTGCard
		if err := rows.Scan(&card.ID, &card.Name, &card.ManaCost, &card.Type, &card.Description); err != nil {
			http.Error(w, "Error scanning card from the database", http.StatusInternalServerError)
			return
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error processing rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

// GetMTGCardHandler handles requests to get a specific MTG card by ID
func GetMTGCardHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	cardID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return
	}

	var card MTGCard
	err = db.QueryRow("SELECT * FROM mtg_cards WHERE id = $1", cardID).Scan(&card.ID, &card.Name, &card.ManaCost, &card.Type, &card.Description)
	if err != nil {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

// ImportMTGCardsHandler handles requests to import MTG cards
func ImportMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	ImportMTGCards()
	w.Write([]byte("MTG cards imported successfully"))
}

// ListCamerasInRadiusHandler handles requests to list traffic cameras within a given radius
func ListCamerasInRadiusHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	latitudeParam := r.URL.Query().Get("latitude")
	longitudeParam := r.URL.Query().Get("longitude")
	radiusParam := r.URL.Query().Get("radius")

	latitude, err := strconv.ParseFloat(latitudeParam, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeParam, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radiusParam, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	var cameras []Camera
	query := `
        SELECT id, location, url
        FROM traffic_cameras
        WHERE earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(latitude, longitude);
    `

	rows, err := db.Query(query, latitude, longitude, radius*1000) // radius is converted to meters
	if err != nil {
		http.Error(w, "Error fetching cameras from the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Location, &camera.URL); err != nil {
			http.Error(w, "Error scanning camera from the database", http.StatusInternalServerError)
			return
		}
		cameras = append(cameras, camera)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error processing rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cameras)
}

// RegisterHandlers sets up the routes and handlers
func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/list", ListCamerasInRadiusHandler).Methods("GET")
	router.HandleFunc("/import", ImportMTGCardsHandler).Methods("POST")
	router.HandleFunc("/card/{id}", GetMTGCardHandler).Methods("GET")
}
