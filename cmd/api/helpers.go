package main

import (
	"errors"

	"hermes-mailer/internal/providers/messagebroker"
)

type JSONResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

const (
	// MailSimple is a constant that represents a simple mail type.
	MailSimple = "simple"
)

// getExchangeKeyBasedOnMailType is a helper function that will return the exchange key based on the mail type.
func getExchangeKeyBasedOnMailType(t string) (string, error) {
	switch t {
	case MailSimple:
		return messagebroker.MailerSendSimpleMail, nil
	default:
		return "", errors.New("invalid email type")
	}
}
