package messagebroker

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	client *rabbitMQ
}

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

func (mp *Publisher) PublishToMailerExchange(key string, message MailerQueueMessageBody, timeout time.Duration) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	jsonMessage, err := message.toJSON()
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

func (mp *Publisher) CloseConn() error {
	mp.client.closeConn()
	return nil
}
