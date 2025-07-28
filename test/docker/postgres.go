package docker

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbName     = "test_quotes"
	dbUser     = "postgres"
	dbPassword = "postgres"

	dbOccurrence     = 2
	dbStartupTimeout = 5 * time.Second
)

type PostgresConfig struct {
	Host          string
	User          string
	Password      string
	Database      string
	Port          string
	MigrationPath string
}

func StartPostgres(t *testing.T) PostgresConfig {
	t.Helper()
	t.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	ctx := t.Context()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:latest",
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
		if err = postgresContainer.Terminate(context.WithoutCancel(ctx)); err != nil {
			log.Fatalf("failed to terminate postgres container: %s", err)
		}
	})

	host := "localhost"
	port := "5432"

	if os.Getenv("RUNS_IN_CI") == "true" {
		host, err = postgresContainer.ContainerIP(ctx)
		require.NoError(t, err)
	} else {
		ports, err := postgresContainer.Ports(ctx)
		require.NoError(t, err)

		for _, v := range ports {
			port = v[0].HostPort
			break
		}
	}

	return PostgresConfig{
		Host:          host,
		User:          dbUser,
		Password:      dbPassword,
		Database:      dbName,
		Port:          port,
		MigrationPath: "../migrations",
	}
}
