package usecases

import (
	"fmt"
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
		messagebroker.MailerSendSimpleMailKey,
		fmt.Sprintf("%s\n%s\n%s\n%s", m.From, m.To, m.Subject, m.Body),
		10*time.Second,
	)

	if err != nil {
		return err
	}

	return nil
}
