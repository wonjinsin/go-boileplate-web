# Setup Instructions

## Install Dependencies

```bash
# Install chi router
go get -u github.com/go-chi/chi/v5

# Tidy up dependencies
go mod tidy

# Run the server
go run ./cmd/server
```

## Test Endpoints

```bash
# Health check
curl http://localhost:8080/healthz

# Create user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# List users
curl http://localhost:8080/users

# Get user by ID
curl http://localhost:8080/users/{id}
```

