package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/models"
	"encoding/json"
	"net/http"
)

// @Summary Get Hourly data fetch of binance from the scrapper
// @Param symbol query string true "BTCUSDT"
// @Param x query string false "5"
// @Produce  json
// @Success 200 {object} []models.HourlyDataFetchOnCurrency
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /data/hourly [get]
func getHourlyDataFetchOnCurrency(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		symbol := r.URL.Query().Get("symbol")
		if !models.Symbol[symbol] {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid symbol"})
			return
		}

		x := r.URL.Query().Get("x")
		if x == "" {
			x = "10"
		}

		rows, err := c.DB.Queryx(`SELECT * FROM hourly_data_fetch_on_currency WHERE currency_name=$1
		ORDER BY inserted_date DESC LIMIT $2`, symbol, x)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving the hourly datas from the database: " + err.Error()})
			return
		}
		defer rows.Close()

		var d []models.HourlyDataFetchOnCurrency
		for rows.Next() {
			var data models.HourlyDataFetchOnCurrency
			if err := rows.StructScan(&data); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving the hourly datas from the database: " + err.Error()})
				return
			}
			d = append(d, data)
		}
		json.NewEncoder(w).Encode(d)
	}
}
