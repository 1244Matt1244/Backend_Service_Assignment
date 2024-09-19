---

# Backend Service Assignment API

This project implements a backend service that handles both **Magic: The Gathering (MTG) Cards** and **Traffic Cameras**. The API allows for various operations such as retrieving MTG cards, listing them, importing them from an API, as well as working with traffic cameras including querying them by ID and finding cameras within a radius. The backend is built using **Go**, with a **PostgreSQL** database, including **PostGIS** for geographic querying.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Swagger Documentation](#swagger-documentation)
- [Environment Variables](#environment-variables)
- [Database](#database)
- [Contributing](#contributing)

## Features

- **MTG Cards:**
  - Fetch a specific card by ID
  - List cards with optional filters (color, rarity, type, name)
  - Import 500 MTG cards from an external API into PostgreSQL

- **Traffic Cameras:**
  - Fetch a camera by ID
  - List cameras within a specified radius using PostGIS

- **Health Check:**
  - Check if the server is running

## Tech Stack

- **Language**: Go (Golang)
- **Frameworks**: Gorilla Mux (for routing), Swaggo (for Swagger documentation)
- **Database**: PostgreSQL with PostGIS extension
- **Containerization**: Docker, Docker Compose
- **Logging**: Go standard logging with file and console output

## Project Structure

```
Backend_Service_Assignment/
├── cmd/
│   ├── camera/
│   │   ├── camera.go           # Camera related database operations
│   ├── mtg/
│   │   ├── mtg.go              # MTG card related database operations
├── handlers/
│   ├── handlers.go             # HTTP handlers for cameras and MTG cards
├── swagger/
│   └── docs.go                 # Swagger docs (auto-generated)
├── Dockerfile                  # Dockerfile for building the Go app
├── docker-compose.yml          # Docker Compose configuration for PostgreSQL and app
├── README.md                   # Project documentation
├── server.log                  # Log file for server output
└── go.mod                      # Go module dependencies
```

## Setup

### Prerequisites

- Docker and Docker Compose
- Go 1.23+
- PostgreSQL with PostGIS extension

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/1244Matt1244/Dev-Assignment.git
   cd Dev-Assignment
   ```

2. Run Docker Compose to start the PostgreSQL database and the Go app:
   ```bash
   docker-compose up --build
   ```

3. The application will start on **localhost:8080**.

## Running the Application

Once Docker Compose is running, the application will be available at `http://localhost:8080`. You can also access the Swagger UI documentation at `http://localhost:8080/swagger/index.html`.

## API Endpoints

### MTG Cards

- **Get MTG card by ID**
  - `GET /cards/{id}`
  - Fetch a specific MTG card by its ID.

- **List MTG cards**
  - `GET /list-cards?color=&rarity=&type=&name=&page=`
  - List MTG cards with optional filters (color, rarity, type, name) and pagination.

- **Import MTG cards**
  - `POST /import-cards`
  - Import 500 MTG cards from an external API into the PostgreSQL database.

### Traffic Cameras

- **Get Camera by ID**
  - `GET /cameras/{id}`
  - Retrieve a camera by its ID.

- **List Cameras by Radius**
  - `GET /cameras?latitude=&longitude=&radius=`
  - Find all traffic cameras within a specified radius from a given point using PostGIS.

### Health Check

- **Check server health**
  - `GET /health`
  - Check if the server is running.

## Swagger Documentation

This project includes API documentation generated with Swagger. To view and interact with the API documentation, visit:

```
http://localhost:8080/swagger/index.html
```

## Environment Variables

Make sure to configure the following environment variables, especially when deploying the app.

- `DATABASE_URL`: The URL for connecting to your PostgreSQL database (e.g., `postgres://username:password@db:5432/mydb`).

Example `.env` file:

```env
DATABASE_URL=postgres://<username>:<password>@<host>:5432/mydb
```

## Database

- **Database:** PostgreSQL with the PostGIS extension for geographic data.
- **Tables:**
  - `traffic_cameras`: Stores information about traffic cameras, including geographic location.
  - `mtg_cards`: Stores Magic: The Gathering card data.

### Initializing the Database

Data for traffic cameras is inserted from a CSV file (`/app/camera/cameras.csv`). If you need to update the camera data, place the updated CSV in the appropriate location, and the application will insert the data on startup.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for any improvements.

---