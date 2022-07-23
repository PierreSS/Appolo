package handlers

import (
	"Appolo-api/app/config"
	"Appolo-api/app/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HandleRequest gere toute les routes du serveur HTTP
func HandleRequest(router *mux.Router, c *config.Client) {

	router.HandleFunc("/", index).Methods("GET")

	// Subrouter to get the url match with the api version
	r := router.PathPrefix("/" + os.Getenv("VERSION")).Subrouter().StrictSlash(true)
	b := models.NewBinance()

	r.HandleFunc("/data/hourly", getHourlyDataFetchOnCurrency(c)).Methods("GET")

	r.HandleFunc("/strategy", getStrategy(c)).Methods("GET")
	r.HandleFunc("/strategy", createStrategy(c)).Methods("POST")
	r.HandleFunc("/strategy", deleteStrategy(c)).Methods("DELETE")

	r.HandleFunc("/strategy/result", getSumOfTradeForStrategy(c)).Methods("GET")

	r.HandleFunc("/account", getAccount(c)).Methods("GET")
	r.HandleFunc("/account", createAccount(c)).Methods("POST")
	r.HandleFunc("/account", deleteAccount(c)).Methods("DELETE")

	r.HandleFunc("/convert", convertFromToWithQuantity(b)).Methods("GET")

	// Binance path
	s := r.PathPrefix("/binance").Subrouter().StrictSlash(true)

	s.HandleFunc("/price", getPriceForSymbol(b)).Methods("GET")
	s.HandleFunc("/account", getBinanceAccount(c, b)).Methods("GET")

	s.HandleFunc("/order/buy", orderMarketBuy(c, b, "BUY")).Methods("POST")
	s.HandleFunc("/order/sell", orderMarketSell(c, b, "SELL")).Methods("POST")

	// Binance futures path
	f := s.PathPrefix("/futures").Subrouter().StrictSlash(true)

	f.HandleFunc("/account", getFuturesAccount(c, b)).Methods("GET")
	f.HandleFunc("/position", getFuturesPosition(c, b)).Methods("GET")

	f.HandleFunc("/order/buy", futuresOrderMarketBuy(c, b, "BUY")).Methods("POST")
	f.HandleFunc("/order/sell", futuresOrderMarketSell(c, b, "SELL")).Methods("POST")

	// Documentation
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

}

// @Summary Index of the api
// @Success 200
// @Router / [get]
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	fmt.Fprintf(w, "<h1>Hi there, welcome to the Appolo api !</h1>")
}
