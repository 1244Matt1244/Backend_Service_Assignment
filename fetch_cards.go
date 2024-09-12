// fetch_cards.go
package mtg

import (
	"encoding/json"
	"fmt"
	"log"
	"my_project/models" // C:/Users/korisnik/Desktop/MTG_Backend_Service_Assignment
	"net/http"
)

// FetchCardsFromAPI fetches cards from the given API URL and processes them
func FetchCardsFromAPI(url string) ([]models.Card, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Cards []models.Card `json:"cards"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	for _, card := range result.Cards {
		log.Printf("Card ID: %s, Name: %s", card.ID, card.Name)
	}

	return result.Cards, nil
}
