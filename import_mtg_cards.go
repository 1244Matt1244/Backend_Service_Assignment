package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func FetchMTGCards(minCards int) ([]MTGCard, error) {
	apiURL := "https://api.magicthegathering.io/v1/cards"
	var allCards []MTGCard
	page := 1

	for len(allCards) < minCards {
		resp, err := http.Get(fmt.Sprintf("%s?page=%d", apiURL, page))
		if err != nil {
			return nil, fmt.Errorf("error fetching MTG cards: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error fetching MTG cards: status code %d", resp.StatusCode)
		}

		var result struct {
			Cards []MTGCard `json:"cards"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding MTG cards JSON: %w", err)
		}

		allCards = append(allCards, result.Cards...)
		page++
	}

	return allCards[:minCards], nil
}

func InsertMTGCard(card MTGCard) error {
	_, err := db.Exec(`INSERT INTO mtg_cards (id, name, mana_cost, type, description)
					   VALUES ($1, $2, $3, $4, $5)
					   ON CONFLICT (id) DO NOTHING`,
		card.ID, card.Name, card.ManaCost, card.Type, card.Description)
	return err
}

func ImportMTGCards() error {
	minCards := 500
	cards, err := FetchMTGCards(minCards)
	if err != nil {
		return err
	}

	for _, card := range cards {
		if err := InsertMTGCard(card); err != nil {
			log.Printf("Error inserting card %s: %v", card.Name, err)
		}
	}

	return nil
}

func importMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := ImportMTGCards()
		if err != nil {
			http.Error(w, "Error importing MTG cards", http.StatusInternalServerError)
			log.Printf("Error importing MTG cards: %v", err)
			return
		}
		w.Write([]byte("MTG cards imported successfully"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func StartApp() {
	connStr := "user=postgres password=Prague1993 dbname=postgres host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open the database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	http.HandleFunc("/import-mtg-cards", importMTGCardsHandler)

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
