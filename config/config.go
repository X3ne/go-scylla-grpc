package config

type Config struct {
	SERVER	ServerConfig
	DB			DatabaseConfig
	JWT			JwtConfig
}

func NewConfig() *Config {
	envs := ValidateEnvs()

	return &Config{
		SERVER: LoadServerConfig(envs),
		DB: LoadDatabaseConfig(envs),
		JWT: LoadJwtConfig(envs),
	}
}
