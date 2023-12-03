API_FOLDER = ./cmd/api
CONSUMER_FOLDER = ./cmd/messageconsumer

run-api:
	go run $(API_FOLDER)

run-consumer:
	go run $(CONSUMER_FOLDER)