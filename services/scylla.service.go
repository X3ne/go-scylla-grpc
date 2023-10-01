package services

import (
	"log"
	"scylla-grpc-adapter/config"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type ScyllaService struct {
	Session	*gocqlx.Session
}

func InitScylla(cfg *config.Config) *ScyllaService {
	log.Println("Connecting to ScyllaDB...")

	cluster := gocql.NewCluster(cfg.DB.Hosts...)


	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}

	cluster.Keyspace = cfg.DB.Keyspace

	log.Println("Connected to ScyllaDB")

	return &ScyllaService{
		Session: &session,
	}
}
