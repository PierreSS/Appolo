package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Binance Candlestick structuration
// [
//   [
//     1499040000000,      // Open time
//     "0.01634790",       // Open
//     "0.80000000",       // High
//     "0.01575800",       // Low
//     "0.01577100",       // Close
//     "148976.11427815",  // Volume
//     1499644799999,      // Close time
//     "2434.19055334",    // Quote asset volume
//     308,                // Number of trades
//     "1756.87402397",    // Taker buy base asset volume
//     "28.46694368",      // Taker buy quote asset volume
//     "17928899.62484339" // Ignore.
//   ]
// ]
func (client *Client) getHourlyCandlestickForSymbol(symbol string) (*HourlyDataFetchOnCurrency, error) {
	resp, err := http.Get(client.binanceURL + "klines?symbol=" + symbol + "&interval=1h&limit=1")
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		var err BinanceAPIError
		json.NewDecoder(resp.Body).Decode(&err)

		return nil, errors.New(err.Message)
	}

	defer resp.Body.Close()

	type result [][]interface{}
	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	var candle HourlyDataFetchOnCurrency
	var parseErr error

	str := make([]string, 5)

	str[0] = res[0][1].(string)
	str[1] = res[0][2].(string)
	str[2] = res[0][3].(string)
	str[3] = res[0][4].(string)
	str[4] = res[0][5].(string)

	candle.OpenTime = millisTimestampToRFC3339Format(int64(res[0][0].(float64)))
	candle.CloseTime = millisTimestampToRFC3339Format(int64(res[0][6].(float64)))

	flo := make([]float64, len(str))
	for i, v := range str {
		flo[i], parseErr = strconv.ParseFloat(v, 64)
		if parseErr != nil {
			return nil, parseErr
		}
	}

	candle.Open = flo[0]
	candle.High = flo[1]
	candle.Low = flo[2]
	candle.Close = flo[3]
	candle.Volume = flo[4]

	return &candle, nil
}
