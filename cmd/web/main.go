package main

import (
	"github.com/shadeshade/bookings-api/pkg/config"
	"github.com/shadeshade/bookings-api/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Fatal(srv.ListenAndServe())
}
