//go:build mutation

package test

import (
	"testing"

	"github.com/gtramontina/ooze"
)

func TestMutation(t *testing.T) {
	t.Setenv("MUTATION_TESTING", "true")
	ooze.Release(
		t,
		ooze.WithRepositoryRoot("../"),
		ooze.WithMinimumThreshold(0.9),
		ooze.Parallel(),
		ooze.IgnoreSourceFiles("^((test|server|pkg|app|config|tools|internal/heartbeat)/.+|main|.+structs|.+interfaces|.+errors)\\.go$"),
		ooze.ForceColors(),
	)
}
