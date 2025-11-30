// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	httpdelivery "user_services/delivery/http" // ✅ aliased
	"user_services/repository"
	"user_services/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	repo, err := repository.NewPostgresUserRepository(dbURL)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	uc := usecase.NewUserUsecase(repo)
	router := httpdelivery.NewRouter(uc) // ✅ using alias

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
