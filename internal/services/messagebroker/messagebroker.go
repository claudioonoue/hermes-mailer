package messagebroker

import "encoding/json"

const (
	MailerExchange = "mailer_exchange"

	MailerQueue = "mailer_queue"

	MailerSendSimpleMail = "mailer_send_simple_mail"
)

type MailerQueueMessage struct {
	Key  string
	Body MailerQueueMessageBody
}

type MailerQueueMessageBody struct {
	ExternalID string `json:"externalID"`
	EmailType  string `json:"emailType"`
}

func (m MailerQueueMessageBody) toJSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MailerQueueMessageBody) fromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}
