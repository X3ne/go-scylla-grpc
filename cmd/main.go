package main

import (
	"scylla-grpc-adapter/config"
	"scylla-grpc-adapter/internal/server"
)

func main() {
	config := config.NewConfig()

	server.LaunchServer(config)
}
