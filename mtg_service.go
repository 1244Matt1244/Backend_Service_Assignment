package mtg

import (
    "fmt"
    "log"
    "myapp/db"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// FetchMTGCards fetches cards from the database
func FetchMTGCards() ([]Card, error) {
    conn, err := db.GetDBConnection() // Reuse the singleton connection
    if err != nil {
        return nil, fmt.Errorf("could not connect to db: %v", err)
    }

    rows, err := conn.Query("SELECT * FROM mtg_cards")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cards []Card
    for rows.Next() {
        var card Card
        if err := rows.Scan(&card.ID, &card.Name, &card.Type); err != nil {
            log.Printf("Error scanning card: %v", err)
            continue
        }
        cards = append(cards, card)
    }

    return cards, nil
}

// TestFetchMTGCards tests the FetchMTGCards function
func TestFetchMTGCards(t *testing.T) {
    cards, err := FetchMTGCards()
    assert.NoError(t, err, "Expected no error while fetching cards")
    assert.NotEmpty(t, cards, "Expected non-empty card list")
}

// FetchCardsFromAPI fetches cards from an external API (if applicable)
func FetchCardsFromAPI() error {
    resp, err := http.Get("https://api.magicthegathering.io/v1/cards")
    if err != nil {
        return fmt.Errorf("Failed to fetch cards from API: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("API returned status: %d", resp.StatusCode)
    }

    return nil
}

// TestFetchCardsFromAPI tests the API fetch function
func TestFetchCardsFromAPI(t *testing.T) {
	err := FetchCardsFromAPI("https://api.magicthegathering.io/v1/cards")
	assert.NoError(t, err, "Expected no error while fetching cards")
}

// TestFetchCardsFromAPIMock tests FetchCardsFromAPI using a mock server
func TestFetchCardsFromAPIMock(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"cards": [{"id": "1", "name": "Test Card"}]}`)
	}))
	defer mockServer.Close()

	err := FetchCardsFromAPI(mockServer.URL)
	assert.NoError(t, err, "Expected no error with mock server")
}
; // End