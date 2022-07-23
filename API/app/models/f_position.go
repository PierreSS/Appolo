package models

import (
	"Appolo-api/app/helpers"
	"encoding/json"
	"time"
)

type FuturesPosition struct {
	Symbol           string    `json:"symbol"`
	Quantity         string    `json:"positionAmt"`
	EntryPrice       string    `json:"entryPrice"`
	UnrealizedProfit string    `json:"unRealizedProfit"`
	Leverage         string    `json:"leverage"`
	UpdatedTime      intToTime `json:"updateTime"`
}

//Int to time.Time alias
type intToTime struct {
	time.Time
}

func (i *intToTime) UnmarshalJSON(data []byte) error {

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	i.Time = helpers.MillisTimestampToRFC3339Format(int64(v.(float64)))

	return nil
}
