# --------------------------------------------------------------------------------------
# ---------------------------------- VARIABLES -----------------------------------------
POSTGRES_USER = postgres
POSTGRES_PASSWORD = postgres

RABBITMQ_DEFAULT_USER = rabbitmq
RABBITMQ_DEFAULT_PASS = rabbitmq

BINARY_FOLDER = ./bin
TMP_FOLDER = ./tmp
API_FOLDER = ./cmd/api
API_BINARY_NAME = api_build
CONSUMER_FOLDER = ./cmd/messageconsumer
CONSUMER_BINARY_NAME = messageconsumer_build

# --------------------------------------------------------------------------------------
# ------------------------------------- API --------------------------------------------
.PHONY: api-local
api-local:
	@echo "Running API..."
	@go run $(API_FOLDER)/*.go

.PHONY: api-build
api-build:
	@echo "Building API..."
	@env CGO_ENABLED=0  go build -ldflags="-s -w" -o $(BINARY_FOLDER)/$(API_BINARY_NAME) $(API_FOLDER)
	@echo "Finished building API!"

.PHONY: api-stop
api-stop:
	@echo "Stopping API..."
	@-pkill -SIGTERM -f "./$(BINARY_FOLDER)/$(API_BINARY_NAME)"
	@echo "API stopped!"

.PHONY: api-run
api-run: api-stop api-build
	@echo "Running API..."
	@env ENV=PROD ./$(BINARY_FOLDER)/$(API_BINARY_NAME) &
	@echo "API is running!"

# --------------------------------------------------------------------------------------
# ------------------------------------- Consumer ---------------------------------------
.PHONY: consumer-local
consumer-local:
	@echo "Running Consumer..."
	@go run $(CONSUMER_FOLDER)/*.go

.PHONY: consumer-build
consumer-build:
	@echo "Building Consumer..."
	@env CGO_ENABLED=0  go build -ldflags="-s -w" -o $(BINARY_FOLDER)/$(CONSUMER_BINARY_NAME) $(CONSUMER_FOLDER)
	@echo "Finished building Consumer!"

.PHONY: consumer-stop
consumer-stop:
	@echo "Stopping Consumer..."
	@-pkill -SIGTERM -f "./$(BINARY_FOLDER)/${CONSUMER_BINARY_NAME}"
	@echo "Consumer stopped!"

.PHONY: consumer-run
consumer-run: consumer-stop consumer-build
	@echo "Running Consumer..."
	@env ENV=PROD ./$(BINARY_FOLDER)/$(CONSUMER_BINARY_NAME) &
	@echo "Consumer is running!"

# --------------------------------------------------------------------------------------
# ------------------------------------- Docker -----------------------------------------
.PHONY: docker-start
docker-start:
	@echo "Starting docker..."
	@env POSTGRES_USER=${POSTGRES_USER} POSTGRES_PASSWORD=${POSTGRES_PASSWORD} RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER} RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS} docker compose up -d
	@echo "Docker started!"

.PHONY: docker-stop
docker-stop:
	@echo "Stopping docker..."
	@docker compose down
	@echo "Docker stopped!"

# --------------------------------------------------------------------------------------
# ---------------------------------- Management ----------------------------------------
.PHONY: clean
clean: api-stop consumer-stop
	@echo "Cleaning..."
	@go clean
	@rm -rf $(BINARY_FOLDER)
	@rm -rf $(TMP_FOLDER)
	@echo "Cleaned!"	