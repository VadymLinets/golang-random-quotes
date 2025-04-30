package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"

	"quote/app"
	"quote/config"
)

func TestValidateApp(t *testing.T) {
	t.Parallel()

	err := fx.ValidateApp(app.Exec(&config.Config{}))
	require.NoError(t, err)
}
