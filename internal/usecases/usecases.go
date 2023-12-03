package usecases

import (
	"hermes-mailer/internal/providers/messagebroker"
)

// Core is the core struct of the usecases layer.
// It contains all the usecases methods.
type Core struct {
	MessagePublisher *messagebroker.Publisher
}

// Setup is the setup struct of the usecases layer.
// It contains all the information necessary to initialize the usecases dependencies.
type Setup struct {
	MessagePublisherConfig *MessageBrokerConfig
}

// MessageBrokerConfig is the configuration struct for the MessageBroker.
type MessageBrokerConfig struct {
	URL string
}

// New creates a brand new usecases Core with all dependencies initialized.
// It receives a pointer to a Setup struct with all the information necessary to initialize the usecases dependencies.
//
// Passing a nil value to any field in the struct will result in the dependency in question not being initialized.
// Example: if you pass a nil value to the MessagePublisherConfig field, the MessagePublisher dependency will not be initialized.
//
// It returns a pointer to the new instantiated Core.
func New(c *Setup) (*Core, error) {
	var err error
	core := &Core{}

	if c.MessagePublisherConfig != nil {
		core.MessagePublisher, err = messagebroker.NewPublisher(c.MessagePublisherConfig.URL)
		if err != nil {
			return nil, err
		}
	}

	return core, nil
}

// Cleanup cleans up all the usecases dependencies.
func (c *Core) Cleanup() {
	c.MessagePublisher.CloseConn()
}
