package main

import (
	"fmt"
	"log"
	"net/http"

	"kikukafandi/book-shop-api/internal/adapter/db"
	httpAdapter "kikukafandi/book-shop-api/internal/adapter/http"
	"kikukafandi/book-shop-api/internal/config"
	"kikukafandi/book-shop-api/internal/usecase"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database, err := config.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto migration
	if err := config.AutoMigrate(database); err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	// ====================================
	// DEPENDENCY INJECTION (Composition Root)
	// Flow: DB Repo → Usecase → Handler → Router
	// ====================================

	// Initialize repositories (adapters for database)
	bookRepo := db.NewBookRepositoryMySQL(database)
	userRepo := db.NewUserRepositoryMySQL(database)
	orderRepo := db.NewOrderRepositoryMySQL(database)

	// Initialize usecases (business logic)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, bookRepo, userRepo)

	// Initialize handlers (adapters for HTTP)
	bookHandler := httpAdapter.NewBookHandler(bookUsecase)
	userHandler := httpAdapter.NewUserHandler(userUsecase)
	orderHandler := httpAdapter.NewOrderHandler(orderUsecase)

	// Initialize router
	router := httpAdapter.NewRouter(bookHandler, userHandler, orderHandler)
	httpRouter := router.Setup()

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", addr)

	if err := http.ListenAndServe(addr, httpRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
