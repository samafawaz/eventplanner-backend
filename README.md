# EventPlanner Backend

Go backend server for the EventPlanner application using Gin framework and Postgres.

## Project Structure

```
internal/
  database/       # DB connection (pgx pool)
  handlers/       # HTTP handlers (Gin)
  models/         # Domain models and request DTOs
  repositories/   # Data access layer
  router/         # Router wiring and middleware
  services/       # Business logic
migrations/
  001_init.sql    # Users table
main.go           # Entry point wiring layers together
```

## Prerequisites

- Go 1.21+
- Postgres 13+

## Database Setup (Postgres)

1) Create a database and user (example):

```sql
CREATE DATABASE eventplanner;
CREATE USER eventplanner_user WITH PASSWORD 'strongpassword';
GRANT ALL PRIVILEGES ON DATABASE eventplanner TO eventplanner_user;
```

2) Apply migration:

```bash
psql "postgresql://eventplanner_user:strongpassword@localhost:5432/eventplanner?sslmode=disable" -f migrations/001_init.sql
```

3) Set environment variable before running the server:

```bash
# Windows PowerShell
$env:DATABASE_URL = "postgresql://eventplanner_user:strongpassword@localhost:5432/eventplanner?sslmode=disable"

# Linux/macOS
export DATABASE_URL="postgresql://eventplanner_user:strongpassword@localhost:5432/eventplanner?sslmode=disable"
```

## Getting Started

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

- `POST /signup` - Register a new user
  - body: `{ "name": string, "email": string, "password": string }`
- `POST /login` - Login with credentials
  - body: `{ "email": string, "password": string }`
- `GET /health` - Health check

## Dependencies

- Gin web framework
- pgx (Postgres driver)
- bcrypt for password hashing
- CORS middleware for cross-origin requests
