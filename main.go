package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// Import the camera package
)

// Define Card and Camera structs only if needed in this file
type Card struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` // Ensure this field is present
}

func GetCardByID(id string, cards []Card) (Card, error) {
	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("card not found")
}

func GetCameraByID(id string, cameras []Camera) (Camera, error) {
	for _, camera := range cameras {
		if camera.ID == id {
			return camera, nil
		}
	}
	return Camera{}, fmt.Errorf("camera not found")
}

func ParseMTGCards(filename string) ([]Card, error) {
	return []Card{
		{ID: "1", Name: "Card 1", Type: "Type A"},
	}, nil
}

func SetupRouter(cards []Card, cameras []Camera) *mux.Router {
	r := mux.NewRouter()

	// MTG Card Routes
	r.HandleFunc("/cards/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		card, err := GetCardByID(id, cards)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(card); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}).Methods("GET")

	// Camera Routes
	r.HandleFunc("/cameras/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		camera, err := GetCameraByID(id, cameras)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(camera); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}
	}).Methods("GET")

	return r
}

func main() {
	cards, err := ParseMTGCards("mtg_cards.csv")
	if err != nil {
		log.Fatal(err)
	}

	cameras := []Camera{
		{ID: "1", Latitude: 40.7128, Longitude: -74.0060},
	}

	r := SetupRouter(cards, cameras)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
