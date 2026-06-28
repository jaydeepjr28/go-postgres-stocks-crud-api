# Go PostgreSQL Stocks CRUD API

A simple RESTful API for managing stock data, built with Go, Gorilla Mux, and PostgreSQL.

## Features

- Complete CRUD operations (Create, Read, Update, Delete) for stock records.
- PostgreSQL database integration using `github.com/lib/pq`.
- Environment variable configuration using `github.com/joho/godotenv`.
- Routing using `github.com/gorilla/mux`.

## Prerequisites

- [Go](https://golang.org/) (v1.25.0 or later recommended)
- [PostgreSQL](https://www.postgresql.org/) database

## Installation and Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/jaydeepjr28/go-postgres-stocks-crud-api.git
   cd go-postgres-stocks-crud-api
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Configure Environment Variables:**
   Create a `.env` file in the root directory and add your PostgreSQL database connection string:
   ```env
   POSTGRES_URL="postgres://username:password@localhost:5432/dbname?sslmode=disable"
   ```
   *(Update the connection details with your actual PostgreSQL credentials).*

4. **Run the API:**
   ```bash
   go run cmd/main.go
   ```
   The server will start on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint                  | Description                   |
|--------|---------------------------|-------------------------------|
| GET    | `/api/stock`              | Retrieve all stocks           |
| GET    | `/api/stock/{id}`         | Retrieve a specific stock     |
| POST   | `/api/newstock`           | Create a new stock            |
| PUT    | `/api/stock/{id}`         | Update an existing stock      |
| PUT    | `/api/deletestock/{id}`   | Delete a stock                |

## Project Structure

```
├── .env                # Environment variables configuration
├── cmd/
│   └── main.go         # Application entry point
├── middleware/
│   └── handler.go      # Request handlers and database logic
├── models/
│   └── models.go       # Data models and structures
├── router/
│   └── router.go       # API route definitions
├── go.mod              # Go module dependencies
└── go.sum              # Go module checksums
```

## License

This project is open-source and available under the [MIT License](LICENSE).
