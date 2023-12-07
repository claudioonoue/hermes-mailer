package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hermes-mailer/internal/providers/messagebroker"
)

type App struct {
	Config          *Config
	MessageConsumer messagebroker.Consumer
}

func main() {
	var err error

	config := newConfig()

	consumer, err := messagebroker.NewConsumer(messagebroker.RabbitMQ, config.MessageBrokerURL)
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := &App{
		Config:          config,
		MessageConsumer: consumer,
	}

	go app.ListenForShutdown()

	ch, err := app.MessageConsumer.ConsumeFromMailerExchange()
	if err != nil {
		log.Fatalf(err.Error())
	}

	for msg := range ch {
		exampleConsumerFunc(msg)
	}
}

func (a *App) ListenForShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	fmt.Println("Shutting down consumer...")
	a.MessageConsumer.Close()
	os.Exit(0)
}

func exampleConsumerFunc(msg messagebroker.MailerQueueMessage) {
	fmt.Println("----------------------------")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Message received!")
	fmt.Println("")
	fmt.Println(msg)
	fmt.Println("")
	fmt.Println("Processing...")
	fmt.Println("")
	time.Sleep(5 * time.Second)
	fmt.Println("Done!")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
