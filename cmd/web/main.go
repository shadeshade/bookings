package main

import (
	"fmt"
	"github.com/shadeshade/bookings-api/config"
	"github.com/shadeshade/bookings-api/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/hotels", handlers.Repo.AllHotels)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
