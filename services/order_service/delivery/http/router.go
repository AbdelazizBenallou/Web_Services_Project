package handler

import (
	"github.com/gorilla/mux"
)

func SetupRouter(handler *OrderHandler) *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", handler.GetOrderByID).Methods("GET")
	router.HandleFunc("/orders/user/{user_id}", handler.GetOrdersByUserID).Methods("GET")
	router.HandleFunc("/orders/{id}", handler.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", handler.DeleteOrder).Methods("DELETE")
	
	return router
}