# --------------------------------------------------------------------------------------
# ---------------------------------- VARIABLES -----------------------------------------
POSTGRES_USER = postgres
POSTGRES_PASSWORD = postgres

RABBITMQ_DEFAULT_USER = rabbitmq
RABBITMQ_DEFAULT_PASS = rabbitmq

API_FOLDER = ./cmd/api
API_BINARY_NAME = api_build
CONSUMER_FOLDER = ./cmd/messageconsumer
CONSUMER_BINARY_NAME = messageconsumer_build

# --------------------------------------------------------------------------------------
# ------------------------------------- API --------------------------------------------
build-api:
	@echo "Building API..."
	@env CGO_ENABLED=0  go build -ldflags="-s -w" -o $(API_BINARY_NAME) $(API_FOLDER)
	@echo "Finished building API!"

run-api: stop-api build-api
	@echo "Running API..."
	@./$(API_BINARY_NAME) &
	@echo "API is running!"

stop-api:
	@echo "Stopping API..."
	@-pkill -SIGTERM -f "./${API_BINARY_NAME}"
	@echo "API stopped!"

# --------------------------------------------------------------------------------------
# ------------------------------------- Consumer ---------------------------------------
build-consumer:
	@echo "Building Consumer..."
	@env CGO_ENABLED=0  go build -ldflags="-s -w" -o $(CONSUMER_BINARY_NAME) $(CONSUMER_FOLDER)
	@echo "Finished building Consumer!"

run-consumer: stop-consumer build-consumer
	@echo "Running Consumer..."
	@./$(CONSUMER_BINARY_NAME) &
	@echo "Consumer is running!"

stop-consumer:
	@echo "Stopping Consumer..."
	@-pkill -SIGTERM -f "./${CONSUMER_BINARY_NAME}"
	@echo "Consumer stopped!"

# --------------------------------------------------------------------------------------
# ------------------------------------- Docker -----------------------------------------
start-docker:
	@echo "Starting docker..."
	@POSTGRES_USER=${POSTGRES_USER} POSTGRES_PASSWORD=${POSTGRES_PASSWORD} RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER} RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS} docker compose up -d
	@echo "Docker started!"

stop-docker:
	@echo "Stopping docker..."
	@docker compose down
	@echo "Docker stopped!"

# --------------------------------------------------------------------------------------
# ---------------------------------- Management ----------------------------------------
clean: stop-api stop-consumer
	@echo "Cleaning..."
	@go clean
	@rm -f $(API_BINARY_NAME)
	@rm -f $(CONSUMER_BINARY_NAME)
	@echo "Cleaned!"	