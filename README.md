# MTG Backend Service Assignment

## Overview

This Go-based backend service is designed to manage Magic: The Gathering (MTG) cards and camera data. The project includes functionalities to interact with these data sources through various services and HTTP endpoints.

## Project Structure

```
myapp/
├── camera.go               # Defines the Camera struct
├── camera_handler.go       # Handles HTTP requests for cameras
├── camera_service.go       # Contains logic for camera operations
├── camera_service_test.go  # Tests for camera service functionality
├── cameras.csv             # CSV file containing camera data
├── db.go                   # Contains database-related code
├── fetch_cards.go          # Fetches MTG cards from an external API
├── handlers.go             # HTTP request handlers
├── main.go                 # Main application entry point
├── models.go               # Defines application models
├── mtg_cards.csv           # CSV file containing MTG card data
├── mtg_service.go          # Contains MTG card service functionality
├── mtg_service_test.go     # Tests for MTG card service functionality
├── routes.go               # Defines application routes
├── test.go                 # Contains additional tests
├── utils.go                # Utility functions
├── Dockerfile              # Dockerfile for containerizing the application
├── docker-compose.yml      # Docker Compose file for multi-container setups
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── README.md               # This file
```

## Dependencies

- `github.com/gorilla/mux`: Router for handling HTTP routes.
- `github.com/stretchr/testify`: Library for testing assertions.

## Setup and Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/1244Matt1244/Dev-Assignment.git
   cd Dev-Assignment
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Build the Application**

   ```bash
   go build -o myapp
   ```

4. **Run the Application**

   ```bash
   ./myapp
   ```

## API Endpoints

- **GET /cards/{id}**: Retrieves MTG card details by ID.
- **GET /cameras/{id}**: Retrieves camera details by ID.

## Testing

To run the tests:

```bash
go test ./...
```

## File Descriptions

- **camera.go**: Defines the `Camera` struct used in the application.
- **camera_handler.go**: Handles HTTP requests related to cameras.
- **camera_service.go**: Contains logic for camera operations and service functions.
- **camera_service_test.go**: Tests for camera-related services.
- **cameras.csv**: CSV file containing camera data.
- **db.go**: Includes database-related code for data management.
- **fetch_cards.go**: Contains logic for fetching MTG card data from an external API.
- **handlers.go**: Contains HTTP request handlers for various endpoints.
- **main.go**: The main entry point for the application, sets up the HTTP server and routes.
- **models.go**: Defines application models used throughout the code.
- **mtg_cards.csv**: CSV file with MTG card data for local testing.
- **mtg_service.go**: Contains MTG card service functionality.
- **mtg_service_test.go**: Tests for MTG card services.
- **routes.go**: Defines application routes and their handlers.
- **test.go**: Contains additional tests for various functionalities.
- **utils.go**: Utility functions used in the application.
- **Dockerfile**: Defines how to build a Docker image for the application.
- **docker-compose.yml**: Defines services and configurations for Docker Compose.

## Contributing

Feel free to open issues or submit pull requests to improve the project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

This version reflects the `cameras.csv` file, ensuring it is properly acknowledged in the project structure and descriptions.