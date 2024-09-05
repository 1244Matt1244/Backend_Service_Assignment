package mtg

import (
	"encoding/csv"
	"log"
	"net/http"
)

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
