package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Get futures active position
func (c *Client) GetFuturesPosition(account_name string) (FuturesPosition, error) {
	resp, err := http.Get(c.APIURL + "binance/futures/position?symbol=BTCUSDT&account_name=" + account_name)
	if err != nil {
		return FuturesPosition{}, err
	} else if resp.StatusCode != http.StatusOK {
		var err APIError
		json.NewDecoder(resp.Body).Decode(&err)

		return FuturesPosition{}, errors.New(err.Message)
	}

	defer resp.Body.Close()

	var res FuturesPosition
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return FuturesPosition{}, err
	}
	return res, nil
}
