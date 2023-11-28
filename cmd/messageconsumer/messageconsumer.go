package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hermes-mailer/internal/services/messagebroker"
)

type app struct {
	MessageConsumer *messagebroker.Consumer
}

func main() {
	App := &app{
		MessageConsumer: messagebroker.NewConsumer("amqp://guest:guest@localhost:5672/"),
	}

	go App.listenForShutdown()

	ch, err := App.MessageConsumer.ConsumeFromMailerExchange(messagebroker.MailerQueue)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case msg := <-ch:
			exampleConsumerFunc(msg)

		default:
			time.Sleep(2 * time.Second)
			fmt.Println("Waiting for messages...")
		}
	}
}

func (a *app) listenForShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	fmt.Println("Shutting down consumer...")
	a.MessageConsumer.CloseConn()
	os.Exit(0)
}

func exampleConsumerFunc(msg string) {
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
