package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

type Hotel struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Hotels []Hotel

func allHotels(w http.ResponseWriter, req *http.Request) {
	hotels := Hotels{
		Hotel{
			Title:   "Test title",
			Desc:    "Test description",
			Content: "Hello World",
		},
	}
	json.NewEncoder(w).Encode(hotels)
}

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hotels", allHotels)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
