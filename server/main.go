package main

import (
	"distributed-key-value-store/server/controller"
	initialize "distributed-key-value-store/server/initialize"
	"distributed-key-value-store/server/service"
)

var KEYSPACE = "key_value_store"

func main() {
	session := initialize.InitCassandraDB(KEYSPACE)
	service := service.NewService(session, KEYSPACE)
	controller := controller.NewController(service)
	initialize.InitServer(8080, controller)
}
