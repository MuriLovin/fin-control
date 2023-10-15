BINARY_PATH=bin
BINARY_NAME=finC
ENTRYPOINT=./cmd/fin-control/main.go
DOCKER_COMPOSE_FILE=docker/docker-compose.yml

compose-up:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

compose-stop:
	docker compose -f $(DOCKER_COMPOSE_FILE) stop

build: clean
	go build -o $(BINARY_PATH)/$(BINARY_NAME) $(ENTRYPOINT)

clean:
	rm -f bin/$(BINARY_NAME)

run: build
	./$(BINARY_PATH)/$(BINARY_NAME)