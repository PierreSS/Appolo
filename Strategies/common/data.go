package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (c *Client) GetHourlyData() (*HourlyDataFetchOnCurrency, error) {
	resp, err := http.Get(c.APIURL + "data/hourly?symbol=BTCUSDT&x=1")
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		var err APIError
		json.NewDecoder(resp.Body).Decode(&err)

		return nil, errors.New(err.Message)
	}

	defer resp.Body.Close()

	var res []HourlyDataFetchOnCurrency
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (c *Client) GetHourlyDatas(rows int) ([]HourlyDataFetchOnCurrency, error) {
	resp, err := http.Get(c.APIURL + "data/hourly?symbol=BTCUSDT&x=" + strconv.Itoa(rows))
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		var err APIError
		json.NewDecoder(resp.Body).Decode(&err)

		return nil, errors.New(err.Message)
	}

	defer resp.Body.Close()

	var res []HourlyDataFetchOnCurrency
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}
