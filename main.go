package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"mtg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Database connection
	var err error
	db, err = mtg.ConnectToDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Setup the router
	r := SetupRouter()

	// Start the server
	log.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// SetupRouter sets up routes for MTG cards and cameras
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// MTG Card Routes
	r.HandleFunc("/cards/{id}", GetMTGCardByID).Methods("GET")
	r.HandleFunc("/import", ImportMTGCards).Methods("POST")
	r.HandleFunc("/list", ListMTGCards).Methods("GET")

	// Camera Routes
	r.HandleFunc("/cameras/{id}", GetCameraByID).Methods("GET")
	r.HandleFunc("/cameras", ListCamerasByRadius).Methods("GET")

	return r
}

// ImportMTGCards imports cards from the MTG API into the database
func ImportMTGCards(w http.ResponseWriter, r *http.Request) {
	err := mtg.ImportCardsFromAPI(db)
	if err != nil {
		http.Error(w, "Error importing cards", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cards imported successfully"))
}

// GetMTGCardByID retrieves an MTG card by its ID from the database
func GetMTGCardByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	card, err := mtg.GetCardByIDFromDB(db, id)
	if err != nil {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	JSONResponse(w, card, http.StatusOK)
}

// ListMTGCards lists MTG cards with search filters and pagination
func ListMTGCards(w http.ResponseWriter, r *http.Request) {
	filters := mtg.CardFilters{
		Color:  r.URL.Query().Get("color"),
		Rarity: r.URL.Query().Get("rarity"),
		Type:   r.URL.Query().Get("type"),
		Name:   r.URL.Query().Get("name"),
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	cards, total, err := mtg.SearchCards(db, filters, page)
	if err != nil {
		http.Error(w, "Error searching cards", http.StatusInternalServerError)
		return
	}

	JSONResponse(w, map[string]interface{}{
		"total": total,
		"page":  page,
		"items": cards,
	}, http.StatusOK)
}

// GetCameraByID retrieves a camera by its ID
func GetCameraByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	camera, err := mtg.GetCameraByIDFromDB(db, id)
	if err != nil {
		http.Error(w, "Camera not found", http.StatusNotFound)
		return
	}

	JSONResponse(w, camera, http.StatusOK)
}

// ListCamerasByRadius lists cameras within a radius of a given location
func ListCamerasByRadius(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")
	radiusStr := r.URL.Query().Get("radius")

	if latStr == "" || lonStr == "" || radiusStr == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	cameras, err := mtg.FindCamerasWithinRadius(db, lat, lon, radius)
	if err != nil {
		http.Error(w, "Error finding cameras", http.StatusInternalServerError)
		return
	}

	JSONResponse(w, map[string]interface{}{
		"total": len(cameras),
		"items": cameras,
	}, http.StatusOK)
}

// JSONResponse sends a JSON response with the given data and status code
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
