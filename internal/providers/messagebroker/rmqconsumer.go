package messagebroker

// RabbitMQConsumer is a wrapper around the RabbitMQ client.
type RabbitMQConsumer struct {
	client *rabbitMQ
}

// ConsumeFromMailerQueue is a method that will consume a message from the Mailer Exchange.
func (mc *RabbitMQConsumer) ConsumeFromMailerQueue() (<-chan MailerQueueMessage, error) {
	msgs, err := mc.client.consume(rabbitMQConsume{
		Queue:     MailerQueue,
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
			err := fromJSONBytes(msg.Body, &body)
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

// Close is a method that will close the connection to the RabbitMQ.
func (mc *RabbitMQConsumer) Close() error {
	mc.client.closeConn()
	return nil
}

// NewRabbitMQConsumer is a function that will return a new RabbitMQ Consumer instance.
func NewRabbitMQConsumer(connectionString string) (*RabbitMQConsumer, error) {
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
		Name:       MailerQueue,
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

	err = client.queueDeclare(rabbitMQQueue{
		Name:       MailerQueue,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	})
	if err != nil {
		return nil, err
	}

	err = client.queueBind(rabbitMQQueueBind{
		Queue:    MailerQueue,
		Key:      "",
		Exchange: MailerQueue,
		NoWait:   false,
	})
	if err != nil {
		return nil, err
	}

	return &RabbitMQConsumer{
		client: client,
	}, nil
}
