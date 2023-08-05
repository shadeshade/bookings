package main

import (
	"encoding/json"
	"fmt"
	"github.com/shadeshade/bookings-api/config"
	"net/http"
)

const portNumber = ":8080"

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

type Hotel struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Hotels []Hotel

func (m *Repository) allHotels(w http.ResponseWriter, req *http.Request) {
	hotels := Hotels{
		Hotel{
			Title:   "Test title",
			Desc:    "Test description",
			Content: "Hello World",
		},
	}
	json.NewEncoder(w).Encode(hotels)
}

func (m *Repository) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func main() {
	var app config.AppConfig

	repo := NewRepo(&app)
	NewHandlers(repo)

	http.HandleFunc("/", Repo.Home)
	http.HandleFunc("/hotels", Repo.allHotels)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
