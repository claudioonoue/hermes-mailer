package usecases

import (
	"time"

	"hermes-mailer/internal/services/messagebroker"
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (c *Core) SendMail(m Mail) error {
	err := c.MessagePublisher.PublishToMailerExchange(
		messagebroker.MailerSendSimpleMail,
		messagebroker.MailerQueueMessageBody{
			ExternalID: "123",
			EmailType:  "simple",
		},
		10*time.Second,
	)

	if err != nil {
		return err
	}

	return nil
}
