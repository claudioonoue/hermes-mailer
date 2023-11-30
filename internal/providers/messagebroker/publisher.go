package messagebroker

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// Publisher is a wrapper around the message broker client.
type Publisher struct {
	client *rabbitMQ
}

// NewPublisher is a function that will return a new Publisher instance.
func NewPublisher(connectionString string) *Publisher {
	client := newRabbitMQ(connectionString)
	client.dial()
	client.openChannel()

	client.exchangeDeclare(rabbitMQExchange{
		Name:       MailerExchange,
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	})

	return &Publisher{
		client: client,
	}
}

// PublishToMailerExchange is a method that will publish a message to the mailer exchange.
func (mp *Publisher) PublishToMailerExchange(key string, message MailerQueueMessageBody, timeout time.Duration) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	jsonMessage, err := message.toJSONBytes()
	if err != nil {
		return err
	}

	err = mp.client.publishWithContext(rabbitMQPublish{
		Ctx:       ctx,
		Exchange:  MailerExchange,
		Key:       key,
		Mandatory: false,
		Immediate: false,
		Msg: amqp091.Publishing{
			ContentType: "text/plain",
			Body:        jsonMessage,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// CloseConn is a method that will close the connection to the message broker.
func (mp *Publisher) CloseConn() error {
	mp.client.closeConn()
	return nil
}
