package messagebroker

// Consumer is a wrapper around the message broker client.
type Consumer struct {
	client *rabbitMQ
}

// NewConsumer is a function that will return a new Consumer instance.
func NewConsumer(connectionString string) *Consumer {
	client := newRabbitMQ(connectionString)
	client.dial()
	client.openChannel()

	setupMailerConsumer(client)

	return &Consumer{
		client: client,
	}
}

// ConsumeFromMailerExchange is a method that will consume a message from the mailer exchange.
func (mc *Consumer) ConsumeFromMailerExchange(queue string) (<-chan MailerQueueMessage, error) {
	msgs, err := mc.client.consume(rabbitMQConsume{
		Queue:     queue,
		Consumer:  "",
		AutoAck:   true,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	})

	if err != nil {
		return nil, err
	}

	ch := make(chan MailerQueueMessage)

	go func() {
		for msg := range msgs {
			var body MailerQueueMessageBody
			err := body.fromJSONBytes(msg.Body)
			if err != nil {
				continue
			}

			ch <- MailerQueueMessage{
				Key:  msg.RoutingKey,
				Body: body,
			}
		}

		close(ch)
	}()

	return ch, nil
}

// CloseConn is a method that will close the connection to the message broker.
func (mc *Consumer) CloseConn() error {
	mc.client.closeConn()
	return nil
}

// setupMailerConsumer is a helper function that will setup the mailer consumer.
func setupMailerConsumer(client *rabbitMQ) {
	client.exchangeDeclare(rabbitMQExchange{
		Name:       MailerExchange,
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	})

	client.queueDeclare(rabbitMQQueue{
		Name:       MailerQueue,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	})

	client.queueBind(rabbitMQQueueBind{
		Queue:    MailerQueue,
		Key:      MailerSendSimpleMail,
		Exchange: MailerExchange,
		NoWait:   false,
	})

}
