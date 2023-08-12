package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/shadeshade/bookings-api/config"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

type Hotels []Hotel

type Hotel struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (m *Repository) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func (m *Repository) AllHotels(w http.ResponseWriter, req *http.Request) {
	hotels := Hotels{
		Hotel{
			Title:   "Test title",
			Desc:    "Test description",
			Content: "Hello World",
		},
	}
	json.NewEncoder(w).Encode(hotels)
}
