# Go HTTP Server Demo

This project demonstrates various aspects of building a production-ready HTTP server in Go, including:

- Basic HTTP routing
- Middleware implementation
- Error handling
- Configuration management
- Structured logging
- Graceful shutdown

## Prerequisites

- Go 1.21 or higher
- Basic understanding of Go programming language

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── handlers.go
│   ├── middleware/
│   │   └── middleware.go
│   └── config/
│       └── config.go
├── go.mod
└── README.md
```

## Setup and Running

1. Clone the repository:

```bash
git clone git@github.com:shivamshahi07/go-http.git
cd go-http
```

2. Install dependencies:

```bash
go mod download
```

3. Run the server:

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

## Available Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/hello` - Hello world endpoint
- `POST /api/v1/echo` - Echo endpoint that returns the request body

## Testing

Run the tests using:

```bash
go test ./...
```

## Features Demonstrated

1. **Routing**: Using the standard `net/http` package
2. **Middleware**: Custom middleware for logging and request tracking
3. **Error Handling**: Structured error responses
4. **Configuration**: Environment-based configuration
5. **Graceful Shutdown**: Proper server shutdown handling
6. **Logging**: Structured logging with request context
