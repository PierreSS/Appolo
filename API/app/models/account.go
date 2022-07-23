package models

import "time"

type Account struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	ApiKey       string    `db:"api_key"`
	SecretKey    string    `db:"secret_key"`
	InsertedDate time.Time `db:"inserted_date"`
}
