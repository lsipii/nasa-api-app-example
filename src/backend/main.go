package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"app/authentication"
	"app/services/nasa/epic"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", fallback)
	router.Handle("/epic-api", authentication.AuthMiddleware(http.HandlerFunc(epicApiEndpoint)))
	log.Fatal(http.ListenAndServe(":8000", router))
}

func fallback(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid endpoint", http.StatusNotFound)
}

func epicApiEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: epicApiEndpoint")

	var query = epic.EpicQuery{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &query)
		if err != nil {
			http.Error(w, "Could parse request", http.StatusBadRequest)
		}
	}

	fmt.Println("query", query)

	epics := epic.GetEpics(query)
	json.NewEncoder(w).Encode(epics)
}

func main() {
	handleRequests()
}
