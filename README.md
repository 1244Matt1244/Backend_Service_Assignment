# Backend Service Assignment

This project is a backend service for handling Magic: The Gathering (MTG) cards and traffic camera data. It includes endpoints to fetch data from APIs, manage a PostgreSQL database, and interact with the data.

## Project Structure

```
Backend_Service_Assignment/
├── cmd/
│   └── main.go           # Entry point
├── internal/
│   ├── camera/
│   │   ├── camera_handler.go
│   │   ├── camera_service.go
│   │   ├── camera_service_test.go
│   │   └── cameras.csv
│   ├── mtg/
│   │   ├── fetch_cards.go
│   │   ├── mtg_service_test.go
│   │   └── mtg_cards.csv
│   ├── db/
│   │   ├── db.go
│   │   └── db_query.go
│   └── utils/
│       └── utils.go
├── configs/
│   ├── docker-compose.yml
│   └── Dockerfile
├── routes/
│   └── routes.go
├── tests/
│   └── test.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

## Features

- **MTG Cards Management**: Fetch MTG cards from an external API and store them in a PostgreSQL database.
- **Camera Management**: Insert and retrieve traffic camera data.
- **PostgreSQL Integration**: Manage all data with PostgreSQL using Docker.
- **Testing**: Unit tests for service components.

## Prerequisites

- **Go 1.23+**
- **Docker** & **Docker Compose**
- **PostgreSQL**

## Setup and Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/1244Matt1244/Dev-Assignment.git
   cd Dev-Assignment
   ```

2. **Set up PostgreSQL**: You can either use Docker Compose to set up the database or configure a local PostgreSQL instance.

3. **Environment Variables**: Create a `.env` file with the following content:
   ```bash
   DATABASE_URL=postgresql://user:password@localhost:5432/myappdb?sslmode=disable
   ```

4. **Run Docker Compose**:
   ```bash
   docker-compose up
   ```

5. **Build and Run the Application**:
   ```bash
   go mod tidy
   go build -o myapp cmd/main.go
   ./myapp
   ```

6. **Run Tests**:
   ```bash
   go test ./internal/...
   ```

## API Endpoints

### MTG Cards

- **GET /api/mtg/cards**: Fetch all MTG cards from the database.
- **POST /api/mtg/cards/fetch**: Fetch cards from the MTG API and store them in the database.

### Cameras

- **GET /api/cameras**: Retrieve all traffic cameras.
- **POST /api/cameras/insert**: Insert camera data from `cameras.csv` into the database.

## Database Schema

### `mtg_cards`
| Column        | Type    | Description                  |
|---------------|---------|------------------------------|
| `id`          | STRING  | Unique card identifier        |
| `name`        | STRING  | Name of the card              |
| `colors`      | STRING  | List of colors associated     |
| `cmc`         | INT     | Converted mana cost           |
| `type`        | STRING  | Type of the card              |
| `subtype`     | STRING  | Subtype of the card           |
| `rarity`      | STRING  | Rarity of the card            |
| `original_text`| STRING | Card's original description   |
| `image_url`   | STRING  | Image URL                     |

### `cameras`
| Column        | Type    | Description                  |
|---------------|---------|------------------------------|
| `id`          | STRING  | Unique camera identifier      |
| `name`        | STRING  | Name of the camera            |
| `latitude`    | FLOAT   | Latitude of the camera        |
| `longitude`   | FLOAT   | Longitude of the camera       |

## Contributions

Contributions are welcome! Please feel free to submit pull requests or open issues.

## License

This project is licensed under the MIT License.

---

This `README.md` provides clear instructions on setup, running, and testing your project, as well as outlining the features and endpoints.