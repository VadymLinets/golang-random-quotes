//go:build mutation

package test

import (
	"os"
	"testing"

	"github.com/gtramontina/ooze"
	"github.com/stretchr/testify/require"
)

func TestMutation(t *testing.T) {
	err := os.Setenv("MUTATION_TESTING", "true")
	require.NoError(t, err)

	ooze.Release(
		t,
		ooze.WithRepositoryRoot("../"),
		ooze.WithMinimumThreshold(0.9),
		ooze.Parallel(),
		ooze.IgnoreSourceFiles("^((test|server|pkg|cmd|config|tools|internal/heartbeat)/.+|main|.+structs|.+interfaces|.+errors)\\.go$"),
		ooze.ForceColors(),
	)
}
