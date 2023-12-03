package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	MessageBrokerURL string
}

const (
	EnvLocal = "LOCAL"
	EnvDev   = "DEV"
	EnvProd  = "PROD"
)

func newConfig() *Config {
	godotenv.Load()

	return &Config{
		Env:              os.Getenv("ENV"),
		MessageBrokerURL: os.Getenv("MESSAGE_BROKER_URL"),
	}
}
