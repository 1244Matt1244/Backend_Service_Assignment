package mtg

import "fmt"

type Card struct {
	ID   string
	Name string
	Type string // Include the 'Type' field here
}

func GetCardByID(id string, cards []Card) (Card, error) {
	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("card not found")
}
