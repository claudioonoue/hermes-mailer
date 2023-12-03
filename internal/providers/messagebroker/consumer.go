package messagebroker

// Consumer is a wrapper around the message broker client.
type Consumer struct {
	client *rabbitMQ
}

// NewConsumer is a function that will return a new Consumer instance.
func NewConsumer(connectionString string) (*Consumer, error) {
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

	err = setupMailerConsumer(client)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		client: client,
	}, nil
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
func (mc *Consumer) CloseConn() {
	mc.client.closeConn()
}

// setupMailerConsumer is a helper function that will setup the mailer consumer.
func setupMailerConsumer(client *rabbitMQ) error {
	var err error

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
		return err
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
		return err
	}

	err = client.queueBind(rabbitMQQueueBind{
		Queue:    MailerQueue,
		Key:      MailerSendSimpleMail,
		Exchange: MailerExchange,
		NoWait:   false,
	})
	if err != nil {
		return err
	}

	return nil
}
