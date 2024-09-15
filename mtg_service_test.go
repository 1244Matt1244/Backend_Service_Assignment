package mtg

import (
	"os"
	"testing"

	"github.com/1244Matt1244/Backend_Service_Assignment/internal/models"
)

func TestParseMTGCards(t *testing.T) {
	// Create a temporary file with sample data for testing
	tmpfile, err := os.CreateTemp("", "test_cards_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Write sample CSV data
	data := `id,name,colors,cmc,type,subtype,rarity,originalText,imageUrl
1,Card One,"red",1,Creature,Human,Common,A powerful creature,http://example.com/image1
2,Card Two,"green",2,Enchantment,Forest,Uncommon,An enchantment card,http://example.com/image2`
	_, err = tmpfile.Write([]byte(data))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Parse the cards
	cards, err := ParseMTGCards(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to parse MTG cards: %v", err)
	}

	// Check if the correct number of cards were parsed
	if len(cards) != 2 {
		t.Errorf("Expected 2 cards, got %d", len(cards))
	}

	// Verify the data of the first card
	if cards[0].ID != "1" || cards[0].Name != "Card One" {
		t.Errorf("Unexpected card data: %+v", cards[0])
	}
}

func TestGetCardByID(t *testing.T) {
	// Create a sample list of cards
	cards := []models.MTGCard{
		{ID: "1", Name: "Card One", Colors: []string{"red"}, CMC: 1, Type: []string{"Creature"}, Subtype: []string{"Human"}, Rarity: "Common", OriginalText: "A powerful creature", ImageURL: "http://example.com/image1"},
		{ID: "2", Name: "Card Two", Colors: []string{"green"}, CMC: 2, Type: []string{"Enchantment"}, Subtype: []string{"Forest"}, Rarity: "Uncommon", OriginalText: "An enchantment card", ImageURL: "http://example.com/image2"},
	}

	// Test getting a valid card by ID
	card, err := GetCardByID("1", cards)
	if err != nil {
		t.Fatalf("Failed to get card by ID: %v", err)
	}

	// Verify the data of the fetched card
	if card.ID != "1" || card.Name != "Card One" {
		t.Errorf("Unexpected card data: %+v", card)
	}

	// Test getting a non-existent card by ID
	_, err = GetCardByID("3", cards)
	if err == nil {
		t.Errorf("Expected error for non-existing card ID, got nil")
	}
}
