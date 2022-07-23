package models

import "time"

// TradeHistory from DB
type TradeHistory struct {
	ID              int       `db:"id"`
	CurrencyName    string    `db:"currency_name"`
	StrategyName    string    `db:"strategy_name"`
	BinanceOrderID  string    `db:"binance_order_id"`
	TransactionTime time.Time `db:"transaction_time"`
	Type            string    `db:"type"`
	Side            string    `db:"side"`
	Price           float64   `db:"price"`
	Quantity        float64   `db:"quantity"`
	Commission      float64   `db:"commission"`
	PNL             float64   `db:"pnl"`
}
