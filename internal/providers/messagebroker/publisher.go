package messagebroker

import (
	"context"
	"errors"
)

// Publisher is an interface that will be implemented by the message broker publisher.
type Publisher interface {
	PublishToMailerQueueWithContext(ctx context.Context, message *MailerQueueMessageBody) error
	Close() error
}

// NewPublisher is a function that will return a new message broker publisher instance.
func NewPublisher(provider string, connectionString string) (Publisher, error) {
	switch provider {
	case RabbitMQ:
		return NewRabbitMQPublisher(connectionString)
	default:
		return nil, errors.New("invalid message broker provider")
	}
}
