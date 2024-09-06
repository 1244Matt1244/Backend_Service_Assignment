package main

// MTGCard represents a card from the MTG API
type MTGCard struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ManaCost    string `json:"mana_cost"`
	Type        string `json:"type"` // Ensure the Type field exists
	Description string `json:"description"`
}

// Camera represents a traffic camera
type Camera struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
