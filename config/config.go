package config

import (
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/go-viper/mapstructure/v2"
	"github.com/joho/godotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
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
	Type              ServerType    `koanf:"SERVER_TYPE" default:"http"`
	Addr              string        `koanf:"SERVER_ADDRESS"`
	CorsMaxAge        int           `koanf:"CORS_MAX_AGE" default:"300"`
	ReadHeaderTimeout time.Duration `koanf:"SERVER_READ_HEADER_TIMEOUT"`
}

type PostgresConfig struct {
	DSN           string `koanf:"POSTGRES_DSN"`
	MigrationPath string `koanf:"POSTGRES_MIGRATION_PATH" default:"migrations"`
}

type QuotesConfig struct {
	RandomQuoteChance float64 `koanf:"RANDOM_QUOTE_CHANCE"`
}

func Get() (*Config, error) {
	_ = godotenv.Load()

	cfg := new(Config)
	if err := defaults.Set(cfg); err != nil {
		return nil, err
	}

	k := koanf.New(".")
	if err := k.Load(env.Provider("", ".", strings.ToUpper), nil); err != nil {
		return nil, err
	}

	if err := k.UnmarshalWithConf("", nil, koanf.UnmarshalConf{
		Tag: "koanf",
		DecoderConfig: &mapstructure.DecoderConfig{
			DecodeHook: mapstructure.ComposeDecodeHookFunc(
				mapstructure.StringToTimeDurationHookFunc(),
				mapstructure.StringToSliceHookFunc(","),
			),
			Result:           &cfg,
			WeaklyTypedInput: true,
			Squash:           true,
		},
	}); err != nil {
		return nil, err
	}

	return cfg, nil
}
