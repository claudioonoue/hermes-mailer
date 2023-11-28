package main

import (
	"fmt"
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

	App.MessageConsumer.ConsumeFromMailerExchange(
		messagebroker.MailerQueue,
		func(msg interface{}) {
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
		})

	var forever chan bool
	<-forever
}
