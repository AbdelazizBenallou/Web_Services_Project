package handler

import (
	"encoding/json"
	"net/http"
	"order_service/domain"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	orderUsecase domain.OrderUsecase
}

func NewOrderHandler(orderUsecase domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	if err := h.orderUsecase.CreateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	
	order, err := h.orderUsecase.GetOrderByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if order == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	orders, err := h.orderUsecase.GetOrdersByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	order.ID = id
	
	if err := h.orderUsecase.UpdateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	
	if err := h.orderUsecase.DeleteOrder(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}