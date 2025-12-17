package messaging

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"product_service/usecase"
)

func ConsumeOrderCreated(
	ch *amqp.Channel,
	stockUC usecase.StockUseCase,
	publisher *Publisher,
) error {

	q, err := ch.QueueDeclare(
		"order_created_queue",
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
		"order.created",
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
			var event OrderCreatedEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Println("invalid order.created:", err)
				continue
			}

			err := stockUC.ReserveForOrder(event.Items)
			if err != nil {
				fail := InventoryFailedEvent{
					OrderID: event.OrderID,
					Reason:  err.Error(),
				}

				data, _ := json.Marshal(fail)
				_ = publisher.Publish("inventory.failed", data)
				continue
			}

			ok := InventoryReservedEvent{
				OrderID: event.OrderID,
				Items:   event.Items,
			}

			data, _ := json.Marshal(ok)
			_ = publisher.Publish("inventory.reserved", data)
		}
	}()

	return nil
}

