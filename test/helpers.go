package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/fx"

	"quote/cmd"
	"quote/config"
	"quote/pkg/database"
)

const (
	dbOccurrence     = 2
	dbStartupTimeout = 5 * time.Second
	waitForAppStart  = 2 * time.Second
	corsMaxAge       = 300
)

func runEssentials(t *testing.T) (*config.Config, *database.Postgres) {
	t.Helper()

	ctx := context.Background()
	dbName := "test_quotes"
	dbUser := "postgres"
	dbPassword := "postgres"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(dbOccurrence).
				WithStartupTimeout(dbStartupTimeout)),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		if err = postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	})

	port := ""
	ports, err := postgresContainer.Ports(ctx)
	for _, v := range ports {
		port = v[0].HostPort
		break
	}

	cfg := &config.Config{
		ServerConfig: config.ServerConfig{
			Addr:              "0.0.0.0:1141",
			CorsMaxAge:        corsMaxAge,
			ReadHeaderTimeout: time.Second,
		},
		PostgresConfig: config.PostgresConfig{
			DSN:           fmt.Sprintf("host=localhost user=postgres password=postgres dbname=test_quotes port=%s sslmode=disable", port),
			MigrationPath: "../migrations",
		},
		QuotesConfig: config.QuotesConfig{
			RandomQuoteChance: 0,
		},
	}

	app := fx.New(cmd.Exec(cfg))
	t.Cleanup(func() {
		if err = app.Stop(ctx); err != nil {
			log.Fatalf("failed to terminate app: %s", err)
		}
	})

	go func() {
		if err = app.Start(ctx); err != nil {
			log.Fatalf("failed to start app: %s", err)
		}
	}()

	time.Sleep(waitForAppStart) // Wait until the app starts

	db := database.NewPostgres(cfg)
	require.NoError(t, db.Start(ctx))
	t.Cleanup(func() {
		if err = db.Stop(ctx); err != nil {
			log.Fatalf("failed to terminate database connection: %s", err)
		}
	})
	return cfg, db
}
