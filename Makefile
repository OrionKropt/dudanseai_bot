APP_NAME := dudanseai_bot
CMD_DIR  := ./cmd/app
BIN_DIR  := ./bin
BIN_FILE := $(BIN_DIR)/$(APP_NAME)
BIN_FILE_DEV := ./tmp/$(APP_NAME)

.PHONY: fmt vet build build-dev clean run test
.PHONY: lint dev tidy
.PHONY: dev-docker-compose-up dev-docker-compose-down
.PHONY: prod-docker-compose-up prod-docker-compose-down

all: build

dev-docker-compose-up:
	sudo docker compose -f docker-compose.dev.yml --env-file .env.dev up --build

dev-docker-compose-down:
	sudo docker compose -f docker-compose.dev.yml --env-file .env.dev down

prod-docker-compose-up:
	sudo docker compose -f docker-compose.prod.yml --env-file .env up -d --build

prod-docker-compose-down:
	sudo docker compose -f docker-compose.prod.yml --env-file .env down

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_FILE) $(CMD_DIR)

build-dev:
	go build -buildvcs=false -o $(BIN_FILE_DEV) $(CMD_DIR)

run: build
	go run $(CMD_DIR)

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	@echo "→ Cleaning binary"
	@rm -rf $(BIN_DIR)
