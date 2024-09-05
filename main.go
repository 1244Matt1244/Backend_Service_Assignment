// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Setup routes
	router := mux.NewRouter()

	// MTG Card routes
	router.HandleFunc("/import", ImportMTGCardsHandler).Methods("POST")
	router.HandleFunc("/mtg/list", ListMTGCardsHandler).Methods("GET")
	router.HandleFunc("/card/{id}", GetMTGCardHandler).Methods("GET")

	// Traffic Camera routes
	router.HandleFunc("/cameras/list", ListCamerasInRadiusHandler).Methods("GET")

	// Start the server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
