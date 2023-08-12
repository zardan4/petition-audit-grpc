package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type MongoConfig struct {
	ConnectionLine string
	Database       string
}

type ServerConfig struct {
	Port int
}

type Config struct {
	DB     MongoConfig
	Server ServerConfig
}

func New() (*Config, error) {
	godotenv.Load()

	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
