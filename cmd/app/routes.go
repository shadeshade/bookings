package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shadeshade/bookings-api/internal/config"
	"github.com/shadeshade/bookings-api/internal/handlers"
	"net/http"
)

func routes(config *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Route("/api", func(mux chi.Router) {
		mux.Get("/rooms", handlers.Repo.GetAllRooms)
		mux.Get("/rooms/{id}", handlers.Repo.GetRoomByID)

		mux.Get("/availabilities", handlers.Repo.GetRoomAvailability)

		mux.Post("/reservations", handlers.Repo.CreateReservation)
	})

	return mux
}
