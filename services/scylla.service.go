package services

import (
	"log"
	"scylla-grpc-adapter/config"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var (
	session	gocqlx.Session
)

func InitScylla(cfg *config.Config) {
	log.Println("Connecting to ScyllaDB...")

	cluster := gocql.NewCluster(cfg.DB.Hosts...)

	cluster.Keyspace = cfg.DB.Keyspace

	var err error
	session, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to ScyllaDB")
}

func CloseScylla() {
	log.Println("Closing ScyllaDB connection...")
	session.Close()
	log.Println("ScyllaDB connection closed")
}
