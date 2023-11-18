package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"hermes-mailer/internal/usecases"
)

func checkAPI(c *fiber.Ctx) error {
	return c.JSON(JSONResponse{
		Data:    nil,
		Message: "API is up and running",
	})
}

type mail struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func sendMail(c *fiber.Ctx) error {
	var mail mail

	if err := c.BodyParser(&mail); err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSONResponse{
			Data:    nil,
			Message: err.Error(),
		})
	}

	err := App.UseCases.SendMail(usecases.Mail{
		From:    mail.From,
		To:      mail.To,
		Subject: mail.Subject,
		Body:    mail.Body,
	})

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
