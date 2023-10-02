package main

import (
	"scylla-grpc-adapter/config"
	"scylla-grpc-adapter/internal/server"
	"scylla-grpc-adapter/services"
)

func main() {
	cfg := config.NewConfig()

	services.InitScylla(cfg)

	server.LaunchServer(cfg)

	defer services.CloseScylla()
}
