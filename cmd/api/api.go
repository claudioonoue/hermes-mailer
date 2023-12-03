package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"hermes-mailer/internal/usecases"
)

type App struct {
	Config   *Config
	Logger   *Logger
	Fiber    *fiber.App
	UseCases *usecases.Core
}

func main() {
	var err error

	config, err := newConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	isProd := config.Env == "prod"

	logger, err := newLogger(isProd)
	if err != nil {
		log.Fatalf(err.Error())
	}

	useCasesCore, err := usecases.New(&usecases.Setup{
		MessagePublisherConfig: &usecases.MessageBrokerConfig{
			URL: config.MessageBrokerURL,
		},
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := &App{
		Config: config,
		Logger: logger,
		Fiber: fiber.New(fiber.Config{
			DisableStartupMessage: isProd,
		}),
		UseCases: useCasesCore,
	}

	for _, route := range app.GetRoutes() {
		app.Fiber.Add(route.Method, route.Path, route.Handler)
	}

	go app.listenForShutdown()

	app.Fiber.Listen(fmt.Sprintf(":%s", config.APIPort))
}

func (a *App) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.shutdown()
	os.Exit(0)
}

func (a *App) shutdown() {
	a.UseCases.Cleanup()
	a.Logger.Sync()
}
