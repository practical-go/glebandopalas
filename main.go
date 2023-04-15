package main

import (
	"atomcourse/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Init Rouer
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB

	// Route Handlers / Endpoints
	r.HandleFunc("/news", api.HandleNews).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
