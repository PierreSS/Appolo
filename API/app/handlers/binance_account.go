package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/helpers"
	"Appolo-api/app/models"
	"encoding/json"
	"net/http"
)

// @Summary Get Account Informations about one or multiple asset
// @Param asset query string false "BTC"
// @Produce  json
// @Success 200 {object} models.BinanceAccount
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/account [get]
func getBinanceAccount(c *config.Client, b *models.Binance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		query := "&timestamp=" + b.GetCurrentTimestamp()
		signature := helpers.Sign([]byte(query), []byte(b.SecretKey))
		query = query + "&signature=" + signature

		req, err := http.NewRequest("GET", b.BaseURL+"account?"+query, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "can't set header for request to binance api"})
			return
		}

		req.Header.Set("X-MBX-APIKEY", b.APIKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance get account failed"})
			return
		}
		defer resp.Body.Close()

		var res models.BinanceAccount
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil || resp.StatusCode != http.StatusOK {
			var gerr models.APIError
			json.NewDecoder(resp.Body).Decode(&gerr)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance get account failed" + gerr.Message})
			return
		}

		asset := r.URL.Query().Get("asset")
		if asset != "" {
			if !models.Asset[asset] {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid asset"})
				return
			}

			var a struct {
				Asset     string
				Available string
			}
			for _, b := range res.Balances {
				if b.Asset == asset {
					a.Asset = b.Asset
					a.Available = b.Available

					json.NewEncoder(w).Encode(a)
					return
				}
			}
		}

		json.NewEncoder(w).Encode(res)
	}
}
