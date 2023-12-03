package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"hermes-mailer/internal/usecases"
)

type app struct {
	Logger   *logger
	Fiber    *fiber.App
	UseCases *usecases.Core
}

var App *app

func main() {
	var err error

	logger, err := newLogger()
	if err != nil {
		panic(err)
	}

	logger.info("Configuring API...")

	logger.info("Initializing UseCases...")

	useCasesCore, err := usecases.New(&usecases.Setup{
		MessagePublisherConfig: &usecases.MessageBrokerConfig{
			URL: "amqp://guest:guest@localhost:5672/",
		},
	})
	if err != nil {
		panic(err)
	}

	App = &app{
		Logger: logger,
		Fiber: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
		UseCases: useCasesCore,
	}

	logger.info("Binding routes...")
	for _, route := range App.GetRoutes() {
		App.Fiber.Add(route.Method, route.Path, route.Handler)
	}

	go App.listenForShutdown()

	logger.info("Serving API...")
	App.Fiber.Listen(":3000")
}

func (a *app) listenForShutdown() {
	a.Logger.info("Listening for shutdown...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.shutdown()
	os.Exit(0)
}

func (a *app) shutdown() {
	a.UseCases.Cleanup()
}
