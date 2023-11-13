package main

import "github.com/gofiber/fiber/v2"

func checkAPI(c *fiber.Ctx) error {
	return c.JSON(JSONResponse{
		Data:    nil,
		Message: "API is up and running",
	})
}
