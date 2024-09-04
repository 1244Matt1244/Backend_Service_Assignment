# Assignment

## Project Overview

This application is a solution for an assignment. The goal is to create a backend service using the Go language and a PostgreSQL database to manage Magic The Gathering cards and traffic camera locations.

## Requirements

- Backend service written in Go
- PostgreSQL to be used as the database
- The project must include a Readme file with instructions on how to set up and run the project on a local machine
- The database and the application should run in their own Docker containers, preferably in their own network

## Installation and Running

1. Clone the repository from GitHub:
   ```sh
   git clone https://github.com/YourUsername/CollectiveMindAssignment.git
   ```
2. Install the necessary dependencies:
   ```sh
   go get -u
   ```
3. Start Docker compose to initialize the database:
   ```sh
   docker-compose up -d
   ```
4. Import the data about cards and cameras into the database using the `import.sh` script:
   ```sh
   ./import.sh
   ```
5. Run the application:
   ```sh
   go run main.go
   ```

## API Documentation

The application provides the following routes:

- `/import` - Imports card data from the API into the database.
- `/list` - Returns a list of cards or cameras based on search parameters.
- `/card/{CARD-ID}` - Returns details about a specific card based on its ID.

## Additional Notes

- The project is designed to be as clear and concise as possible, with minimal boilerplate code.
