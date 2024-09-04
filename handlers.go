// handlers.go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListMTGCardsHandler handles requests to list all MTG cards
func ListMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	var cards []MTGCard
	rows, err := db.Query("SELECT * FROM mtg_cards")
	if err != nil {
		http.Error(w, "Error fetching cards", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var card MTGCard
		if err := rows.Scan(&card.ID, &card.Name, &card.ManaCost, &card.Type, &card.Description); err != nil {
			http.Error(w, "Error scanning card", http.StatusInternalServerError)
			return
		}
		cards = append(cards, card)
	}

	json.NewEncoder(w).Encode(cards)
}

// GetMTGCardHandler handles requests to get a specific MTG card by ID
func GetMTGCardHandler(w http.ResponseWriter, r *http.Request) {
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

	json.NewEncoder(w).Encode(card)
}

// ImportMTGCardsHandler handles requests to import MTG cards
func ImportMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	ImportMTGCards()
	w.Write([]byte("MTG cards imported successfully"))
}

// ListCamerasHandler handles requests to list traffic cameras
func ListCamerasHandler(w http.ResponseWriter, r *http.Request) {
	var cameras []Camera
	rows, err := db.Query("SELECT * FROM traffic_cameras")
	if err != nil {
		http.Error(w, "Error fetching cameras", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

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
