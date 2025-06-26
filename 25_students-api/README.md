
# Students API

This is a simple RESTful API for managing student records, built with Go. It uses SQLite as the database and supports basic CRUD operations for student data. The API is configured using a YAML file and leverages the `cleanenv` package for environment configuration.

## Features
- Create a new student (`POST /api/students`)
- Retrieve a student by ID (`GET /api/students/{id}`)
- List all students (`GET /api/students`)
- SQLite database for persistent storage
- YAML-based configuration
- Input validation for student data
- Graceful server shutdown

## Prerequisites
To run this project, you need the following installed:
- [Go](https://golang.org/dl/) (version 1.22.5 or later)
- [Git](https://git-scm.com/downloads)

## Installation
Follow these steps to set up and run the project locally:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/codebyankita/students-api.git
   cd students-api
   ```

2. **Install dependencies**:
   The project uses Go modules. Install the required packages by running:
   ```bash
   go mod tidy
   ```
   This will download all dependencies listed in `go.mod`, including:
   - `github.com/ilyakaznacheev/cleanenv` for configuration parsing
   - `github.com/mattn/go-sqlite3` for SQLite database support
   - `github.com/go-playground/validator/v10` for input validation
   - Other indirect dependencies (see `go.mod` for details)

3. **Set up the configuration file**:
   Create a `config` directory and a `local.yaml` file inside it with the following content:
   ```yaml
   env: "dev"
   storage_path: "storage/storage.db"
   http_server:
     address: "localhost:8082"
   ```
   You can place this file in `config/local.yaml`. Ensure the `storage` directory exists for the SQLite database:
   ```bash
   mkdir -p storage
   ```

4. **Run the application**:
   Use the following command to start the API server with the configuration file:
   ```bash
   go run cmd/students-api/main.go -config config/local.yaml
   ```
   The server will start on `localhost:8082` (or the address specified in `local.yaml`).

## Project Structure
```
students-api/
├── cmd/
│   └── students-api/
│       └── main.go          # Entry point of the application
├── config/
│   └── config.go           # Configuration loading logic
├── internal/
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go  # HTTP handlers for student endpoints
│   ├── storage/
│   │   └── sqlite/
│   │       └── sqlite.go   # SQLite storage implementation
│   ├── types/
│   │   └── types.go        # Data structures (Student struct)
│   └── utils/
│       └── response/
│           └── response.go # Utility for formatting HTTP responses
├── config/
│   └── local.yaml          # Configuration file (create manually)
├── storage/
│   └── storage.db          # SQLite database (created automatically)
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
└── README.md               # This file
```

## API Endpoints
The API provides the following endpoints:

### 1. Create a Student
- **Endpoint**: `POST /api/students`
- **Description**: Creates a new student record.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  }
  ```
- **Response**:
  - Success: `201 Created`
    ```json
    {
      "id": 1
    }
    ```
  - Error: `400 Bad Request` or `500 Internal Server Error`
    ```json
    {
      "status": "Error",
      "error": "field Name is required field"
    }
    ```

### 2. Get a Student by ID
- **Endpoint**: `GET /api/students/{id}`
- **Description**: Retrieves a student by their ID.
- **Response**:
  - Success: `200 OK`
    ```json
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 20
    }
    ```
  - Error: `400 Bad Request` or `500 Internal Server Error`

### 3. Get All Students
- **Endpoint**: `GET /api/students`
- **Description**: Retrieves a list of all students.
- **Response**:
  - Success: `200 OK`
    ```json
    [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "age": 20
      },
      ...
    ]
    ```
  - Error: `500 Internal Server Error`

## Configuration
The application uses a YAML configuration file (`config/local.yaml`). The structure is as follows:
- `env`: Environment name (e.g., `dev`, `prod`)
- `storage_path`: Path to the SQLite database file
- `http_server.address`: Address for the HTTP server (e.g., `localhost:8082`)

You can override the config file path using the `CONFIG_PATH` environment variable or the `-config` flag.

## Dependencies
Key dependencies include:
- `github.com/ilyakaznacheev/cleanenv`: For parsing YAML configuration
- `github.com/mattn/go-sqlite3`: SQLite driver for Go
- `github.com/go-playground/validator/v10`: For validating input data

Run `go get -u github.com/ilyakaznacheev/cleanenv` to ensure the configuration library is installed.

## Running Tests
To run tests (if any are added in the future):
```bash
go test ./...
```

## Building the Application
To build a binary:
```bash
go build -o students-api cmd/students-api/main.go
```
Run the binary with:
```bash
./students-api -config config/local.yaml
```

## Logging
The application uses `slog` for structured logging. Logs include:
- Server startup and shutdown events
- Student creation and retrieval events
- Errors during request processing

## Graceful Shutdown
The server handles `SIGINT` and `SIGTERM` signals for graceful shutdown, allowing up to 5 seconds to complete ongoing requests.

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make changes and commit (`git commit -m "Add feature"`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details (if applicable).

## Contact
For issues or questions, please open an issue on the [GitHub repository](https://github.com/codebyankita/students-api).
