package sql

import (
	"Appolo-api/app/config"
	"Appolo-api/app/models"

	"github.com/jmoiron/sqlx"
)

func CreateStrategy(c *config.Client, input *models.Strategy) error {
	_, err := c.DB.NamedExec(`INSERT INTO strategy (name, currency_name, leverage, interval, account_name) 
	VALUES (:name, :currency_name, :leverage, :interval, :account_name)`, input)
	if err != nil {
		return err
	}

	return nil
}

func GetStrategy(c *config.Client, name string) ([]models.Strategy, error) {
	var rows *sqlx.Rows
	var err error

	if name != "" {
		rows, err = c.DB.Queryx(`SELECT * FROM strategy WHERE name=$1`, name)
	} else {
		rows, err = c.DB.Queryx(`SELECT * FROM strategy`)
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var dest []models.Strategy
	for rows.Next() {
		var strat models.Strategy
		if err := rows.StructScan(&strat); err != nil {
			return nil, err
		}
		dest = append(dest, strat)
	}
	return dest, nil
}

func DeleteStrategy(c *config.Client, name string) error {
	_, err := c.DB.Exec(`DELETE FROM strategy WHERE name =$1`, name)
	if err != nil {
		return err
	}

	return nil
}
