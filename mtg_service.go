package mtg

import (
	"encoding/csv"
	"fmt"
	"my_project/models" // C:/Users/korisnik/Desktop/MTG_Backend_Service_Assignment
	"os"
	"strconv"
)

// ParseMTGCards parses MTG cards from a CSV file
func ParseMTGCards(filename string) ([]models.MTGCard, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cards []models.MTGCard
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
		cards = append(cards, models.MTGCard{
			ID:           record[0],
			Name:         record[1],
			Colors:       []string{record[2]},
			CMC:          atoi(record[3]),
			Type:         []string{record[4]},
			Subtype:      []string{record[5]},
			Rarity:       record[6],
			OriginalText: record[7],
			ImageURL:     record[8],
		})
	}

	return cards, nil
}

// GetCardByID retrieves a card by its ID
func GetCardByID(id string, cards []models.MTGCard) (models.MTGCard, error) {
	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}
	return models.MTGCard{}, fmt.Errorf("card not found")
}

// Helper function to convert string to int
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
