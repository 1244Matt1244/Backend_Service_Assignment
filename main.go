package main

import (
	"encoding/json"
	"log"
	"math"
	"models"
	"mtg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Parse MTG cards and cameras
	cards, err := mtg.ParseMTGCards("mtg_cards.csv")
	if err != nil {
		log.Fatal("Error parsing MTG cards:", err)
	}

	cameras, err := mtg.ParseCameras("cameras.csv")
	if err != nil {
		log.Fatal("Error parsing cameras:", err)
	}

	// Setup the router
	r := SetupRouter(cards, cameras)

	// Start the server
	log.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// SetupRouter sets up routes for MTG cards and cameras
func SetupRouter(cards []models.MTGCard, cameras []models.Camera) *mux.Router {
	r := mux.NewRouter()

	// MTG Card Routes
	r.HandleFunc("/cards/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		card, err := mtg.GetCardByID(id, cards)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		JSONResponse(w, card, http.StatusOK)
	}).Methods("GET")

	// Camera Routes
	r.HandleFunc("/cameras/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		camera, err := mtg.GetCameraByID(id, cameras)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		JSONResponse(w, camera, http.StatusOK)
	}).Methods("GET")

	// List Cameras with search by location
	r.HandleFunc("/cameras", func(w http.ResponseWriter, r *http.Request) {
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

		var results []models.Camera
		for _, camera := range cameras {
			if calculateDistance(lat, lon, camera.Latitude, camera.Longitude) <= radius {
				results = append(results, camera)
			}
		}

		JSONResponse(w, map[string]interface{}{
			"total": len(results),
			"items": results,
		}, http.StatusOK)
	}).Methods("GET")

	return r
}

// calculateDistance calculates the distance between two points on the earth (specified in decimal degrees)
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius of the Earth in kilometers
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

// JSONResponse sends a JSON response with the given data and status code
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
