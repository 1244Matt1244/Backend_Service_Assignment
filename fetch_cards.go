package mtg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// FetchCardsFromAPI fetches MTG cards from the API and stores them in the database
func FetchCardsFromAPI(apiURL string, db *sqlx.DB) error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("error fetching cards from API: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Cards []models.Card `json:"cards"` // Ensure 'models.Card' is the correct type
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("error decoding API response: %v", err)
	}

	for _, card := range data.Cards {
		_, err := db.NamedExec(`INSERT INTO mtg_cards (id, name, colors, cmc, type, subtype, rarity, image_url, original_text) 
		VALUES (:id, :name, :colors, :cmc, :type, :subtype, :rarity, :image_url, :original_text)
		ON CONFLICT (id) DO NOTHING`, card)
		if err != nil {
			return fmt.Errorf("error inserting card into database: %v", err)
		}
	}

	return nil
}
