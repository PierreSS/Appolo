package sql

import (
	"Appolo-api/app/config"
	"Appolo-api/app/models"

	"github.com/jmoiron/sqlx"
)

func CreateAccount(c *config.Client, input *models.Account) error {
	_, err := c.DB.NamedExec(`INSERT INTO account (name, api_key, secret_key) 
		VALUES (:name, :api_key, :secret_key)`, input)
	if err != nil {
		return err
	}

	return nil
}

func GetAccount(c *config.Client, name string) ([]models.Account, error) {
	var rows *sqlx.Rows
	var err error

	if name != "" {
		rows, err = c.DB.Queryx(`SELECT * FROM account WHERE name=$1`, name)
	} else {
		rows, err = c.DB.Queryx(`SELECT * FROM account`)
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var dest []models.Account
	for rows.Next() {
		var acc models.Account
		if err := rows.StructScan(&acc); err != nil {
			return nil, err
		}
		dest = append(dest, acc)
	}
	return dest, nil
}

func DeleteAccount(c *config.Client, name string) error {
	_, err := c.DB.Exec(`DELETE FROM account WHERE name =$1`, name)
	if err != nil {
		return err
	}

	return nil
}
