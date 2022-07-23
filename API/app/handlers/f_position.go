package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/helpers"
	"Appolo-api/app/models"
	"Appolo-api/app/sql"
	"encoding/json"
	"net/http"
)

// @Summary Get Futures position informations
// @Param symbol query string true "BTCUSDT"
// @Param account_name query string true "my_account"
// @Produce  json
// @Success 200 {object} models.FuturesPosition
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/futures/position [get]
func getFuturesPosition(c *config.Client, b *models.Binance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		input := helpers.IsInvalidInput(r.URL.Query(), "symbol", "account_name")
		if input != "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid " + input})
			return
		}

		acc, err := sql.GetAccount(c, r.URL.Query().Get("account_name"))
		if err != nil || acc == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your account: " + err.Error()})
			return
		}

		query := "symbol=" + r.URL.Query().Get("symbol") + "&timestamp=" + b.GetCurrentTimestamp()
		signature := helpers.Sign([]byte(query), []byte(acc[0].SecretKey))
		query = query + "&signature=" + signature

		req, err := http.NewRequest("GET", b.FuturesBaseURL+"v2/positionRisk?"+query, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "can't set header for request to binance futures api"})
			return
		}
		req.Header.Set("X-MBX-APIKEY", acc[0].ApiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			var gerr models.APIError
			json.NewDecoder(resp.Body).Decode(&gerr)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance futures get position failed: " + gerr.Message})
			return
		}
		defer resp.Body.Close()

		var res []models.FuturesPosition
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance futures get position failed : " + err.Error()})
			return
		}

		if res[0].Quantity[0:1] == "-" {
			res[0].Quantity = res[0].Quantity[1:]
		}

		json.NewEncoder(w).Encode(res[0])
	}
}
