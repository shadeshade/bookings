package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shadeshade/bookings-api/pkg/config"
	"github.com/shadeshade/bookings-api/pkg/handlers"
	"net/http"
)

func routes(config *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/hotels", handlers.Repo.AllHotels)
	mux.Post("/add-hotel", handlers.Repo.AddHotel)

	return mux
}
