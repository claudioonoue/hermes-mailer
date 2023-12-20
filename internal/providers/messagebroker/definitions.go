package messagebroker

const (
	// RabbitMQ is the name of the RabbitMQ message broker provider.
	RabbitMQ = "rabbitmq"

	// MailerQueue is the name of the mailer queue.
	MailerQueue = "mailer_queue"
)

// MailerQueueMessage is a struct that represents a message that will be published to the mailer exchange.
type MailerQueueMessage struct {
	Key  string
	Body MailerQueueMessageBody
}

// MailerQueueMessageBody is a struct that represents the body of a message that will be published to the mailer exchange.
type MailerQueueMessageBody struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Type    string `json:"type"`
}
