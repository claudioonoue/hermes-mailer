package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler)
	}

	app.Listen(":3000")
}
