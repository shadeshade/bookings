package main

import (
	"github.com/shadeshade/bookings-api/internal/config"
	"github.com/shadeshade/bookings-api/internal/handlers"
	"github.com/shadeshade/bookings-api/internal/helpers"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Fatal(srv.ListenAndServe())
}
