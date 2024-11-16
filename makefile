MAIN_PKG_DIR = ./server
SRC = ./server/...
BIN_DIR = ./bin
BIN_NAME = runner

.PHONY: build run clean

# builds the project, entry file at MAIN_PKG_DIR
build:
	go build -o $(BIN_DIR)/$(BIN_NAME) ${MAIN_PKG_DIR}

# runs the server using docker
run:
	docker-compose -f ./docker/server.yml up

# deletes built binaries
clean:
	rm -rf $(BIN_DIR)