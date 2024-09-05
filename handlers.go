// handlers.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
)

// ImportMTGCardsHandler handles the request to import MTG cards
func ImportMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Unable to read CSV file", http.StatusInternalServerError)
		return
	}

	for _, record := range records {
		if len(record) == 5 {
			card := MTGCard{
				ID:          record[0],
				Name:        record[1],
				ManaCost:    record[2],
				Type:        record[3],
				Description: record[4],
			}
			log.Printf("Processing card: %+v", card)
			// Process card (e.g., save to database)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Import successful"))
}

// ImportCamerasHandler handles the request to import traffic cameras
func ImportCamerasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var cameras []Camera
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cameras)
	if err != nil {
		http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
		return
	}

	for _, camera := range cameras {
		log.Printf("Processing camera: ID=%d, Location=%s, URL=%s",
			camera.ID, camera.Location, camera.URL)
		// Process camera (e.g., save to database)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cameras import successful"))
}
