package usecases

import (
	"errors"
	"time"

	"hermes-mailer/internal/services/messagebroker"
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
	Type    string
}

func (c *Core) SendMail(m Mail) error {
	var err error

	mailType, err := getMailType(m.Type)
	if err != nil {
		return err
	}

	err = c.MessagePublisher.PublishToMailerExchange(
		mailType,
		messagebroker.MailerQueueMessageBody{
			ExternalID: "123",
		},
		10*time.Second,
	)

	if err != nil {
		return err
	}

	return nil
}

func getMailType(t string) (string, error) {
	switch t {
	case messagebroker.MailerSendSimpleMail:
		return messagebroker.MailerSendSimpleMail, nil
	default:
		return "", errors.New("invalid email type")
	}
}
