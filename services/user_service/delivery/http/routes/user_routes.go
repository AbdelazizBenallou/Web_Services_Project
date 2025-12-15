package routes

import (
	"user_service/delivery/http/handler"
	"encoding/json"
	"net/http"
)

func SetupUserRoutes(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Use Go 1.22+ pattern matching
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})

	return mux
}
