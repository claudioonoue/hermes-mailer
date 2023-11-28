package messagebroker

type Consumer struct {
	client *rabbitMQ
}

func NewConsumer(connectionString string) *Consumer {
	client := newRabbitMQ(connectionString)
	client.dial()
	client.openChannel()

	setupMailerConsumer(client)

	return &Consumer{
		client: client,
	}
}

func (mc *Consumer) ConsumeFromMailerExchange(
	queue string,
	callback func(msg interface{}),
) error {
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
		return err
	}

	go func() {
		for msg := range msgs {
			callback(string(msg.Body))
		}
	}()

	return nil
}

func (mc *Consumer) CloseConn() error {
	return nil
}

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
		Key:      MailerSendSimpleMailKey,
		Exchange: MailerExchange,
		NoWait:   false,
	})

}
