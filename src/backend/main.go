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
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", fallback)
	myRouter.HandleFunc("/epic-api", epicApiEndpoint)
	myRouter.Use(authentication.Middleware)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func fallback(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid endpoint", http.StatusNotFound)
}

func epicApiEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: epicApiEndpoint")
	articles := epic.GetArticles()
	json.NewEncoder(w).Encode(articles)
}

func main() {
	handleRequests()
}
