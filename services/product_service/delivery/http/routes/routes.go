package routes

import (
	"net/http"

	"product_service/delivery/http/handler"
)

func Setup(
	categoryHandler *handler.CategoryHandler,
	productHandler *handler.ProductHandler,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /categories", categoryHandler.Create)
	mux.HandleFunc("GET /categories", categoryHandler.GetAll)
	mux.HandleFunc("GET /categories/{id}", categoryHandler.GetByID)

	
	mux.HandleFunc("POST /products", productHandler.Create)
	mux.HandleFunc("GET /products", productHandler.GetAll)
	mux.HandleFunc("GET /products/{id}", productHandler.GetByID)
	mux.HandleFunc("GET /categories/{category_id}/products", productHandler.GetByCategory)

	return mux
}

