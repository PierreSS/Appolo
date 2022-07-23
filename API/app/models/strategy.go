package models

type Strategy struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	CurrencyName string `db:"currency_name"`
	Leverage     int    `db:"leverage"`
	Interval     string `db:"interval"`
	AccountName  string `db:"account_name"`
}
