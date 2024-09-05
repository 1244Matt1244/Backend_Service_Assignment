// main.go
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
	router.HandleFunc("/import", ImportMTGCardsHandler).Methods("GET")          // Changed to GET as per the assignment
	router.HandleFunc("/list", ListMTGCardsHandler).Methods("GET")              // Changed route to /list as per the assignment
	router.HandleFunc("/card/{id:[a-f0-9]+}", GetMTGCardHandler).Methods("GET") // Added regex pattern to match the card ID
	router.HandleFunc("/list", ListCamerasInRadiusHandler).Methods("GET")       // Changed route to /list for camera list

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


