// models.go
package models

// MTGCard represents a card from the MTG API
type MTGCard struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ManaCost    string `json:"mana_cost"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Camera represents a traffic camera
type Camera struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
	URL      string `json:"url"`
}
