package common

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// NewClient initialize the credentials for the strategies
// You should always call this function before using the strategies.
func NewClient() (*Client, error) {
	godotenv.Load(".env")
	if _, b := os.LookupEnv("API_URL"); !b {
		return &Client{}, errors.New("env variables not set")
	}

	return &Client{
		APIURL: os.Getenv("API_URL"),
	}, nil
}
