// Package http handles HTTP delivery layer.
package http

import (
	"product_service/usecase"

	"github.com/gorilla/mux"
)

// NewRouter sets up the HTTP routes.
func NewRouter(uc *usecase.ProductUsecase) *mux.Router {
	r := mux.NewRouter()
	handler := NewProductHandler(uc)

	api := r.PathPrefix("/api/v1").Subrouter()
	products := api.PathPrefix("/products").Subrouter()

	products.HandleFunc("", handler.GetAll).Methods("GET")
	products.HandleFunc("", handler.Create).Methods("POST")
	products.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	products.HandleFunc("/{id}", handler.Update).Methods("PUT")
	products.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
	products.HandleFunc("/{id}/stock", handler.UpdateStock).Methods("PATCH")

	return r
}
