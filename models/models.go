package models

// Destination struct for decoded json data
type Destination struct {
	Image       string  `json:"image"`
	AltImage    string  `json:"altImage"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Continent   string  `json:"continent"`
	Stars       float64 `json:"stars"`
	Description string  `json:"description"`
}
