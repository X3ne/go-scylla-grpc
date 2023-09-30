package config

type DatabaseConfig struct {
	Host string
	Port string
	Keyspace string
}

func LoadDatabaseConfig(envs map[string]string) DatabaseConfig {
	return DatabaseConfig{
		Host: envs["DB_HOST"],
		Port: envs["DB_PORT"],
		Keyspace: envs["DB_KEYSPACE"],
	}
}
