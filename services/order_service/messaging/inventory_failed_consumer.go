package messaging

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"order_service/domain"
	"order_service/repository"
)

func ConsumeInventoryFailed(
	ch *amqp.Channel,
	orderRepo repository.OrderRepository,
) error {

	q, err := ch.QueueDeclare(
		"inventory_failed_queue",
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
		"inventory.failed",
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
			var event InventoryFailedEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Println("invalid inventory.failed:", err)
				continue
			}

			err := orderRepo.UpdateStatus(
				event.OrderID,
				domain.StatusCancelled,
			)
			if err != nil {
				log.Println("failed to cancel order:", err)
			}
		}
	}()

	return nil
}

