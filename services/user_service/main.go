package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"user_service/config"
	"user_service/delivery/http/handler"
	"user_service/delivery/http/routes"
	"user_service/messaging"
	"user_service/repository"
	"user_service/usecase"

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

	publisher := messaging.NewRabbitPublisher(ch)

	// -------------------------
	// Application
	// -------------------------
	userRepo := repository.NewPostgresRepository(db)
	userUC := usecase.NewUserUseCase(userRepo, publisher)
	userHandler := handler.NewUserHandler(userUC)

	router := routes.SetupUserRoutes(userHandler)

	log.Println("User Service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

