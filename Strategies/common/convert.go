package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (c *Client) Convert(from, to, quantity string) (float64, error) {
	resp, err := http.Get(c.APIURL + "convert?from=" + from + "&to=" + to + "&quantity=" + quantity)
	if err != nil {
		return 0, err
	} else if resp.StatusCode != http.StatusOK {
		var err APIError
		json.NewDecoder(resp.Body).Decode(&err)

		return 0, errors.New(err.Message)
	}
	defer resp.Body.Close()

	var res Convert
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, err
	}

	return res.Quantity, nil
}
