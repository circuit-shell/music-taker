# Playlist Builder API

A Go-based REST API for managing song metadata and playlists. This project
demonstrates the implementation of a clean architecture using Go, SQLite, and
the Gin web framework.

## 🚀 Features

- RESTful API endpoints for song management
- SQLite database for data persistence
- Clean architecture implementation
- Comprehensive test coverage
- API documentation
- Error handling
- Input validation

## 📋 Prerequisites

- Go 1.21 or later
- SQLite3
- Make (optional, for using Makefile commands)

## 🛠️ Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/playlist-builder-back.git
cd playlist-builder-back
```

2. Install dependencies:

```bash
go mod download
```

3. Run the application:

```bash
go run cmd/api/main.go
```

## 🏗️ Project Structure

```
playlist-builder-back/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handler/
│   │   │   ├── greeting.go
│   │   │   ├── greeting_test.go
│   │   │   ├── song.go
│   │   │   └── song_test.go
│   │   └── router/
│   │       └── router.go
│   ├── model/
│   │   └── song.go
│   ├── repository/
│   │   ├── sqlite/
│   │   │   ├── song.go
│   │   │   └── song_test.go
│   │   └── interfaces.go
│   └── service/
│       ├── song.go
│       └── song_test.go
├── pkg/
│   ├── database/
│   │   └── sqlite.go
│   └── greeting/
│       ├── greeting.go
│       └── greeting_test.go
├── go.mod
├── go.sum
└── README.md
```

## 🔄 API Endpoints

### Songs

#### Create a Song

```bash
curl -X POST http://localhost:8080/api/v1/songs \
-H "Content-Type: application/json" \
-d '{
    "title": "Bohemian Rhapsody",
    "artist": "Queen",
    "album": "A Night at the Opera",
    "year": 1975,
    "genre": "Rock"
}'
```

#### Get All Songs

```bash
curl http://localhost:8080/api/v1/songs
```

### Greeting

#### Get Greeting

```bash
curl "http://localhost:8080/api/v1/greeting?name=John"
```

## 🧪 Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Generate coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📊 Using jq for Pretty Output

Pretty print all songs:

```bash
curl -s http://localhost:8080/api/v1/songs | jq '.'
```

Filter songs by year:

```bash
curl -s http://localhost:8080/api/v1/songs | jq '.[] | select(.year == 2020)'
```

Group songs by genre:

```bash
curl -s http://localhost:8080/api/v1/songs | jq 'group_by(.genre)'
```

## 🔧 Development

### Adding a New Endpoint

1. Define the model in `internal/model/`
2. Create repository interface in `internal/repository/interfaces.go`
3. Implement repository in `internal/repository/sqlite/`
4. Create service in `internal/service/`
5. Create handler in `internal/api/handler/`
6. Add route in `internal/api/router/router.go`
7. Add tests for all components

### Database Migrations

The SQLite database schema is automatically created when the application starts.
See `pkg/database/sqlite.go` for the schema definition.

## 📝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.

## 🤝 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [SQLite](https://www.sqlite.org/index.html)
- [testify](https://github.com/stretchr/testify)

## 📞 Contact

Your Name - [@yourusername](https://twitter.com/yourusername)

Project Link:
[https://github.com/yourusername/playlist-builder-back](https://github.com/yourusername/playlist-builder-back)

````

Optional: You might also want to add a `Makefile` to simplify common operations:

```makefile
.PHONY: build run test clean

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out
	rm -f *.db

# Format code
fmt:
	go fmt ./...

# Verify dependencies
verify:
	go mod verify

# Tidy dependencies
tidy:
	go mod tidy

# Run all quality checks
check: fmt verify test

# Install development tools
dev-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
lint:
	golangci-lint run

.DEFAULT_GOAL := run
````
