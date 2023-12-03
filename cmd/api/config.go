package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	APIPort          string
	MessageBrokerURL string
}

func newConfig() (*Config, error) {
	godotenv.Load()

	return &Config{
		Env:              os.Getenv("ENV"),
		APIPort:          os.Getenv("API_PORT"),
		MessageBrokerURL: os.Getenv("MESSAGE_BROKER_URL"),
	}, nil
}
