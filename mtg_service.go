package mtg

import (
	"encoding/json"
	"fmt"
	"models"
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

	// Define a structure to parse the API response
	var data struct {
		Cards []models.MTGCard `json:"cards"` // Ensure 'models.MTGCard' is the correct type
	}

	// Decode the JSON response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("error decoding API response: %v", err)
	}

	// Start a transaction for batch insert
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Insert each card into the database within the transaction
	for _, card := range data.Cards {
		_, err := tx.NamedExec(`INSERT INTO mtg_cards (id, name, colors, cmc, type, subtype, rarity, image_url, original_text) 
		VALUES (:id, :name, :colors, :cmc, :type, :subtype, :rarity, :image_url, :original_text)
		ON CONFLICT (id) DO NOTHING`, card)
		if err != nil {
			tx.Rollback() // Rollback the transaction on error
			return fmt.Errorf("error inserting card %s into database: %v", card.ID, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
