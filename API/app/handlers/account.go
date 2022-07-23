package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/helpers"
	"Appolo-api/app/models"
	"Appolo-api/app/sql"
	"encoding/json"
	"net/http"
)

// @Summary Create a new account
// @Param name query string true "my_account"
// @Param api_key query string true "38fzefzef74723jh"
// @Param secret_key query string true "38fzefzef74723jh"
// @Produce  json
// @Success 200
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /account [post]
func createAccount(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		input := helpers.IsInvalidInput(r.URL.Query(), "name", "api_key", "secret_key")
		if input != "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid " + input})
			return
		}

		account := models.Account{
			Name:      r.URL.Query().Get("name"),
			ApiKey:    r.URL.Query().Get("api_key"),
			SecretKey: r.URL.Query().Get("secret_key"),
		}

		if err := sql.CreateAccount(c, &account); err != nil {
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error inserting account in database : " + err.Error()})
			return
		}

		// prep, err := c.DB.PrepareNamed(`INSERT INTO account (name, api_key, secret_key)
		// VALUES (:name, :api_key, :secret_key) RETURNING *`)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error inserting in db : " + err.Error()})
		// 	return
		// }
		// defer prep.Close()

		// account := models.Account{
		// 	Name:      name,
		// 	ApiKey:    api_key,
		// 	SecretKey: secret_key,
		// }

		// account2 := &models.Account{}
		// if err := prep.Get(account2, account); err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error inserting in db : " + err.Error()})
		// 	return
		// }

		json.NewEncoder(w).Encode(nil)
	}
}

// @Summary Get one or multiple account
// @Param name query string false "my_account"
// @Produce  json
// @Success 200 {object} []models.Account
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /account [get]
func getAccount(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 	var rows *sqlx.Rows
		// 	var err error

		name := r.URL.Query().Get("name")
		// 	if name != "" {
		// 		rows, err = c.DB.Queryx(`SELECT * FROM account WHERE name=$1`, name)
		// 	} else {
		// 		rows, err = c.DB.Queryx(`SELECT * FROM account`)
		// 	}

		// 	if err != nil {
		// 		w.WriteHeader(http.StatusInternalServerError)
		// 		json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your account: " + err.Error()})
		// 		return
		// 	}
		// 	defer rows.Close()

		// 	var a []models.Account
		// 	for rows.Next() {
		// 		var acc models.Account
		// 		if err := rows.StructScan(&acc); err != nil {
		// 			w.WriteHeader(http.StatusInternalServerError)
		// 			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your account: " + err.Error()})
		// 			return
		// 		}
		// 		a = append(a, acc)
		// 	}

		dest, err := sql.GetAccount(c, name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "There was a problem retrieving your account from the database: " + err.Error()})
			return
		}

		json.NewEncoder(w).Encode(dest)
	}
}

// @Summary Delete an account
// @Param name query string true "my_account"
// @Produce  json
// @Success 200
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /account [delete]
func deleteAccount(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid name"})
			return
		}

		if err := sql.DeleteAccount(c, name); err != nil {
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error deleting account in database : " + err.Error()})
			return
		}

		//var a models.Account

		// if err := c.DB.Get(&a, `DELETE FROM account WHERE name =$1 RETURNING *`, name); err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error deleting in db : " + err.Error()})
		// 	return
		// }

		json.NewEncoder(w).Encode(nil)
	}
}
