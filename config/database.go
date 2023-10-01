package config

import "strings"

type DatabaseConfig struct {
	Hosts []string
	Keyspace string
}

func LoadDatabaseConfig(envs map[string]string) DatabaseConfig {

	hosts := strings.Split(envs["DB_HOSTS"], ",")

	return DatabaseConfig{
		Hosts: hosts,
		Keyspace: envs["DB_KEYSPACE"],
	}
}
