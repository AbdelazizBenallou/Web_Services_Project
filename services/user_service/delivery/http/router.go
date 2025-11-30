// Package http handles HTTP delivery layer.
package http

import (
	"user_services/usecase"

	"github.com/gorilla/mux"
)

// NewRouter sets up the HTTP routes.
func NewRouter(uc *usecase.UserUsecase) *mux.Router {
	r := mux.NewRouter()
	handler := NewUserHandler(uc)

	api := r.PathPrefix("/api/v1").Subrouter()
	users := api.PathPrefix("/users").Subrouter()

	users.HandleFunc("", handler.GetAll).Methods("GET")
	users.HandleFunc("/{id}", handler.GetByID).Methods("GET")

	return r
}
