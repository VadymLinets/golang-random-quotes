package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
	PostgresConfig
	QuotesConfig
}

type ServerType string

const (
	HTTP ServerType = "http"
	GRPC ServerType = "grpc"
)

type ServerConfig struct {
	Type              ServerType    `env:"SERVER_TYPE" envDefault:"http"`
	Addr              string        `env:"SERVER_ADDRESS"`
	CorsMaxAge        int           `env:"CORS_MAX_AGE" envDefault:"300"`
	ReadHeaderTimeout time.Duration `env:"SERVER_READ_HEADER_TIMEOUT"`
}

type PostgresConfig struct {
	DSN           string `env:"POSTGRES_DSN"`
	MigrationPath string `env:"POSTGRES_MIGRATION_PATH" envDefault:"migrations"`
}

type QuotesConfig struct {
	RandomQuoteChance float64 `env:"RANDOM_QUOTE_CHANCE"`
}

func Get() (*Config, error) {
	_ = godotenv.Load()

	cfg, err := env.ParseAs[Config]()

	return &cfg, err
}
