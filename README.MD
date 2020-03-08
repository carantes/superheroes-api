# Superheroes API
Golang app to manage superheroes data using the [superhero API](https://superheroapi.com/).


## Project Structure

The goal here is to keep the project struture very simple, thinking in medium size web applications.

```
.
├── build                     # Compiled files
├── bundles
    ├── superheroesbundle     # Superheroes bundle
├── core                      # Core/base structs and interfaces
    ├── config.go             # Configuration from environment variables
    ├── controller.go         # Controller base struct
    ├── server.go             # Server
└── main.go                   # Root main file
```

## Requirements (WIP)

- GO

## Environments

Dotenv load environment variables from `.env` file.

```
API_PORT=8080
API_PREFIX=/api/v1/
```

## Commands (WIP)

- Start the app

    `go run .`

- Run Tests

    `go test`


## Packages

- [Gorilla MUX](github.com/gorilla/mux): HTTP Router and for building webservers
- [GO DotEnv](github.com/joho/godotenv): Load environment variables from .env
- [UUID](github.com/satori/go.uuid): Generate and Parse Universally Unique Identifier (UUID)
- [Testify](github.com/stretchr/testify/assert): Easy assertions
- [SQL Mock](github.com/DATA-DOG/go-sqlmock): Mock database interactions
