package messagebroker

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQ struct {
	URL  string
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

type rabbitMQExchange struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

type rabbitMQQueue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type rabbitMQQueueBind struct {
	Queue    string
	Key      string
	Exchange string
	NoWait   bool
	Args     amqp.Table
}

type rabbitMQPublish struct {
	Ctx       context.Context
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}

type rabbitMQConsume struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func newRabbitMQ(URL string) *rabbitMQ {
	return &rabbitMQ{
		URL:  URL,
		Conn: nil,
		Ch:   nil,
	}
}

func (mb *rabbitMQ) dial() error {
	var err error
	mb.Conn, err = amqp.Dial(mb.URL)
	if err != nil {
		return err
	}
	return nil
}

func (mb *rabbitMQ) openChannel() error {
	var err error
	mb.Ch, err = mb.Conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func (mb *rabbitMQ) exchangeDeclare(exchange rabbitMQExchange) error {
	err := mb.Ch.ExchangeDeclare(
		exchange.Name,
		exchange.Type,
		exchange.Durable,
		exchange.AutoDelete,
		exchange.Internal,
		exchange.NoWait,
		exchange.Args,
	)
	if err != nil {
		return err
	}
	return nil
}

func (mb *rabbitMQ) queueDeclare(queue rabbitMQQueue) error {
	_, err := mb.Ch.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	)
	if err != nil {
		return err
	}
	return nil
}

func (mb *rabbitMQ) queueBind(queueBind rabbitMQQueueBind) error {
	err := mb.Ch.QueueBind(
		queueBind.Queue,
		queueBind.Key,
		queueBind.Exchange,
		queueBind.NoWait,
		queueBind.Args,
	)
	if err != nil {
		return err
	}
	return nil
}

func (mb *rabbitMQ) publishWithContext(publish rabbitMQPublish) error {
	return mb.Ch.PublishWithContext(
		publish.Ctx,
		publish.Exchange,
		publish.Key,
		publish.Mandatory,
		publish.Immediate,
		publish.Msg,
	)
}

func (mb *rabbitMQ) consume(consume rabbitMQConsume) (<-chan amqp.Delivery, error) {
	return mb.Ch.Consume(
		consume.Queue,
		consume.Consumer,
		consume.AutoAck,
		consume.Exclusive,
		consume.NoLocal,
		consume.NoWait,
		consume.Args,
	)
}

func (mb *rabbitMQ) closeConn() {
	if mb.Conn != nil {
		mb.Conn.Close()
	}

	if mb.Ch != nil {
		mb.Ch.Close()
	}
}
