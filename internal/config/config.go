package config

import (
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
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		Port:  os.Getenv("PORT"),
		DBURL: os.Getenv("DB_URL"),
	}
	return cfg, nil
}
