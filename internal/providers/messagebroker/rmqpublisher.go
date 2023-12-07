package messagebroker

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// RabbitMQPublisher is a wrapper around the RabbitMQ client.
type RabbitMQPublisher struct {
	client *rabbitMQ
}

// PublishToMailerExchangeWithContext is a method that will publish a message to Mailer Exchange.
func (mp *RabbitMQPublisher) PublishToMailerExchangeWithContext(ctx context.Context, key string, message *MailerQueueMessageBody) error {
	var err error

	jsonMessage, err := toJSONBytes(&message)
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

// Close is a method that will close the connection to the RabbitMQ.
func (mp *RabbitMQPublisher) Close() error {
	mp.client.closeConn()
	return nil
}

// NewRabbitMQPublisher is a function that will return a new RabbitMQ Publisher instance.
func NewRabbitMQPublisher(connectionString string) (*RabbitMQPublisher, error) {
	var err error

	client := newRabbitMQ(connectionString)

	err = client.dial()
	if err != nil {
		return nil, err
	}

	err = client.openChannel()
	if err != nil {
		return nil, err
	}

	err = client.exchangeDeclare(rabbitMQExchange{
		Name:       MailerExchange,
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	})
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{
		client: client,
	}, nil
}
