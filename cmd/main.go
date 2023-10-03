package main

import (
	"github.com/X3ne/go-scylla-grpc/config"
	"github.com/X3ne/go-scylla-grpc/internal/server"
	"github.com/X3ne/go-scylla-grpc/services"
)

func main() {
	cfg := config.NewConfig()

	services.InitScylla(cfg)

	server.LaunchServer(cfg)

	defer services.CloseScylla()
}
