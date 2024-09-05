// mtg_service.go
package mtg

import (
	"fmt"
	"log"
	"net/http"
)

// ImportMTGCards fetches MTG card data and inserts it into the database
func ImportMTGCards() {
	cards, err := FetchMTGCards()
	if err != nil {
		log.Fatalf("Error fetching MTG cards: %v", err)
	}

	for _, card := range cards {
		if err := InsertMTGCard(card); err != nil {
			log.Printf("Error inserting card %s: %v", card.Name, err)
		}
	}
}

// FetchCardsFromAPI fetches cards from the MTG API and stores them in the database
func FetchCardsFromAPI() error {
	resp, err := http.Get("https://api.magicthegathering.io/v1/cards")
	if err != nil {
		log.Printf("Error while calling MTG API: %v", err)
		return fmt.Errorf("failed to fetch cards from MTG API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response from MTG API: %d", resp.StatusCode)
	}

	// Process the response (store data in DB)
	// Handle DB errors
	err = StoreCardsInDB(cards)
	if err != nil {
		log.Printf("Database error: %v", err)
		return fmt.Errorf("failed to store cards in the database")
	}

	return nil
}

// InsertMTGCard inserts an MTG card into the database (dummy implementation)
func InsertMTGCard(card MTGCard) error {
	_, err := db.Exec("INSERT INTO mtg_cards (id, name, mana_cost, type, description) VALUES ($1, $2, $3, $4, $5)",
		card.ID, card.Name, card.ManaCost, card.Type, card.Description)
	return err
}
