package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"

	"quote/cmd"
	"quote/config"
)

func TestValidateApp(t *testing.T) {
	t.Parallel()

	err := fx.ValidateApp(cmd.Exec(&config.Config{}))
	require.NoError(t, err)
}
