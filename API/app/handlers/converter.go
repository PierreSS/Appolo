package handlers

import (
	"Appolo-api/app/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Convert from asset one to two with a given quantity
// @Param from query string true "USDT"
// @Param to query string true "BTC"
// @Param quantity query string true "100"
// @Produce  json
// @Success 200 {object} models.Converter
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /convert [get]
func convertFromToWithQuantity(b *models.Binance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from, to, quantity := r.URL.Query().Get("from"), r.URL.Query().Get("to"), r.URL.Query().Get("quantity")
		if from == "" || to == "" || quantity == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid params"})
			return
		}

		symbol := to + from
		if !models.Symbol[symbol] {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid symbol"})
			return
		}

		q, err := strconv.ParseFloat(quantity, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "can't convert : " + err.Error()})
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

		p, err := strconv.ParseFloat(res.Price, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "can't convert : " + err.Error()})
			return
		}

		var con models.Converter

		con.From = from
		con.To = to
		con.Quantity = 1 / p * q

		json.NewEncoder(w).Encode(con)
	}
}
