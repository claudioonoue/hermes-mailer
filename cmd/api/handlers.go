package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"hermes-mailer/internal/providers/messagebroker"
)

func (a *App) CheckAPI(c *fiber.Ctx) error {
	return c.JSON(JSONResponse{
		Data:    nil,
		Message: "API is up and running",
	})
}

type Mail struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Type    string `json:"type"`
}

func (a *App) SendMail(c *fiber.Ctx) error {
	var mail Mail

	if err := c.BodyParser(&mail); err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSONResponse{
			Data:    nil,
			Message: err.Error(),
		})
	}

	exchangeKey, err := getExchangeKeyBasedOnMailType(mail.Type)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSONResponse{
			Data:    nil,
			Message: err.Error(),
		})
	}

	err = a.MessagePublisher.PublishToMailerExchange(
		exchangeKey,
		messagebroker.MailerQueueMessageBody{
			From:    mail.From,
			To:      mail.To,
			Subject: mail.Subject,
			Body:    mail.Body,
			Type:    mail.Type,
		},
		10*time.Second,
	)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSONResponse{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.JSON(JSONResponse{
		Data:    nil,
		Message: "Mail sent successfully",
	})
}
