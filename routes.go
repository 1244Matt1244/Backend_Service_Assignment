package main

import "github.com/gorilla/mux"

// SetupRoutes initializes the HTTP routes for the application
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/mtg/cards", ListMTGCardsHandler).Methods("GET")
	r.HandleFunc("/mtg/cards/{id:[0-9]+}", GetMTGCardHandler).Methods("GET")
	r.HandleFunc("/mtg/cards/import", ImportMTGCardsHandler).Methods("POST")
	r.HandleFunc("/cameras", ListCamerasHandler).Methods("GET")
	r.HandleFunc("/cameras/{id:[0-9]+}", GetCameraByIDHandler).Methods("GET")
	r.HandleFunc("/cameras", AddCameraHandler).Methods("POST")
	r.HandleFunc("/cameras/{id:[0-9]+}", DeleteCameraHandler).Methods("DELETE")
	r.HandleFunc("/cameras/radius", ListCamerasInRadiusHandler).Methods("GET")
	return r
}
