package main

import (
	"scylla-grpc-adapter/config"
	"scylla-grpc-adapter/internal/app"
	"scylla-grpc-adapter/internal/server"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg)

	server.LaunchServer(cfg, app)
}
