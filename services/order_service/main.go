package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"order_service/config"
	"order_service/delivery/http/handler"
	"order_service/delivery/http/routes"
	"order_service/messaging"
	"order_service/repository"
	"order_service/usecase"

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
	// Rabbit Consumer
	// -------------------------
	userViewRepo := repository.NewUserViewPostgres(db)
	if err := messaging.ConsumeUserRegistered(ch, userViewRepo); err != nil {
		log.Fatal(err)
	}
// -------------------------
// Inventory Consumers
// -------------------------
	if err := messaging.ConsumeInventoryReserved(ch, orderRepo); err != nil {
		log.Fatal(err)
	}

	if err := messaging.ConsumeInventoryFailed(ch, orderRepo); err != nil {
		log.Fatal(err)
	}
	// -------------------------
	// Application
	// -------------------------
	publisher := messaging.NewRabbitPublisher(ch)

	orderRepo := repository.NewPostgresRepository(db)
	orderUC := usecase.NewOrderUseCase(orderRepo, userViewRepo, publisher, )
	orderHandler := handler.NewOrderHandler(orderUC)


	router := routes.SetupOrderRoutes(orderHandler)

	log.Println("Order Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

