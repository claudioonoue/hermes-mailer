package messagebroker

import "encoding/json"

const (
	// MailerExchange is the name of the mailer exchange.
	MailerExchange = "mailer_exchange"

	// MailerQueue is the name of the mailer queue.
	MailerQueue = "mailer_queue"

	// MailerSendSimpleMail is the routing key for the mailer send simple mail message.
	MailerSendSimpleMail = "mailer_send_simple_mail"
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

// toJSON is a method that will convert the MailerQueueMessageBody to JSON bytes.
func (m MailerQueueMessageBody) toJSONBytes() ([]byte, error) {
	return json.Marshal(m)
}

// fromJSON is a method that will convert the JSON Bytes to MailerQueueMessageBody.
func (m *MailerQueueMessageBody) fromJSONBytes(data []byte) error {
	return json.Unmarshal(data, m)
}
