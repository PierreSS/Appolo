package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/helpers"
	"Appolo-api/app/models"
	"Appolo-api/app/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// saveFuturesOrderToTrade register a new trading line in db
func saveFuturesOrderToTrade(c *config.Client, o *models.FuturesOrder, strat string) error {
	trade := models.TradeHistory{
		CurrencyName:    o.Symbol,
		StrategyName:    strat,
		BinanceOrderID:  strconv.FormatInt(o.OrderID, 10),
		TransactionTime: helpers.MillisTimestampToRFC3339Format(o.UpdateTime),
		Type:            o.Type,
		Side:            o.Side,
	}

	p, err := strconv.ParseFloat(o.Price, 64)
	if err != nil {
		return err
	}
	trade.Price = p

	q, err := strconv.ParseFloat(o.OrigQty, 64)
	if err != nil {
		return err
	}
	trade.Quantity = q

	com, err := strconv.ParseFloat(o.Commission, 64)
	if err != nil {
		return err
	}
	trade.Commission = com

	pnl, err := strconv.ParseFloat(o.RealizedPNL, 64)
	if err != nil {
		return err
	}
	trade.PNL = pnl

	_, dberr := c.DB.NamedExec(`INSERT INTO trade_history (currency_name, strategy_name, binance_order_id, transaction_time, type, side, price, quantity, commission, pnl) 
		VALUES (:currency_name, :strategy_name, :binance_order_id, :transaction_time, :type, :side, :price, :quantity, :commission, :pnl)`, trade)
	if dberr != nil {
		return dberr
	}

	return nil
}

// get a strategy from input request
// func getStrat(c *config.Client, strat string) (models.Strategy, error) {
// 	var s models.Strategy

// 	if err := c.DB.Get(&s, "SELECT * FROM strategy WHERE name=$1", strat); err != nil {
// 		return models.Strategy{}, err
// 	}

// 	if s.Name == "" {
// 		return models.Strategy{}, errors.New("wrong strategy name")
// 	}

// 	return s, nil
// }

// change leverage for the given strategy
func switchLeverage(b *models.Binance, s *models.Strategy, a *models.Account) error {
	query := "symbol=" + s.CurrencyName + "&leverage=" + strconv.Itoa(s.Leverage) + "&recvWindow=100000&timestamp=" + b.GetCurrentTimestamp()
	signature := helpers.Sign([]byte(query), []byte(a.SecretKey))
	query = query + "&signature=" + signature

	req, err := http.NewRequest("POST", b.FuturesBaseURL+"v1/leverage?"+query, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-MBX-APIKEY", a.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		var gerr models.APIError
		json.NewDecoder(resp.Body).Decode(&gerr)

		return errors.New(gerr.Message)
	}
	defer resp.Body.Close()

	return nil
}

// get pnl, commission and price for last trade on binance
func getTradeInformations(b *models.Binance, o *models.FuturesOrder, a *models.Account) error {
	type AccountTradeListBinance struct {
		Commission  string `json:"commission"`
		OrderID     int64  `json:"orderId"`
		Price       string `json:"price"`
		RealizedPNL string `json:"realizedPnl"`
	}

	query := "symbol=" + o.Symbol + "&limit=5" + "&timestamp=" + b.GetCurrentTimestamp()
	signature := helpers.Sign([]byte(query), []byte(a.SecretKey))
	query = query + "&signature=" + signature

	req, err := http.NewRequest("GET", b.FuturesBaseURL+"v1/userTrades?"+query, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-MBX-APIKEY", a.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		var gerr models.APIError
		json.NewDecoder(resp.Body).Decode(&gerr)

		return errors.New(gerr.Message)
	}
	defer resp.Body.Close()

	var res []AccountTradeListBinance
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	for _, v := range res {
		if v.OrderID == o.OrderID {
			o.Price = v.Price
			o.RealizedPNL = v.RealizedPNL
			o.Commission = v.Commission

			break
		}
	}

	return nil
}

// @Summary Buy from binance futures using market order
// @Param strategy query string true "my_first_strategy"
// @Param quantity query string true "0.01"
// @Produce  json
// @Success 200 {object} models.FuturesOrder
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/futures/order/buy [post]
func futuresOrderMarketBuy(c *config.Client, b *models.Binance, side string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		input := helpers.IsInvalidInput(r.URL.Query(), "quantity", "strategy")
		if input != "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid " + input})
			return
		}

		// Get the strategy from input and check if currency is valid
		strat, err := sql.GetStrategy(c, r.URL.Query().Get("strategy"))
		if err != nil || strat == nil || !models.Symbol[strat[0].CurrencyName] {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid strategy"})
			return
		}

		// Get account object from account_name strategy to retrieve credentials
		acc, err := sql.GetAccount(c, strat[0].AccountName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your account: " + err.Error()})
			return
		}

		if err := switchLeverage(b, &strat[0], &acc[0]); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "switch leverage failed : " + err.Error()})
			return
		}

		query := "symbol=" + strat[0].CurrencyName + "&side=" + side + "&type=MARKET&quantity=" + r.URL.Query().Get("quantity") + "&timestamp=" + b.GetCurrentTimestamp()
		signature := helpers.Sign([]byte(query), []byte(acc[0].SecretKey))
		query = query + "&signature=" + signature

		req, err := http.NewRequest("POST", b.FuturesBaseURL+"v1/order?"+query, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "can't set header for request to binance api"})
			return
		}
		req.Header.Set("X-MBX-APIKEY", acc[0].ApiKey)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil || resp.StatusCode != http.StatusOK {
			var gerr models.APIError
			json.NewDecoder(resp.Body).Decode(&gerr)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance futures order failed: " + gerr.Message})
			return
		}
		defer resp.Body.Close()

		var res models.FuturesOrder
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance futures order failed : " + err.Error()})
			return
		}

		if err := getTradeInformations(b, &res, &acc[0]); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "request of binance get trade informations failed : " + err.Error()})
			return
		}

		if err := saveFuturesOrderToTrade(c, &res, strat[0].Name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "save futures order to trade failed : " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

// @Summary sell from binance futures using market order
// @Param strategy query string true "my_first_strategy"
// @Param quantity query string true "0.01"
// @Produce  json
// @Success 200 {object} models.FuturesOrder
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /binance/futures/order/sell [post]
func futuresOrderMarketSell(c *config.Client, b *models.Binance, side string) http.HandlerFunc {
	return futuresOrderMarketBuy(c, b, side)
}
