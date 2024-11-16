# Getting Started

This project provides a persistent and distributed key-value storage service, with support for client context deadlines for better scalability.

The project contains server and client code. To start off, make sure to run the server using docker. The config files for docker and docker compose can be found in `/docker`.
```bash
docker-compose -f ./docker/server.yml up
# OR
make run
```

To run the server without using docker, make sure to have a running instance of Cassandra DB and change the address in the environment variables, whose sample can be found in `/env.sample`. The server can then be built using the following command:
```bash
go build -o ./bin/runner ./server
# OR
make build
```

Then simply run the executable file.
```bash
./bin/runner
```

## Development

To facilitate development, a hot-reload tool for golang Fresh has been used. This will re-build and re-run the program when changes are made. The config settings for Fresh are in `/fresh.yaml`. To use fresh, simply run the command:
```bash
fresh
```
To run the service without using Fresh, use either of the options provided in the previous section.

# Project Structure
```
├───client
│       client.go
│
├───config
│       env.go
│
├───docker
│       server.dockerfile
│       server.yml
│       wait-for-it.sh
│
└───server
    │   main.go
    │
    ├───controller
    │       controller.go
    │
    ├───domain
    │       domain.go
    │
    ├───initialize
    │       init_cassandra_db.go
    │       init_server.go
    │
    └───service
            service.go
```

> `client/` contains the client-side code for interacting with the key-value store.
 - `client.go`: Main client logic.

> `config/` contains configuration-related code.
 - `env.go`: Manages environment variables.

> `docker/` contains Docker-related files.
 - `server.dockerfile`: Dockerfile for the server.
 - `server.yml`: Docker Compose file for the server.
 - `wait-for-it.sh`: Script to wait for services to be ready.

> `server/` contains the server-side code for the key-value store.
 - `main.go`: Entry point for the server application.
 - `controller/`: Manages the server controllers.
    - `controller.go`: Main controller logic.
 - `domain/`: Contains the core business logic and domain models.
    - `domain.go`: Main domain logic.
 - `initialize/`: Contains initialization code for the server and Cassandra DB.
    - `init_cassandra_db.go`: Cassandra DB initialization logic.
    - `init_server.go`: Server initialization logic.
 - `service/`: Contains the service layer code.
    - `service.go`: Main service logic.

# Usage

The project supports the following simple commands using sockets with context deadlines to improve the scalability of the service. After connecting to the service, it will expect one of the commands listed below. The sample client in `client/client.go` uses the terminal to simulate the interaction between the server and potential clients. Here is the list of commands:

| Feature | Command |
| - | - | 
| Get key | `GET <KEY>` |
| Put value in key | `PUT <KEY> <VALUE>` |
| List all key value pairs | `LIST` |
| Delete key-value pair | `DELETE <KEY>` |