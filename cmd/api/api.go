package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"hermes-mailer/internal/providers/messagebroker"
)

// App is the main application struct.
type App struct {
	Config           *Config
	Logger           *Logger
	Fiber            *fiber.App
	MessagePublisher messagebroker.Publisher
}

func main() {
	var err error

	config := newConfig()

	isProd := config.Env == EnvProd

	logger, err := newLogger(isProd)
	if err != nil {
		log.Fatalf(err.Error())
	}

	publisher, err := messagebroker.NewPublisher(messagebroker.RabbitMQ, config.MessageBrokerURL)
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := &App{
		Config: config,
		Logger: logger,
		Fiber: fiber.New(fiber.Config{
			DisableStartupMessage: isProd,
		}),
		MessagePublisher: publisher,
	}

	for _, route := range app.GetRoutes() {
		app.Fiber.Add(route.Method, route.Path, route.Handler)
	}

	go app.ListenForShutdown()

	app.Fiber.Listen(fmt.Sprintf(":%s", config.APIPort))
}

// ListenForShutdown listens for a shutdown signal and calls the Shutdown method.
func (a *App) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Shutdown()
	os.Exit(0)
}

// Shutdown calls the cleanup methods for the application.
func (a *App) Shutdown() {
	a.Logger.Sync()
	a.MessagePublisher.Close()
}
