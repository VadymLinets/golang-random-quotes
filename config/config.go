package config

import (
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerConfig
	GormConfig
	QuotesConfig
}

type ServerConfig struct {
	Addr              string        `envconfig:"SERVER_ADDRESS"`
	CorsMaxAge        int           `envconfig:"CORS_MAX_AGE" default:"300"`
	ReadHeaderTimeout time.Duration `envconfig:"SERVER_READ_HEADER_TIMEOUT"`
}

type GormConfig struct {
	Dialect       string `envconfig:"GORM_DIALECT"`
	DSN           string `envconfig:"GORM_DSN"`
	MigrationPath string `envconfig:"GORM_MIGRATION_PATH" default:"migrations"`
}

type QuotesConfig struct {
	RandomQuoteChance float64 `envconfig:"RANDOM_QUOTE_CHANCE"`
}

var (
	once   sync.Once
	config = new(Config)
)

func Get() (*Config, error) {
	var err error
	once.Do(func() {
		var cfg Config
		_ = godotenv.Load()

		if err = envconfig.Process("", &cfg); err != nil {
			return
		}

		config = &cfg
	})

	return config, err
}
