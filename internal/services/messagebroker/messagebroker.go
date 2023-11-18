package messagebroker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}

func Connect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	return conn
}

func OpenChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func DeclareQueue(ch *amqp.Channel, queue Queue) amqp.Queue {
	q, err := ch.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return q
}
