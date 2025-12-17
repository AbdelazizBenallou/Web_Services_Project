package messaging

import "github.com/streadway/amqp"

type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(ch *amqp.Channel) *Publisher {
	return &Publisher{ch: ch}
}

func (p *Publisher) Publish(routingKey string, body []byte) error {
	return p.ch.Publish(
		"events",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

