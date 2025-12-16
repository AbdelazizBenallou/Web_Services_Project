package main

import (
	"log"
	"net/http"

	"product_service/config"
	"product_service/delivery/http/handler"
	"product_service/delivery/http/routes"
	"product_service/repository"
	"product_service/usecase"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repositories
	categoryRepo := repository.NewCategoryPostgres(db)
	productRepo  := repository.NewProductPostgres(db)
	stockRepo    := repository.NewStockPostgres(db)

	// UseCases
	categoryUC := usecase.NewCategoryUseCase(categoryRepo)
	productUC  := usecase.NewProductUseCase(
		productRepo,
		categoryRepo,
		stockRepo,
	)

	// Handlers
	categoryHandler := handler.NewCategoryHandler(categoryUC)
	productHandler  := handler.NewProductHandler(productUC)

	// Router
	router := routes.Setup(categoryHandler, productHandler)

	log.Println("Product Service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

