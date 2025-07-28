package common

import (
	"context"
	"log"
	"testing"
	"time"

	"go.uber.org/fx"

	"quote/app"
	"quote/config"
)

const waitForAppStart = 2 * time.Second

func RunApp(t *testing.T, cfg *config.Config) {
	t.Helper()

	ctx := t.Context()
	quotesApp := fx.New(app.Exec(cfg))

	t.Cleanup(func() {
		if err := quotesApp.Stop(context.WithoutCancel(ctx)); err != nil {
			log.Fatalf("failed to terminate app: %s", err)
		}
	})

	go func() {
		if err := quotesApp.Start(ctx); err != nil {
			log.Fatalf("failed to start app: %s", err)
		}
	}()

	time.Sleep(waitForAppStart) // Wait until the app starts
}
