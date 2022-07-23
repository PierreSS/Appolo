package models

import (
	"time"

	"github.com/lib/pq"
)

// HourlyDataFetchOnCurrency fetch data from the scrapping db
type HourlyDataFetchOnCurrency struct {
	ID           int             `db:"id" json:"id"`
	CurrencyName string          `db:"currency_name" json:"currencyName"`
	Open         float64         `db:"open" json:"open"`
	High         float64         `db:"high" json:"high"`
	Low          float64         `db:"low" json:"low"`
	Close        float64         `db:"close" json:"close"`
	Volume       float64         `db:"volume" json:"volume"`
	OpenTime     time.Time       `db:"open_time" json:"openTime"`
	CloseTime    time.Time       `db:"close_time" json:"closeTime"`
	RSI          float64         `db:"rsi" json:"rsi"`
	SMA          float64         `db:"sma" json:"sma"`
	Stochastic   float64         `db:"stochastic" json:"stochastic"`
	Chop         float64         `db:"chopiness_index" json:"chopinessIndex"`
	Bbands       pq.Float64Array `db:"bollinger_bands" json:"bollingerBands" swaggertype:"array,number"`
	OBV          float64         `db:"obv" json:"obv"`
	Supertrend   string          `db:"supertrend" json:"supertrend"`
	Ichimoku     pq.Float64Array `db:"ichimoku" json:"ichimoku" swaggertype:"array,number"`
	MA           float64         `db:"ma" json:"ma"`
	InsertedDate time.Time       `db:"inserted_date" json:"insertedDate"`
}
