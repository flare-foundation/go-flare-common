package logger

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// captureStdout runs fn with standard output redirected to a pipe and returns
// what was written. The logger reads os.Stdout when Set builds its core, so
// fn must call Set itself.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)

	orig := os.Stdout
	os.Stdout = w
	t.Cleanup(func() {
		os.Stdout = orig
		Set(DefaultConfig())
	})

	fn()

	require.NoError(t, w.Close())
	os.Stdout = orig

	out, err := io.ReadAll(r)
	require.NoError(t, err)

	return string(out)
}

func TestStructuredFunctionsCarryFields(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Console: true})
		Infow("Round submitted", "voting_round", 12345)
		With("protocol_id", 100).Infow("Round finalised")
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Len(t, lines, 2)
	require.Contains(t, lines[0], "Round submitted")
	require.Contains(t, lines[0], `"voting_round": 12345`)
	require.Contains(t, lines[1], "Round finalised")
	require.Contains(t, lines[1], `"protocol_id": 100`, "With field is on the line")
}

func TestCallerIsTheCallSite(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Console: true})
		Infow("Via package function")
		Logger().Infow("Via Logger")
	})

	for line := range strings.SplitSeq(strings.TrimSpace(out), "\n") {
		require.Contains(t, line, "output_test.go", "caller must be the call site, not logger.go: %s", line)
	}
}
