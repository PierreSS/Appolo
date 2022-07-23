package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/models"
	"Appolo-api/app/sql"
	"encoding/json"
	"net/http"
)

// @Summary Get the total pnl and commission for a strategy
// @Param strategy_name query string true "my_strategy"
// @Produce  json
// @Success 200 {object} models.TotalPNLAndCommission
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /strategy/result [get]
func getSumOfTradeForStrategy(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strategyName := r.URL.Query().Get("strategy_name")

		dest, err := sql.GetSumStrategyFromTradeHistory(c, strategyName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your strategy from the database: " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(dest)
	}
}
