package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"hermes-mailer/internal/usecases"
)

type app struct {
	Fiber    *fiber.App
	UseCases *usecases.Core
}

var App *app

func main() {
	App = &app{
		Fiber: fiber.New(),
		UseCases: usecases.New(&usecases.Setup{
			MessagePublisherConfig: &usecases.MessageBrokerConfig{
				URL: "amqp://guest:guest@localhost:5672/",
			},
		}),
	}

	for _, route := range App.GetRoutes() {
		App.Fiber.Add(route.Method, route.Path, route.Handler)
	}

	go App.listenForShutdown()

	App.Fiber.Listen(":3000")
}

func (a *app) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down API...")
	a.shutdown()
	os.Exit(0)
}

func (a *app) shutdown() {
	fmt.Println("Cleanup tasks...")

	a.UseCases.Cleanup()
}
