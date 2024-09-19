package mtg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Initialize Redis client
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis server address
})

// MTGCard represents a card from the MTG API
type MTGCard struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Colors       []string `json:"colors"`
	CMC          float64  `json:"cmc"`
	Type         string   `json:"type"`
	Subtype      string   `json:"subtype"`
	Rarity       string   `json:"rarity"`
	OriginalText string   `json:"original_text"`
	ImageURL     string   `json:"image_url"`
}

// CardFilters represents filters for searching MTG cards
type CardFilters struct {
	Color  string
	Rarity string
	Type   string
	Name   string
}

// FetchCardsFromAPI fetches cards from the MTG API.
func fetchCardsFromAPI(page, pageSize int) ([]MTGCard, error) {
	apiURL := "https://api.magicthegathering.io/v1/cards?page=" + strconv.Itoa(page) + "&pageSize=" + strconv.Itoa(pageSize)
	log.Println("Fetching cards from API:", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error fetching cards from API: %v", err)
		return nil, fmt.Errorf("error fetching cards from API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: API returned status code %d", resp.StatusCode)
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var data struct {
		Cards []MTGCard `json:"cards"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding API response: %v", err)
		return nil, fmt.Errorf("error decoding API response: %w", err)
	}

	return data.Cards, nil
}

// ImportCardsFromAPI imports 500 MTG cards into the database.
func ImportCardsFromAPI(db *sql.DB) error {
	totalCards := 500
	pageSize := 100
	var allCards []MTGCard

	for page := 1; page <= (totalCards / pageSize); page++ {
		cards, err := fetchCardsFromAPI(page, pageSize)
		if err != nil {
			log.Printf("Error fetching cards from API: %v", err)
			return fmt.Errorf("error fetching cards from API: %w", err)
		}
		allCards = append(allCards, cards...)
		log.Printf("Fetched %d cards from page %d", len(cards), page)
	}

	for _, card := range allCards {
		if card.ImageURL == "" {
			card.ImageURL = "N/A"
		}
		if card.OriginalText == "" {
			card.OriginalText = "N/A"
		}
		if card.Subtype == "" {
			card.Subtype = "N/A"
		}

		colorsJSON, err := json.Marshal(card.Colors)
		if err != nil {
			log.Printf("Error converting Colors to JSON: %v, Card: %+v", err, card)
			return fmt.Errorf("error converting Colors to JSON: %w", err)
		}

		_, err = db.Exec(`INSERT INTO mtg_cards (id, name, colors, cmc, type, subtype, rarity, image_url, original_text)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (id) DO NOTHING`,
			card.ID, card.Name, colorsJSON, card.CMC, card.Type, card.Subtype, card.Rarity, card.ImageURL, card.OriginalText)
		if err != nil {
			log.Printf("Error inserting card into database: %v, Card: %+v", err, card)
			return fmt.Errorf("error inserting card: %w", err)
		}
		log.Printf("Inserted card into database: %s (%s)", card.Name, card.ID)
	}

	log.Println("Finished importing 500 cards into the database")
	return nil
}

// GetCardByID retrieves a specific MTG card from the database using its ID.
func GetCardByID(db *sql.DB, id string) (MTGCard, error) {
	var card MTGCard

	cachedCard, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		log.Printf("Cache miss for card ID: %s", id)
	} else if err != nil {
		log.Printf("Error accessing Redis: %v", err)
	} else {
		json.Unmarshal([]byte(cachedCard), &card)
		log.Printf("Cache hit for card ID: %s", id)
		return card, nil
	}

	var colorsJSON []byte
	err = db.QueryRow("SELECT id, name, colors, cmc, type, subtype, rarity, image_url, original_text FROM mtg_cards WHERE id = $1", id).
		Scan(&card.ID, &card.Name, &colorsJSON, &card.CMC, &card.Type, &card.Subtype, &card.Rarity, &card.ImageURL, &card.OriginalText)
	if err != nil {
		log.Printf("Error retrieving card with ID %s: %v", id, err)
		return MTGCard{}, fmt.Errorf("card not found: %w", err)
	}

	if err := json.Unmarshal(colorsJSON, &card.Colors); err != nil {
		log.Printf("Error unmarshalling colors: %v", err)
		return MTGCard{}, fmt.Errorf("error unmarshalling colors: %w", err)
	}

	cardJSON, _ := json.Marshal(card)
	rdb.Set(ctx, id, cardJSON, 10*time.Minute)

	log.Printf("Card found: %+v", card)
	return card, nil
}

// SearchCards searches and filters MTG cards based on provided filters.
func SearchCards(db *sql.DB, filters CardFilters, page int) ([]MTGCard, int, int, error) {
	query := "SELECT id, name, colors, cmc, type, subtype, rarity, image_url, original_text FROM mtg_cards WHERE 1=1"
	args := []interface{}{}
	argCounter := 1

	if filters.Color != "" {
		query += fmt.Sprintf(" AND colors::jsonb @> $%d::jsonb", argCounter)
		args = append(args, fmt.Sprintf(`["%s"]`, strings.Title(filters.Color)))
		argCounter++
	}
	if filters.Rarity != "" {
		query += fmt.Sprintf(" AND rarity ILIKE $%d", argCounter)
		args = append(args, filters.Rarity)
		argCounter++
	}
	if filters.Type != "" {
		query += fmt.Sprintf(" AND type ILIKE $%d", argCounter)
		args = append(args, filters.Type)
		argCounter++
	}

	query += fmt.Sprintf(" LIMIT 10 OFFSET $%d", argCounter)
	args = append(args, (page-1)*10)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("Error querying cards: %v", err)
		return nil, 0, 0, fmt.Errorf("error querying cards: %w", err)
	}
	defer rows.Close()

	var cards []MTGCard
	var colorsJSON []byte
	for rows.Next() {
		var card MTGCard
		if err := rows.Scan(&card.ID, &card.Name, &colorsJSON, &card.CMC, &card.Type, &card.Subtype, &card.Rarity, &card.ImageURL, &card.OriginalText); err != nil {
			log.Printf("Error scanning card row: %v", err)
			return nil, 0, 0, fmt.Errorf("row scan error: %w", err)
		}

		if err := json.Unmarshal(colorsJSON, &card.Colors); err != nil {
			log.Printf("Error unmarshalling colors: %v", err)
			return nil, 0, 0, fmt.Errorf("error unmarshalling colors: %w", err)
		}

		cards = append(cards, card)
	}

	var totalCards int
	if err := db.QueryRow("SELECT COUNT(*) FROM mtg_cards").Scan(&totalCards); err != nil {
		log.Printf("Error counting total cards: %v", err)
		return nil, 0, 0, fmt.Errorf("error counting total cards: %w", err)
	}

	totalPages := (totalCards + 9) / 10
	return cards, totalPages, totalCards, nil
}
