package models

// TotalPNLAndCommission is a custom returning struct for sum_trade handler
type TotalPNLAndCommission struct {
	PNL        float64 `db:"pnl"`
	Commission float64 `db:"commission"`
}
