package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"myapp/internal/db"
	"myapp/internal/mtg"
	"myapp/pkg/camera"
)

func main() {
	fmt.Println("Starting application...")

	// Initialize the router
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/import", ImportMTGCardsHandler).Methods("POST")
	router.HandleFunc("/mtg/list", ListMTGCardsHandler).Methods("GET")
	router.HandleFunc("/card/{id}", GetMTGCardHandler).Methods("GET")
	router.HandleFunc("/cameras/list", ListCamerasInRadiusHandler).Methods("GET")

	// Start the server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Define your handler functions below
func ImportMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

func ListMTGCardsHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

func GetMTGCardHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}

func ListCamerasInRadiusHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}
