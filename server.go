package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/martinlindhe/base36"
)

func getUrls(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(urls)
}

type URL struct {
	ID             string `json:"id"`
	OrigialPath    string `json:"originalPath"`
	RedirectedPath string `json:"redirectedPath"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/urls", getUrls).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

var encodedURL1 = base36.EncodeBytes([]byte(strings.ToLower("https://www.google.com")))

var urls = []URL{
	URL{
		ID:             uuid.New().String(),
		OrigialPath:    encodedURL1,
		RedirectedPath: string(base36.DecodeToBytes(encodedURL1)),
	},
	URL{
		ID:             uuid.New().String(),
		OrigialPath:    "https://www.facebook.com",
		RedirectedPath: "https://www.x.com",
	},
}
