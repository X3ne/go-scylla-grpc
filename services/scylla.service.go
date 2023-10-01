package services

import (
	"fmt"
	"log"
	"scylla-grpc-adapter/config"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func createUserTable(session gocqlx.Session, keyspace string) {
	if err := session.ExecStmt(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s.users (
			id text,
			value text,
			PRIMARY KEY (id, value)
		) WITH CLUSTERING ORDER BY (value ASC)`,
		keyspace,
	)); err != nil {
		log.Fatal("Error creating table", err)
	}


}

func initKeyspace(session gocqlx.Session, keyspace string) {
	if err := session.ExecStmt(fmt.Sprintf(
		`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`,
		keyspace,
	)); err != nil {
		log.Fatal("Error creating keyspace", err)
	}
}

type ScyllaService struct {
	Session	*gocqlx.Session
}

func InitScylla(cfg *config.Config) *ScyllaService {
	log.Println("Connecting to ScyllaDB...")

	cluster := gocql.NewCluster(cfg.DB.Hosts...)

	cluster.Keyspace = cfg.DB.Keyspace

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}

	initKeyspace(session, cfg.DB.Keyspace)

	createUserTable(session, cfg.DB.Keyspace)


	log.Println("Connected to ScyllaDB")

	return &ScyllaService{
		Session: &session,
	}
}
