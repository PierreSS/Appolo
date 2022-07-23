package handlers

import (
	"Appolo-api/app/models"
	"encoding/json"
	"net/http"
)

// @Summary Get price for a given symbol
// @Param symbol query string true "BTCUSDT"
// @Produce  json
// @Success 200 {object} models.Prices
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/price [get]
func getPriceForSymbol(b *models.Binance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		symbol := r.URL.Query().Get("symbol")
		if !models.Symbol[symbol] {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid symbol"})
			return
		}

		resp, err := http.Get(b.BaseURL + "ticker/price?symbol=" + symbol)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance crypto prices failed"})
			return
		}
		defer resp.Body.Close()

		var res models.Prices
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil || resp.StatusCode != http.StatusOK {
			var gerr models.APIError
			json.NewDecoder(resp.Body).Decode(&gerr)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "reading of binance price failed" + gerr.Message})
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}
