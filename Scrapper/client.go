package main

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// NewClient initialize the credentials for the scrapper
// You should always call this function before using the scrapper.
func NewClient() (*Client, error) {
	// connect to postgres database and ping it
	db, err := sqlx.Connect("postgres", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}

	return &Client{
		db:          db,
		binanceURL:  os.Getenv("BINANCE_URL"),
		taapiURL:    os.Getenv("TAAPI_URL"),
		taapiAPIKey: os.Getenv("TAAPI_API_KEY"),
	}, nil
}
