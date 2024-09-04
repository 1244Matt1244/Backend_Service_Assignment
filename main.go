package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Fetching environment variables set in the Docker Compose file
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Constructing the dataSourceName
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Initialize the database with retry logic
	for i := 0; i < 5; i++ { // Retry 5 times
		err := InitializeDB(dataSourceName)
		if err != nil {
			log.Printf("Failed to initialize database, retrying in 5 seconds: %v", err)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	if db == nil {
		log.Fatalf("Failed to initialize database after retries")
	}

	// Setup routes
	router := mux.NewRouter()

	// MTG Card routes
	router.HandleFunc("/import", ImportMTGCardsHandler).Methods("POST")
	router.HandleFunc("/list", ListMTGCardsHandler).Methods("GET")
	router.HandleFunc("/card/{id}", GetMTGCardHandler).Methods("GET")

	// Traffic Camera routes
	router.HandleFunc("/list", ListCamerasInRadiusHandler).Methods("GET")

	// Start the server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
