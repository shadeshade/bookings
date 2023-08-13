package handlers

import (
	"encoding/json"
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

type Hotel struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Hotels []Hotel

var hotels Hotels

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func (m *Repository) AllHotels(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(hotels)
}

func (m *Repository) AddHotel(w http.ResponseWriter, r *http.Request) {
	var newHotel Hotel
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&newHotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hotels = append(hotels, newHotel)

	json.NewEncoder(w).Encode(hotels)
}
