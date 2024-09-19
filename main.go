package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "Backend_Service_Assignment/docs" // Import Swagger docs

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger" // Swagger HTTP handler
)

var db *sql.DB

// Camera represents a traffic camera with ID, name, latitude, and longitude
type Camera struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// @title Backend Service Assignment API
// @version 1.0
// @description This is the Backend Service API for handling MTG cards and Cameras.
// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	// Initialize logging
	initLogging()

	// Establish database connection
	connectToDB()

	// Insert cameras from CSV
	err := InsertCamerasFromCSV(db, "/app/camera/cameras.csv")
	if err != nil {
		log.Fatalf("Error inserting cameras from CSV: %v", err)
	}

	// Setup router and start server
	r := SetupRouter()

	// Swagger handler
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// Initialize logging to log to both file and console
func initLogging() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file:", err)
	}

	// Log to both file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Establishes database connection with error handling
func connectToDB() {
	var err error
	log.Println("Connecting to the database...")
	db, err = sql.Open("postgres", "postgres://postgres:Prague1993@db:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check if the database is reachable
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot reach the database: %v", err)
	}
	log.Println("Connected to the database successfully.")
}

// InsertCamerasFromCSV reads cameras from a CSV file and inserts them into the database
func InsertCamerasFromCSV(db *sql.DB, filePath string) error {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Read CSV records and insert into the database
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading CSV record: %v", err)
			continue
		}

		// Parse latitude and longitude as floats
		name := record[0]
		lat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Invalid latitude for %s: %v", name, err)
			continue
		}
		lon, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Printf("Invalid longitude for %s: %v", name, err)
			continue
		}

		// Insert into the traffic_cameras table with UPSERT logic to prevent duplicates
		query := `INSERT INTO traffic_cameras (name, latitude, longitude, location) 
                  VALUES ($1, $2, $3, ST_SetSRID(ST_MakePoint($3, $2), 4326))
                  ON CONFLICT (name) DO UPDATE SET 
                  latitude = EXCLUDED.latitude, longitude = EXCLUDED.longitude`
		_, err = db.Exec(query, name, lat, lon)
		if err != nil {
			log.Printf("Error inserting camera %s: %v", name, err)
		} else {
			log.Printf("Inserted or updated camera: %s", name)
		}
	}
	log.Println("Camera data insertion complete.")
	return nil
}

// SetupRouter sets up routes for MTG cards, cameras, and health check
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Health Check Route
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	// Camera Routes
	r.HandleFunc("/cameras/{id}", GetCameraByIDHandler).Methods("GET")
	r.HandleFunc("/cameras", ListCamerasByRadiusHandler).Methods("GET")

	// CORS middleware for allowing requests
	r.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins for simplicity
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})))

	return r
}

func handlersMTGCardHandler(db *sql.DB) {
	panic("unimplemented")
}

// HealthCheck returns a simple message indicating server status
// @Summary Health Check
// @Description Check if the server is running
// @Tags health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check called.")
	w.Write([]byte("Server is running"))
}

// GetCameraByIDHandler retrieves a camera by its ID with error handling
// @Summary Get camera by ID
// @Description Retrieve a specific camera by its ID
// @Tags cameras
// @Accept json
// @Produce json
// @Param id path string true "Camera ID"
// @Success 200 {object} Camera
// @Failure 404 {string} string "Camera not found"
// @Router /cameras/{id} [get]
func GetCameraByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Printf("Fetching camera with ID: %s", id)

	cameraData, err := GetCameraByID(db, id)
	if err != nil {
		log.Printf("Error fetching camera with ID %s: %v", id, err)
		http.Error(w, fmt.Sprintf("Error fetching camera with ID %s: %v", id, err), http.StatusNotFound)
		return
	}
	JSONResponse(w, cameraData, http.StatusOK)
}

// ListCamerasByRadiusHandler lists cameras within a radius of a given location with error logging
// @Summary List cameras by radius
// @Description List cameras within a specified radius of a given location
// @Tags cameras
// @Accept json
// @Produce json
// @Param latitude query number true "Latitude"
// @Param longitude query number true "Longitude"
// @Param radius query number true "Radius in meters"
// @Success 200 {object} []Camera
// @Failure 400 {string} string "Invalid parameters"
// @Router /cameras [get]
func ListCamerasByRadiusHandler(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")
	radiusStr := r.URL.Query().Get("radius")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		log.Printf("Invalid latitude: %v", err)
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		log.Printf("Invalid longitude: %v", err)
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		log.Printf("Invalid radius: %v", err)
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	log.Printf("Finding cameras within a radius of %f meters at (%f, %f)", radius, lat, lon)
	cameras, err := FindCamerasWithinRadius(db, lat, lon, radius)
	if err != nil {
		log.Printf("Error finding cameras: %v", err)
		http.Error(w, "Error finding cameras", http.StatusInternalServerError)
		return
	}

	JSONResponse(w, map[string]interface{}{
		"total": len(cameras),
		"items": cameras,
	}, http.StatusOK)
}

// GetCameraByID retrieves a camera by its ID
func GetCameraByID(db *sql.DB, id string) (Camera, error) {
	query := `SELECT id, name, latitude, longitude FROM traffic_cameras WHERE id = $1`
	row := db.QueryRow(query, id)

	var camera Camera
	if err := row.Scan(&camera.ID, &camera.Name, &camera.Latitude, &camera.Longitude); err != nil {
		log.Printf("Error fetching camera with ID %s: %v", id, err)
		return Camera{}, fmt.Errorf("camera not found: %w", err)
	}
	log.Printf("Fetched camera: %+v", camera)
	return camera, nil
}

// FindCamerasWithinRadius finds cameras within a specified radius (in meters) from a given point
func FindCamerasWithinRadius(db *sql.DB, latitude, longitude, radius float64) ([]Camera, error) {
	query := `
		SELECT id, name, latitude, longitude 
		FROM traffic_cameras 
		WHERE ST_DWithin(
			ST_SetSRID(ST_MakePoint($1, $2), 4326),
			ST_SetSRID(ST_MakePoint(longitude, latitude), 4326),
			$3
		)
	`
	rows, err := db.Query(query, longitude, latitude, radius)
	if err != nil {
		log.Printf("Error querying cameras within radius: %v", err)
		return nil, fmt.Errorf("error querying cameras: %w", err)
	}
	defer rows.Close()

	var cameras []Camera
	for rows.Next() {
		var camera Camera
		if err := rows.Scan(&camera.ID, &camera.Name, &camera.Latitude, &camera.Longitude); err != nil {
			log.Printf("Error scanning camera row: %v", err)
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		cameras = append(cameras, camera)
		log.Printf("Found camera: %+v", camera)
	}
	log.Printf("Found %d cameras within radius", len(cameras))
	return cameras, nil
}

// JSONResponse sends a JSON response with error handling
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
