package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Port     string `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"postgres"`
		Password string `envconfig:"DB_PASSWORD" default:"postgres"`
		Name     string `envConfig:"DB_NAME" default:"greendeco"`
		SSLMode  string `envConfig:"SSL_MODE" default:"disable"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
