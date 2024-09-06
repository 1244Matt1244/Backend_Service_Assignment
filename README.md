Here's an updated `README.md` for your project:

```markdown
# MTG Backend Service Assignment

## Overview

This project is a Go-based backend service designed to manage Magic: The Gathering (MTG) cards and camera data. It includes functionalities to retrieve card details and camera information, and to interact with these data sources via HTTP endpoints.

## Project Structure

```
myproject/
├── card.go                # Contains the Card struct and GetCardByID function
├── main.go                # Main application entry point, sets up HTTP server and routes
├── mtg_service.go         # Contains MTG card service functionality
├── mtg_service_test.go    # Contains tests for MTG card service functionality
├── fetch_cards.go         # Fetches MTG cards from an external API
├── mtg_cards.csv          # CSV file containing MTG card data
├── README.md              # This file
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── Dockerfile             # Dockerfile for containerizing the application
├── docker-compose.yml     # Docker Compose file for multi-container setups
```

## Dependencies

- `github.com/gorilla/mux`: Router for handling HTTP routes.
- `github.com/stretchr/testify`: Library for testing assertions.

## Setup and Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/mtg-backend-service.git
   cd mtg-backend-service
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Build the Application**

   ```bash
   go build -o myapp.exe
   ```

4. **Run the Application**

   ```bash
   ./myproject.exe
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

- **card.go**: Defines the `Card` struct and the `GetCardByID` function.
- **main.go**: Sets up the HTTP server and routes.
- **mtg_service.go**: Contains service functions related to MTG cards.
- **mtg_service_test.go**: Tests for the MTG service functions.
- **fetch_cards.go**: Fetches card data from an external API.
- **mtg_cards.csv**: Contains card data for local testing.
- **Dockerfile**: Defines how to build the Docker image for the application.
- **docker-compose.yml**: Defines services and configurations for Docker Compose.

## Contributing

Feel free to open issues or submit pull requests to improve the project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```