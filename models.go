package main

// MTGCard represents a card from the MTG API
type MTGCard struct {
	ID           string   `json:"id"`           // Unique card ID
	Name         string   `json:"name"`         // Card name
	Colors       []string `json:"colors"`       // Colors of the card (e.g., red, green)
	CMC          int      `json:"cmc"`          // Converted mana cost
	Types        []string `json:"type"`         // Card types (e.g., Creature, Sorcery)
	Subtypes     []string `json:"subtype"`      // Card subtypes (e.g., Human, Elf)
	Rarity       string   `json:"rarity"`       // Card rarity (e.g., Common, Rare)
	OriginalText string   `json:"originalText"` // Original card text or abilities
	ImageURL     string   `json:"imageUrl"`     // URL of the card's image
}

// Camera represents a traffic camera
type Camera struct {
	ID        string  `json:"id"`        // Unique camera ID
	Name      string  `json:"name"`      // Name of the camera's location
	Latitude  float64 `json:"latitude"`  // Latitude coordinate of the camera
	Longitude float64 `json:"longitude"` // Longitude coordinate of the camera
}
