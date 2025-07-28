package common

import (
	"fmt"
	"testing"
	"time"

	"quote/config"
	"quote/test/docker"
)

const (
	corsMaxAge = 300
)

func GetConfig(t *testing.T, postgresConfig docker.PostgresConfig) *config.Config {
	t.Helper()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresConfig.Host,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Database,
		postgresConfig.Port,
	)

	return &config.Config{
		ServerConfig: config.ServerConfig{
			Type:              config.HTTP,
			Addr:              "0.0.0.0:45245",
			CorsMaxAge:        corsMaxAge,
			ReadHeaderTimeout: time.Second,
		},
		PostgresConfig: config.PostgresConfig{
			DSN:           dsn,
			MigrationPath: postgresConfig.MigrationPath,
		},
		QuotesConfig: config.QuotesConfig{
			RandomQuoteChance: 0,
		},
	}
}
