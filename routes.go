package main

import (
	"Backend_Service_Assignment/internal/mtg"
	"Backend_Service_Assignment/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// ImportHandler handles the import route
func ImportHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiURL := "https://api.magicthegathering.io/v1/cards" // Update with the correct API endpoint
		err := mtg.FetchCardsFromAPI(apiURL, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JSONResponse(w, map[string]string{"status": "import successful"}, http.StatusOK)
	}
}

func JSONResponse(w http.ResponseWriter, map[string]string map[string]string, i int) {
	panic("unimplemented")
}

// ListHandler handles the list route
func ListHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement pagination and filtering
	}
}

// CardHandler handles the card details route
func CardHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement fetching card details by ID
	}
}

func main() {
	r := mux.NewRouter()

	// Setup database connection
	db, err := sqlx.Open("postgres", "user=postgres dbname=postgres sslmode=disable") // Update connection string
	if err != nil {
		log.Fatal(err)
	}

	r.HandleFunc("/import", ImportHandler(db)).Methods("POST")
	r.HandleFunc("/list", ListHandler(db)).Methods("GET")
	r.HandleFunc("/card/{id:[0-9]+}", CardHandler(db)).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
