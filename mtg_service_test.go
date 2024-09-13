package mtg

import (
    "github.com/1244Matt1244/Backend_Service_Assignment/internal/models"
    "testing"
)


func TestParseMTGCards(t *testing.T) {
	// Create a temporary file with sample data for testing
	tmpfile, err := os.CreateTemp("", "test_cards_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	data := `id,name,colors,cmc,type,subtype,rarity,originalText,imageUrl
1,Card One,"red",1,Creature,Human,Common,A powerful creature,http://example.com/image1
2,Card Two,"green",2,Enchantment,Forest,Uncommon,An enchantment card,http://example.com/image2`
	_, err = tmpfile.Write([]byte(data))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	cards, err := ParseMTGCards(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to parse MTG cards: %v", err)
	}

	if len(cards) != 2 {
		t.Errorf("Expected 2 cards, got %d", len(cards))
	}

	if cards[0].ID != "1" || cards[0].Name != "Card One" {
		t.Errorf("Unexpected card data: %+v", cards[0])
	}
}

func TestGetCardByID(t *testing.T) {
	cards := []models.go{
		{ID: "1", Name: "Card One", Colors: []string{"red"}, CMC: 1, Type: []string{"Creature"}, Subtype: []string{"Human"}, Rarity: "Common", OriginalText: "A powerful creature", ImageURL: "http://example.com/image1"},
		{ID: "2", Name: "Card Two", Colors: []string{"green"}, CMC: 2, Type: []string{"Enchantment"}, Subtype: []string{"Forest"}, Rarity: "Uncommon", OriginalText: "An enchantment card", ImageURL: "http://example.com/image2"},
	}

	card, err := GetCardByID("1", cards)
	if err != nil {
		t.Fatalf("Failed to get card by ID: %v", err)
	}

	if card.ID != "1" || card.Name != "Card One" {
		t.Errorf("Unexpected card data: %+v", card)
	}

	_, err = GetCardByID("3", cards)
	if err == nil {
		t.Errorf("Expected error for non-existing card ID, got nil")
	}
}
