package usecases

import (
	"fmt"

	"hermes-mailer/internal/services/messagebroker"
)

// Core is the core of the usecases layer. \n
// It contains all the usecases of the application.
type Core struct {
	MessagePublisher *messagebroker.Publisher
}

// Setup is the setup of the usecases layer.
type Setup struct {
	MessagePublisherConfig *MessageBrokerConfig
}

type MessageBrokerConfig struct {
	URL string
}

// New creates a brand new usecases Core with all dependencies initialized.
func New(c *Setup) *Core {
	fmt.Println("Initializing UseCases dependencies...")

	return &Core{
		MessagePublisher: initMessagePublisher(c.MessagePublisherConfig),
	}
}

func (c *Core) Cleanup() {
	fmt.Println("Cleaning UseCases dependencies...")

	c.MessagePublisher.CloseConn()
}

func initMessagePublisher(mbc *MessageBrokerConfig) *messagebroker.Publisher {
	if mbc == nil {
		return nil
	}

	fmt.Println("Initializing MessagePublisher...")

	messagePublisher := messagebroker.NewPublisher(mbc.URL)

	return messagePublisher
}
