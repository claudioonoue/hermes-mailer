package messagebroker

import (
	"errors"
)

// Consumer is an interface that will be implemented by the message broker consumer.
type Consumer interface {
	ConsumeFromMailerExchange() (<-chan MailerQueueMessage, error)
	Close() error
}

// NewConsumer is a function that will return a new message broker consumer instance.
func NewConsumer(provider string, connectionString string) (Consumer, error) {
	switch provider {
	case RabbitMQ:
		return NewRabbitMQConsumer(connectionString)
	default:
		return nil, errors.New("invalid message broker provider")
	}
}
