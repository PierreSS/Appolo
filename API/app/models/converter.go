package models

// Converter define the quantity converting 2 assets
type Converter struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Quantity float64 `json:"quantity"`
}
