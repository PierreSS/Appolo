package models

// Prices define the price of a symbol
type Prices struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
