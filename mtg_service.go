package mtg

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Card struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	ManaCost    string `json:"mana_cost"`
	Description string `json:"description"`
}

func ParseMTGCards(filename string) ([]Card, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cards []Card
	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		cards = append(cards, Card{
			ID:          record[0],
			Name:        record[1],
			ManaCost:    record[2],
			Type:        record[3],
			Description: record[4],
		})
	}

	return cards, nil
}

func GetCardByID(id string, cards []Card) (Card, error) {
	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("card not found")
}
