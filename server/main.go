package main

import (
	"distributed-key-value-store/config"
	"distributed-key-value-store/server/controller"
	initialize "distributed-key-value-store/server/initialize"
	"distributed-key-value-store/server/service"
	"os"
	"strconv"
)

var KEYSPACE = "key_value_store"
var DB_ADDRESS = "localhost:9042"
var PORT = 8080

func main() {
	config.LoadEnvironmentVariables(".env")
	if os.Getenv("DB_ADDRESS") != "" {
		DB_ADDRESS = os.Getenv("DB_ADDRESS")
	}

	if port, err := strconv.Atoi(os.Getenv("PORT")); err != nil {
		PORT = port
	}

	if os.Getenv("KEY_SPACE") != "" {
		KEYSPACE = os.Getenv("KEY_SPACE")
	}

	session := initialize.InitCassandraDB(KEYSPACE, DB_ADDRESS)
	service := service.NewService(session, KEYSPACE)
	controller := controller.NewController(service)
	initialize.InitServer(PORT, controller)
}
