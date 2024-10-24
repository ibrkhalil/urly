package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/deatil/go-encoding/base62"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type URL struct {
	ID             string `json:"id"`
	OrigialPath    string `json:"originalPath"`
	RedirectedPath string `json:"redirectedPath"`
}

var urls = []URL{
	{
		ID:             "123",
		OrigialPath:    encodedURL1,
		RedirectedPath: string(decodedURL1),
	},
	{
		ID:             uuid.New().String(),
		OrigialPath:    encodedURL2,
		RedirectedPath: string(decodedURL2),
	},
}

var encodedURL1 = base62.StdEncoding.EncodeToString([]byte(strings.ToLower("https://www.google.com")))
var encodedURL2 = base62.StdEncoding.EncodeToString([]byte(strings.ToLower("https://www.facebook.com")))
var decodedURL1, _ = base62.StdEncoding.DecodeString(encodedURL1)
var decodedURL2, _ = base62.StdEncoding.DecodeString(encodedURL2)

func getUrls(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(urls)
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok {
		panic("OMG: No id params")
	}

	found := false

	for _, v := range urls {
		if id[0] == v.ID {
			json.NewEncoder(w).Encode(v)
			found = true
		}
	}
	if !found {
		// panic(("OMG: Didn't find ID of " + id[0]))
		errorHandler(w, http.StatusNotFound)

	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/urls", getUrls).Methods("GET")
	router.HandleFunc("/url", getUrl).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Something wrong happened, Scream and run in circles!")
	}
}
