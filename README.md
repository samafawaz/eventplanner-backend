# EventPlanner Backend

Go backend server for the EventPlanner application using Gin framework.

## Getting Started

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

- `POST /signup` - Register a new user
- `POST /login` - Login with credentials
- `GET /health` - Health check

## Dependencies

- Gin web framework
- bcrypt for password hashing
- CORS middleware for cross-origin requests

## Development

The backend uses in-memory storage for Phase 0. All user data will be lost when the server restarts.
