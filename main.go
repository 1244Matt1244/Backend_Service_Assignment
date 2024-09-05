// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/import", ImportMTGCardsHandler)
	http.HandleFunc("/import-cameras", ImportCamerasHandler)

	// Add other handlers as needed

	log.Fatal(http.ListenAndServe(":8080", nil))
}
