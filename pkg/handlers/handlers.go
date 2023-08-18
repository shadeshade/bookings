package handlers

import (
	"fmt"
	"github.com/shadeshade/bookings-api/pkg/config"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Availability endpoint")
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostAvailability endpoint")
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact endpoint")
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reservation endpoint")
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostReservation endpoint")
}
