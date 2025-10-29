package main

import (
	"log"
	"os"

	"eventplanner-backend/internal/database"
	"eventplanner-backend/internal/handlers"
	"eventplanner-backend/internal/repositories"
	"eventplanner-backend/internal/router"
	"eventplanner-backend/internal/services"
)

func main() {
	// Read database URL from environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Initialize database connection pool
	pool, err := database.NewPostgresPool(databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Wire dependencies
	userRepo := repositories.NewUserRepository(pool)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	// Build router and start server
	r := router.New(authHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
