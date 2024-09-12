package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter sets up routes for MTG cards and cameras
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// MTG Card Routes
	r.HandleFunc("/cards/{id:[0-9]+}", GetMTGCardHandler).Methods("GET")
	r.HandleFunc("/cards", ListCards).Methods("GET")
	r.HandleFunc("/cards/import", ImportMTGCardsHandler).Methods("POST")

	// Camera Routes
	r.HandleFunc("/cameras/{id:[0-9]+}", GetCameraByIDHandler).Methods("GET")
	r.HandleFunc("/cameras", ListCamerasHandler).Methods("GET")
	r.HandleFunc("/cameras", AddCameraHandler).Methods("POST")
	r.HandleFunc("/cameras/{id:[0-9]+}", DeleteCameraHandler).Methods("DELETE")
	r.HandleFunc("/cameras/radius", ListCamerasInRadiusHandler).Methods("GET")

	return r
}

// JSONResponse sends a JSON response with a given status code
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
