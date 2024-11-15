package main

import initialize "distributed-key-value-store/server/initialize"

func main() {
	// initialize.InitServer(8080)
	initialize.InitCassandraDB("key_value_store")
}
