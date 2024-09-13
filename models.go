package main

// MTGCard represents a card from the MTG API
type MTGCard struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Colors       []string `json:"colors"`
	CMC          int      `json:"cmc"`
	Type         []string `json:"type"`
	Subtype      []string `json:"subtype"`
	Rarity       string   `json:"rarity"`
	OriginalText string   `json:"originalText"`
	ImageURL     string   `json:"imageUrl"`
}

// Camera represents a traffic camera
type Camera struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
