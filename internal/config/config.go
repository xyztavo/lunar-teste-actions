package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConfig)

type Config struct {
	Port  string
	DBURL string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:  os.Getenv("PORT"),
		DBURL: os.Getenv("DB_URL"),
	}

	if cfg.Port == "" || cfg.DBURL == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
