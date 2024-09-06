package mtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCardByID(t *testing.T) {
	cards := []Card{
		{ID: "1", Name: "Card 1", Type: "Type A"},
	}

	card, err := GetCardByID("1", cards)
	assert.NoError(t, err)
	assert.Equal(t, "Card 1", card.Name)
	assert.Equal(t, "Type A", card.Type)
}
