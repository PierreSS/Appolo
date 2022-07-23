package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/helpers"
	"Appolo-api/app/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// saveOrderToTrade register a new trading line in db
func saveOrderToTrade(c *config.Client, o *models.Order, strat string) error {
	trade := models.TradeHistory{
		CurrencyName:    o.Symbol,
		StrategyName:    strat,
		BinanceOrderID:  o.OrderID,
		TransactionTime: helpers.MillisTimestampToRFC3339Format(int64(o.TransactionTime)),
		Type:            o.Type,
		Side:            o.Side,
		PNL:             0,
	}

	var perr error
	f := func(s string) float64 {
		if perr != nil {
			return 0
		}

		var v float64
		v, perr = strconv.ParseFloat(s, 64)
		return v
	}

	var price, qty, com float64
	for _, v := range o.Fills {
		price = f(v.Price)
		qty = f(v.Quantity)
		com = f(v.Commission)

		if perr != nil {
			return perr
		}

		trade.Price += price
		trade.Quantity += qty
		trade.Commission += com
	}
	trade.Price = trade.Price / float64(len(o.Fills))

	_, err := c.DB.NamedExec(`INSERT INTO trade_history (currency_name, strategy_name, binance_order_id, transaction_time, type, side, price, quantity, commission, pnl) 
		VALUES (:currency_name, :strategy_name, :binance_order_id, :transaction_time, :type, :side, :price, :quantity, :commission, :pnl)`, trade)
	if err != nil {
		return err
	}

	return nil
}

// @Summary Buy from binance using market order
// @Param strategy query string true "my_first_strategy"
// @Param quantity query string true "0.01"
// @Produce  json
// @Success 200 {object} models.Order
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/order/buy [post]
func orderMarketBuy(c *config.Client, b *models.Binance, side string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		quantity := r.URL.Query().Get("quantity")
		if quantity == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid quantity"})
			return
		}

		strategy := r.URL.Query().Get("strategy")
		if strategy == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid strategy"})
			return
		}
		var s models.Strategy

		if err := c.DB.Get(&s, "SELECT * FROM strategy WHERE name=$1", strategy); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid strategy : " + err.Error()})
			return
		}

		if s.Name == "" || !models.Symbol[s.CurrencyName] {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid strategy"})
			return
		}

		query := "symbol=" + s.CurrencyName + "&side=" + side + "&type=MARKET&quantity=" + quantity + "&timestamp=" + b.GetCurrentTimestamp()
		signature := helpers.Sign([]byte(query), []byte(b.SecretKey))
		query = query + "&signature=" + signature

		req, err := http.NewRequest("POST", b.BaseURL+"order?"+query, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "can't set header for request to binance api"})
			return
		}
		req.Header.Set("X-MBX-APIKEY", b.APIKey)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil || resp.StatusCode != http.StatusOK {
			var gerr models.APIError
			json.NewDecoder(resp.Body).Decode(&gerr)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance order failed: " + gerr.Message})
			return
		}
		defer resp.Body.Close()

		var res models.Order
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance order failed"})
			return
		}
		if err := saveOrderToTrade(c, &res, s.Name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "save order to trade failed : " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

// @Summary sell from binance using market order
// @Param strategy query string true "my_first_strategy"
// @Param quantity query string true "0.01"
// @Produce  json
// @Success 200 {object} models.Order
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/order/sell [post]
func orderMarketSell(c *config.Client, b *models.Binance, side string) http.HandlerFunc {
	return orderMarketBuy(c, b, side)
}
