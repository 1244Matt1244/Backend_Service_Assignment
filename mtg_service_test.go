package mtg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCardByID(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set up the expected query and result
	mock.ExpectQuery("SELECT (.+) FROM mtg_cards WHERE id = ?").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "colors", "cmc", "type", "subtype", "rarity", "image_url", "original_text"}).
			AddRow("1", "Card One", "[]", 3.0, "Creature", "Human", "Rare", "http://example.com/image.png", "This is the original text"))

	// Call the function to test
	card, err := GetCardByID(db, "1")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "1", card.ID)
	assert.Equal(t, "Card One", card.Name)
	assert.Equal(t, "Human", card.Subtype)
	assert.Equal(t, "Rare", card.Rarity)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
