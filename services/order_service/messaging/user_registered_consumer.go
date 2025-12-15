package messaging

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"order_service/repository"
)

type UserRegisteredEvent struct {
	UserID int64 `json:"user_id"`
}

func ConsumeUserRegistered(
	ch *amqp.Channel,
	userViewRepo *repository.UserViewPostgres,
) error {

	q, err := ch.QueueDeclare(
		"user_registered_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"user.registered",
		"events",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var event UserRegisteredEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Println("invalid user.registered event:", err)
				continue
			}

			if err := userViewRepo.Insert(event.UserID); err != nil {
				log.Println("failed to insert into user_view:", err)
				continue
			}

			log.Println("user_view inserted for user_id:", event.UserID)
		}
	}()

	return nil
}

