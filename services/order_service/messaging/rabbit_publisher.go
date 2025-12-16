package messaging

import "github.com/streadway/amqp"

type RabbitPublisher struct {
	ch *amqp.Channel
}

func NewRabbitPublisher(ch *amqp.Channel) *RabbitPublisher {
	return &RabbitPublisher{ch: ch}
}

func (p *RabbitPublisher) Publish(eventName string, body []byte) error {
	return p.ch.Publish(
		"events",   // exchange
		eventName, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

