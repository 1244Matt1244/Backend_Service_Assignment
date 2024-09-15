package main

import (
	"Backend_Service_Assignment/internal/mtg"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// ImportHandler handles the import route
func ImportHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiURL := "https://api.magicthegathering.io/v1/cards"
		err := mtg.FetchCardsFromAPI(apiURL, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JSONResponse(w, map[string]string{"status": "import successful"}, http.StatusOK)
	}
}

// ListHandler handles the list route with pagination and filtering
func ListHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		if pageStr == "" {
			pageStr = "1"
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}

		// Get filters from query parameters (color, rarity, etc.)
		filters := mtg.CardFilters{
			Color:  r.URL.Query().Get("color"),
			Rarity: r.URL.Query().Get("rarity"),
			Type:   r.URL.Query().Get("type"),
			Name:   r.URL.Query().Get("name"),
		}

		cards, total, err := mtg.SearchCards(db, filters, page)
		if err != nil {
			http.Error(w, "Error fetching cards", http.StatusInternalServerError)
			return
		}

		JSONResponse(w, map[string]interface{}{
			"total": total,
			"page":  page,
			"items": cards,
		}, http.StatusOK)
	}
}

// CardHandler handles fetching card details by ID
func CardHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		card, err := mtg.GetCardByIDFromDB(db, id)
		if err != nil {
			http.Error(w, "Card not found", http.StatusNotFound)
			return
		}

		JSONResponse(w, card, http.StatusOK)
	}
}

// JSONResponse sends a JSON response with the given data and status code
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()

	// Setup database connection
	db, err := sqlx.Open("postgres", "user=postgres dbname=postgres sslmode=disable") // Update connection string
	if err != nil {
		log.Fatal(err)
	}

	// Define the routes
	r.HandleFunc("/import", ImportHandler(db)).Methods("POST")
	r.HandleFunc("/list", ListHandler(db)).Methods("GET")
	r.HandleFunc("/card/{id:[0-9]+}", CardHandler(db)).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
