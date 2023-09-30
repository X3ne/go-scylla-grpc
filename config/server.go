package config

type ServerConfig struct {
	Host string
	Port string
}

func LoadServerConfig(envs map[string]string) ServerConfig {
	return ServerConfig{
		Host: envs["HOST"],
		Port: envs["PORT"],
	}
}
