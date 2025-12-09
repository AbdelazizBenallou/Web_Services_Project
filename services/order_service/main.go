package main

import (
	"fmt"
	"log"
	"net/http"
	"order_service/config"
	"order_service/delivery/http"
	"order_service/repository"
	"order_service/usecase"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repository, usecase, and handler
	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	// Use router from router.go
	router := handler.SetupRouter(orderHandler)

	// Add health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Service is healthy"))
	}).Methods("GET")

	// Start server
	port := ":8082"
	fmt.Printf("Order Service is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
