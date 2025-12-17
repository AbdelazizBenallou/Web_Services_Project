package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"product_service/config"
	"product_service/delivery/http/handler"
	"product_service/delivery/http/routes"
	"product_service/messaging"
	"product_service/repository"
	"product_service/usecase"

	"github.com/streadway/amqp"
)

func main() {
	// -------------------------
	// Database
	// -------------------------
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// -------------------------
	// RabbitMQ
	// -------------------------
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("RABBITMQ_URL is not set")
	}
	if !strings.HasPrefix(rabbitURL, "amqp://") && !strings.HasPrefix(rabbitURL, "amqps://") {
		log.Fatal("invalid RABBITMQ_URL: " + rabbitURL)
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// -------------------------
	// Repositories
	// -------------------------
	categoryRepo := repository.NewCategoryPostgres(db)
	productRepo  := repository.NewProductPostgres(db)
	stockRepo    := repository.NewStockPostgres(db)

	// -------------------------
	// UseCases
	// -------------------------
	categoryUC := usecase.NewCategoryUseCase(categoryRepo)
	productUC := usecase.NewProductUseCase(
		productRepo,
		categoryRepo,
		stockRepo,
	)
	stockUC := usecase.NewStockUseCase(stockRepo)

	// -------------------------
	// Rabbit Publisher
	// -------------------------
	publisher := messaging.NewPublisher(ch)

	// -------------------------
	// Rabbit Consumers
	// -------------------------
	if err := messaging.ConsumeOrderCreated(
		ch,
		stockUC,
		publisher,
	); err != nil {
		log.Fatal(err)
	}

	// -------------------------
	// HTTP Handlers
	// -------------------------
	categoryHandler := handler.NewCategoryHandler(categoryUC)
	productHandler := handler.NewProductHandler(productUC)

	router := routes.Setup(categoryHandler, productHandler)

	log.Println("Product Service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

