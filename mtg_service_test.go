// mtg_service_test.go
package mtg

import (
	"fmt"
	"log"
	"myapp/db" // Import your db package
)

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
