package usecases

import (
	"errors"
	"time"

	"hermes-mailer/internal/providers/messagebroker"
)

const (
	// MailSimple is a constant that represents a simple mail type.
	MailSimple = "simple"
)

// Mail is a struct that represents a mail information.
type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
	Type    string
}

// EnqueueMail is a use case that will enqueue a mail to be sent.
func (c *Core) EnqueueMail(m Mail) error {
	var err error

	exchangeKey, err := getExchangeKeyBasedOnMailType(m.Type)
	if err != nil {
		return err
	}

	err = c.MessagePublisher.PublishToMailerExchange(
		exchangeKey,
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

// getExchangeKeyBasedOnMailType is a helper function that will return the exchange key based on the mail type.
func getExchangeKeyBasedOnMailType(t string) (string, error) {
	switch t {
	case MailSimple:
		return messagebroker.MailerSendSimpleMail, nil
	default:
		return "", errors.New("invalid email type")
	}
}
