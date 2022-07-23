package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type response struct {
	Data []struct {
		ID     string                 `json:"id"`
		Result map[string]interface{} `json:"result"`
		Errors []interface{}          `json:"errors"`
	} `json:"data"`
}

func (c *Client) getIndicatorsFromTaapi(symbol string, hdfoc *HourlyDataFetchOnCurrency) (*HourlyDataFetchOnCurrency, error) {
	symbol = strings.Split(symbol, "USDT")[0] + "/USDT"

	var jsonStr = []byte(`{
		"secret": "` + c.taapiAPIKey + `",
    	"construct": {
			"exchange": "binance",
			"symbol": "` + symbol + `",
			"interval": "1h",
			"indicators": [
				{
					"id": "rsi",
					"indicator": "rsi",
					"backtrack": 1
				},
				{
					"id": "sma",
					"indicator": "sma",
					"optInTimePeriod": 7,
					"backtrack": 1
				},
				{
					"id": "stoch",
					"indicator": "stoch",
					"kPeriod": "70",
                    "kSmooth": "3",
					"backtrack": 1
				},
				{
					"id": "chop",
					"indicator": "chop",
					"backtrack": 1
				},
				{
					"id": "bbands",
					"indicator": "bbands",
					"backtrack": 1
				},
				{
					"id": "obv",
					"indicator": "obv",
					"backtrack": 1
				},
				{
					"id": "supertrend",
					"indicator": "supertrend",
                    "period": 10,
					"backtrack": 1
				},
                {
                    "id": "ichimoku",
                    "indicator": "ichimoku",
                    "conversionPeriod": 20,
                    "basePeriod": 60,
                    "spanPeriod": 120,
                    "displacement": 30,
                    "backtrack": 1 
                },
                {
					"id": "ma",
					"indicator": "ma",
                    "optInTimePeriod": 90,
					"backtrack": 1
				}
			]
		}
	}`)

	newr, err := http.NewRequest("POST", c.taapiURL+"bulk", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	newr.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(newr)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		type errorTaapi struct {
			E string `json:"error"`
		}
		var e errorTaapi
		json.NewDecoder(resp.Body).Decode(&e)

		return nil, errors.New(e.E)
	}
	defer resp.Body.Close()

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	for i, v := range res.Data {
		if len(v.Errors) > 0 {
			return nil, errors.New(v.Errors[0].(string))
		}

		switch v.ID {
		case "rsi":
			hdfoc.RSI = res.Data[i].Result["value"].(float64)
		case "sma":
			hdfoc.SMA = res.Data[i].Result["value"].(float64)
		case "stoch":
			hdfoc.Stochastic = res.Data[i].Result["valueK"].(float64)
		case "chop":
			hdfoc.Chop = res.Data[i].Result["value"].(float64)
		case "bbands":
			hdfoc.Bbands = append(hdfoc.Bbands, res.Data[i].Result["valueUpperBand"].(float64), res.Data[i].Result["valueMiddleBand"].(float64), res.Data[i].Result["valueLowerBand"].(float64))
		case "obv":
			hdfoc.OBV = res.Data[i].Result["value"].(float64)
		case "supertrend":
			hdfoc.Supertrend = res.Data[i].Result["valueAdvice"].(string)
		case "ichimoku":
			hdfoc.Ichimoku = append(hdfoc.Ichimoku, res.Data[i].Result["conversion"].(float64), res.Data[i].Result["base"].(float64))
		case "ma":
			hdfoc.MA = res.Data[i].Result["value"].(float64)
		}
	}
	return hdfoc, nil
}
