package models

import (
	"Appolo-api/app/config"
	"os"
	"strconv"
	"time"
)

type Binance struct {
	BaseURL   string
	APIKey    string
	SecretKey string

	FuturesBaseURL string
	// FuturesAPIKey    string
	// FuturesSecretKey string
}

// formatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func (b *Binance) GetCurrentTimestamp() string {
	return strconv.FormatInt(formatTimestamp(time.Now()), 10)
}

// NewBinance initialize an API binance instance with API key and secret key.
// You should always call this function before using binance API.
func NewBinance() *Binance {
	if config.UseTest {
		return &Binance{
			BaseURL:   os.Getenv("TESTNET_URL"),
			APIKey:    os.Getenv("TESTNET_BINANCE_API_KEY"),
			SecretKey: os.Getenv("TESTNET_BINANCE_SECRET_KEY"),

			FuturesBaseURL: os.Getenv("TESTNET_BINANCE_FUTURE_URL"),
			//FuturesAPIKey:    os.Getenv("TESTNET_BINANCE_FUTURE_API_KEY"),
			//FuturesSecretKey: os.Getenv("TESTNET_BINANCE_FUTURE_SECRET_KEY"),
		}
	}
	// not implemented yet
	return &Binance{
		BaseURL: os.Getenv("TESTNET_URL"),
	}
}
