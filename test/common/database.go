package common //nolint:revive

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"quote/config"
	"quote/pkg/database"
)

func PostgresClient(t *testing.T, cfg *config.Config) *database.Postgres {
	t.Helper()

	ctx := t.Context()
	db := database.NewPostgres(cfg)

	err := db.Start(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		if err = db.Stop(context.WithoutCancel(ctx)); err != nil {
			log.Fatalf("failed to terminate database connection: %s", err)
		}
	})

	return db
}
