BINARY_PATH=bin
BINARY_NAME=finC
ENTRYPOINT=./cmd/fin-control/main.go
DOCKER_COMPOSE_FILE=docker/docker-compose.yml

run: setup build
	./$(BINARY_PATH)/$(BINARY_NAME)

setup:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

build: clean
	go build -o $(BINARY_PATH)/$(BINARY_NAME) $(ENTRYPOINT)

teardown: clean
	docker compose -f $(DOCKER_COMPOSE_FILE) stop

clean:
	rm -rf $(BINARY_PATH)
