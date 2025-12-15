package routes

import (
	"net/http"
	"order_service/delivery/http/handler"
)

func SetupOrderRoutes(h *handler.OrderHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", h.Create)
	return mux
}

