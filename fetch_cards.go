package mtg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchCardsFromAPI(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Cards []Card `json:"cards"` // Use the imported Card struct
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	for _, card := range result.Cards {
		fmt.Printf("Card ID: %s, Name: %s\n", card.ID, card.Name)
	}

	return nil
}
