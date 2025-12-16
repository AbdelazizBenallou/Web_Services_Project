package usecase

type EventPublisher interface {
	Publish(eventName string, payload []byte) error
}

