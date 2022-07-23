package main

import (
	"encoding/json"
	"log"
)

func (client *Client) newHourlyDataFetchOnCurrencyRow(hdfoc *HourlyDataFetchOnCurrency) error {

	prep, err := client.db.PrepareNamed(`INSERT INTO hourly_data_fetch_on_currency (currency_name, open, high, low, close, volume, open_time, close_time, rsi, sma, stochastic, chopiness_index, bollinger_bands, obv, supertrend, ichimoku, ma) 
	VALUES (:currency_name, :open, :high, :low, :close, :volume, :open_time, :close_time, :rsi, :sma, :stochastic, :chopiness_index, :bollinger_bands, :obv, :supertrend, :ichimoku, :ma) RETURNING *`)
	if err != nil {
		return err
	}

	hdfoc2 := &HourlyDataFetchOnCurrency{}
	if err := prep.Get(hdfoc2, hdfoc); err != nil {
		return err
	}

	l, _ := json.Marshal(hdfoc2)
	log.Println("Inserting in hourly_data_fetch_on_currency : " + string(l))

	return nil
}
