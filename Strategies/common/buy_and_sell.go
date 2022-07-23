package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (c *Client) BuyAndSell(side, strategyName, quantity string) error {
	resp, err := http.Post(c.APIURL+"binance/futures/order/"+side+"?strategy="+strategyName+"&quantity="+quantity, "", nil)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		var err APIError
		json.NewDecoder(resp.Body).Decode(&err)

		return errors.New(err.Message)
	}
	defer resp.Body.Close()

	return nil
}
