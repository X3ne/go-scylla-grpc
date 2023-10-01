package app

import (
	"scylla-grpc-adapter/config"
	"scylla-grpc-adapter/services"
)

type App struct {
	Config *config.Config
	Db		 *services.ScyllaService
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
		Db: services.InitScylla(cfg),
	}
}
