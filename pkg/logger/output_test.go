package logger

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"regexp"
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
		Set(Config{Level: "INFO"})
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
		Set(Config{Level: "INFO"})
		Infow("Via package function")
		Logger().Infow("Via Logger")
	})

	for line := range strings.SplitSeq(strings.TrimSpace(out), "\n") {
		require.Contains(t, line, "output_test.go", "caller must be the call site, not logger.go: %s", line)
	}
}

var isoTimestamp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$`)

func TestJSONFormat(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		With("voting_round", 12345).Infow("Round submitted", "contract", "submit1")
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Len(t, lines, 1, "one object per line, got %q", out)

	var line map[string]any
	require.NoError(t, json.Unmarshal([]byte(lines[0]), &line))

	require.Equal(t, "info", line["level"], "lowercase level name")
	require.Equal(t, "Round submitted", line["msg"])
	require.EqualValues(t, 12345, line["voting_round"], "With field is a top-level key")
	require.Equal(t, "submit1", line["contract"])
	require.Regexp(t, isoTimestamp, line["ts"], "ISO 8601 UTC timestamp")
	require.Contains(t, line["caller"], "output_test.go")
}

func TestConsoleFormatIsPlainWhenNotATerminal(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Infow("Round submitted", "voting_round", 12345)
	})

	require.NotContains(t, out, "\033[", "no colour escapes on a pipe")
	require.Contains(t, out, "INFO")
	require.Regexp(t, `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z\t`, out, "ISO 8601 UTC timestamp")
}

func TestUnknownFormatFallsBackToConsole(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: "xml"})
	})

	require.Contains(t, out, "Invalid logger format")
	require.Contains(t, out, `"format": "xml"`)
}

func TestDeprecatedFileOptionWarnsAndStillWritesStdout(t *testing.T) {
	file := filepath.Join(t.TempDir(), "service.log")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", File: file})
		Infow("Round submitted")
		SyncFileLogger()
	})

	require.Contains(t, out, "Deprecated logger file output is set")
	require.Contains(t, out, "Round submitted", "standard output is always written")
	require.NotContains(t, out, "Syncing file logger", "sync is silent")

	content, err := os.ReadFile(file)
	require.NoError(t, err)
	require.Contains(t, string(content), "Round submitted", "the file is still written this release")
}
