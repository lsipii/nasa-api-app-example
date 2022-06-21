package main

import (
	"encoding/json"
	"fmt"
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
	epics := epic.GetEpics()
	json.NewEncoder(w).Encode(epics)
}

func main() {
	handleRequests()
}
