package config

import (
	"time"
)

type JwtConfig struct {
	SecretKey			string
	TokenDuration time.Duration
}

func LoadJwtConfig(envs map[string]string) JwtConfig {
	hr, _ := time.ParseDuration(envs["JWT_TOKEN_DURATION"])
	return JwtConfig{
		SecretKey: 	 envs["JWT_SECRET_KEY"],
		TokenDuration: hr,
	}
}
