package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/helpers"
	"Appolo-api/app/models"
	"Appolo-api/app/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Create a new strategy
// @Param symbol query string true "BTCUSDT"
// @Param name query string true "my_first_strategy_BTCUSDT"
// @Param interval query string true "1h"
// @Param leverage query string true "10"
// @Param account_name query string true "undefined"
// @Produce  json
// @Success 200
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /strategy [post]
func createStrategy(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		input := helpers.IsInvalidInput(r.URL.Query(), "symbol", "name", "interval", "leverage", "account_name")
		if input != "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid " + input})
			return
		}

		strat := models.Strategy{
			Name:         r.URL.Query().Get("name"),
			CurrencyName: r.URL.Query().Get("symbol"),
			Interval:     r.URL.Query().Get("interval"),
			AccountName:  r.URL.Query().Get("account_name"),
		}
		strat.Leverage, _ = strconv.Atoi(r.URL.Query().Get("leverage"))

		if err := sql.CreateStrategy(c, &strat); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error inserting strategy in db : " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(nil)
	}
}

// @Summary Get one or multiple strategy
// @Param name query string false "my_first_strategy"
// @Produce  json
// @Success 200 {object} []models.Strategy
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /strategy [get]
func getStrategy(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")

		dest, err := sql.GetStrategy(c, name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your strategy from the database: " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(dest)
	}
}

// @Summary Delete a strategy
// @Param name query string true "my_first_strategy"
// @Produce  json
// @Success 200
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /strategy [delete]
func deleteStrategy(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid name"})
			return
		}

		if err := sql.DeleteStrategy(c, name); err != nil {
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error deleting strategy in database : " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(nil)
	}
}
