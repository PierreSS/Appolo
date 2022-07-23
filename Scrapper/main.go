package main

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
)

type Client struct {
	db          *sqlx.DB
	binanceURL  string
	taapiURL    string
	taapiAPIKey string
}

// BinanceAPIError define API errors
type BinanceAPIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

// OHLCV from binance
type OHLCV struct {
	Open      float64 `db:"open"`
	High      float64 `db:"high"`
	Low       float64 `db:"low"`
	Close     float64 `db:"close"`
	Volume    float64 `db:"volume"`
	OpenTime  string  `db:"open_time"`
	CloseTime string  `db:"close_time"`
}

// HourlyDataFetchOnCurrency is used to insert fetched datas
type HourlyDataFetchOnCurrency struct {
	ID           int    `db:"id"`
	CurrencyName string `db:"currency_name"`
	OHLCV
	RSI        float64 `db:"rsi"`
	SMA        float64 `db:"sma"`
	Stochastic float64 `db:"stochastic"`
	Chop       float64 `db:"chopiness_index"`
	//Bbands       SliceToArray    `db:"bollinger_bands"`
	Bbands       pq.Float64Array `db:"bollinger_bands"`
	OBV          float64         `db:"obv"`
	Supertrend   string          `db:"supertrend"`
	Ichimoku     pq.Float64Array `db:"ichimoku"`
	MA           float64         `db:"ma"`
	InsertedDate string          `db:"inserted_date"`
}

//type SliceToArray []float64

func main() {
	c := cron.New()
	c.AddFunc("5 0 * * * *", func() {

		log.Println("Starting.")

		godotenv.Load(".env")
		if _, b := os.LookupEnv("BINANCE_URL"); !b {
			log.Println("Env variables not set.")
			return
		}

		client, err := NewClient()
		if err != nil {
			log.Println("Error generating credentials / connecting db : " + err.Error())
			return
		}
		defer client.db.Close()

		// Register OHLCV from binance to db
		currency := "BTCUSDT"
		cohlcv, err := client.getHourlyCandlestickForSymbol(currency)
		if err != nil {
			log.Println("Error connecting to binance : " + err.Error())
			return
		}
		cohlcv.CurrencyName = currency

		hourlyData, err := client.getIndicatorsFromTaapi(currency, cohlcv)
		if err != nil {
			log.Println("Error connecting to taapi : " + err.Error())
			return
		}

		if err := client.newHourlyDataFetchOnCurrencyRow(hourlyData); err != nil {
			log.Println("Error registering into the database : ", err.Error())
		}

		log.Println("Done.")

	})

	c.Start()
	for {
		time.Sleep(time.Second)
	}
}
