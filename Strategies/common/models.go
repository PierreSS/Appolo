package common

import (
	"time"

	"github.com/lib/pq"
)

// APIError define API errors
type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

type Client struct {
	APIURL string
}

// HourlyDataFetchOnCurrency fetch data from the scrapping db
type HourlyDataFetchOnCurrency struct {
	ID           int             `json:"id"`
	CurrencyName string          `json:"currencyName"`
	Open         float64         `json:"open"`
	High         float64         `json:"high"`
	Low          float64         `json:"low"`
	Close        float64         `json:"close"`
	Volume       float64         `json:"volume"`
	OpenTime     time.Time       `json:"openTime"`
	CloseTime    time.Time       `json:"closeTime"`
	RSI          float64         `json:"rsi"`
	SMA          float64         `json:"sma"`
	Stochastic   float64         `json:"stochastic"`
	Chop         float64         `json:"chopinessIndex"`
	Bbands       pq.Float64Array `json:"bollingerBands" swaggertype:"array,number"`
	OBV          float64         `json:"obv"`
	Supertrend   string          `json:"supertrend"`
	Ichimoku     pq.Float64Array `json:"ichimoku" swaggertype:"array,number"`
	MA           float64         `json:"ma"`
	InsertedDate time.Time       `json:"insertedDate"`
}

type Convert struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Quantity float64 `json:"quantity"`
}

// Contains all the config.json struct
// type File struct {
// 	AtClosingBoughtQuantity  string `json:"at_closing_bought_quantity"`
// 	SupertrendBoughtQuantity string `json:"supertrend_bought_quantity"`
// }

type FuturesPosition struct {
	Symbol           string    `json:"symbol"`
	Quantity         string    `json:"positionAmt"`
	EntryPrice       string    `json:"entryPrice"`
	UnrealizedProfit string    `json:"unRealizedProfit"`
	Leverage         string    `json:"leverage"`
	UpdatedTime      time.Time `json:"updateTime"`
}
