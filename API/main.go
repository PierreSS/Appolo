//go:generate swag init

package main

import (
	"Appolo-api/app/config"
	"Appolo-api/app/handlers"
	"log"
	"net/http"
	"os"

	_ "Appolo-api/docs"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Set le port
func balanceTonPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port + ".")
	}
	return ":" + port
}

// @title Appolo API
// @version 1
// @description This is the documentation for the golang api of appolo-api

// @contact.name API Support
// @contact.email pierre.saintsorny@gmail.com

// @BasePath /1.0
func main() {
	godotenv.Load(".env")
	if _, b := os.LookupEnv("VERSION"); !b {
		log.Println("Env variables not set.")
		return
	}
	c, err := config.InitClient()
	if err != nil {
		log.Println("Client inititation failed : " + err.Error())
		return
	}
	defer c.DB.Close()

	addr := balanceTonPort()
	r := mux.NewRouter().StrictSlash(true)
	handler := cors.Default().Handler(r)

	handlers.HandleRequest(r, c)

	log.Println("Starting server.")
	log.Fatal(http.ListenAndServe(addr, handler))
}
