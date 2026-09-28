package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// captureStdout runs fn with standard output redirected to a pipe and returns
// what was written. The handler binds os.Stdout when Set builds it, so fn
// must call Set itself.
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

// nextLine returns the source reference of the line after the call.
func nextLine() string {
	_, _, line, _ := runtime.Caller(1)

	return fmt.Sprintf("logger/logger_test.go:%d", line+1)
}

var isoTimestamp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$`)

func TestJSONFormat(t *testing.T) {
	var src string
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		src = nextLine()
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
	require.Regexp(t, isoTimestamp, line["time"], "ISO 8601 UTC timestamp")
	require.Equal(t, src, line["source"], "source is the call site")
}

func TestConsoleFormat(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Infow("Round submitted", "voting_round", 12345)
		Logger().Errorw("Submit failed", "error", errors.New("nonce too low"))
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Len(t, lines, 2)
	require.Regexp(t, "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{3}Z\tINFO\t\\S+\tRound submitted\tvoting_round=12345$", lines[0])
	require.Regexp(t, "\tERROR\t\\S+\tSubmit failed\terror=\"nonce too low\"$", lines[1])
	require.NotContains(t, out, "\033[", "no escape codes")
}

func TestPrintfAndPrintFunctions(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	var src string
	out := captureStdout(t, func() {
		Set(Config{Level: "DEBUG"})
		src = nextLine()
		Debugf("payload %d bytes", 412)
		Warn("tip ", 12, " blocks behind")
	})

	require.Contains(t, out, src+"\tpayload 412 bytes", "the line reports its caller")
	require.Contains(t, out, "\tDEBUG\t")
	require.Contains(t, out, "\tpayload 412 bytes\n", "no fields, nothing trailing")
	require.Contains(t, out, "\tWARN\t")
	require.Contains(t, out, "\ttip 12 blocks behind\n")
}

func TestLevelFilter(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "WARN"})
		Infow("Hidden")
		Warnw("Shown")
	})

	require.NotContains(t, out, "Hidden")
	require.Contains(t, out, "Shown")
}

func TestInvalidConfigKeepsThePreviousLogger(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Set(Config{Level: "LOUD"})
		Set(Config{Level: "INFO", Format: "xml"})
		Infow("Still logging")
	})

	require.Contains(t, out, "Invalid logger configuration, keeping the previous one\terror=\"unknown logger level \\\"LOUD\\\"")
	require.Contains(t, out, `error="unknown logger format \"xml\"`)
	require.Contains(t, out, "Still logging")
}

// Legacy file settings are ignored and reported at startup.
func TestFileKeysAreIgnoredAndNamed(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	var src string
	file := filepath.Join(t.TempDir(), "indexer.log")
	out := captureStdout(t, func() {
		src = nextLine()
		Set(Config{Level: "INFO", File: file, MaxFileSize: 100, MaxBackups: 10, MaxAgeDays: 30, Console: true})
		Infow("Round submitted")
	})

	require.Contains(t, out, src+`	Ignoring logger keys that no longer do anything	keys="file max_file_size max_backups max_age_days console"`,
		"the line is about the caller's configuration, so it reports the caller")
	require.Contains(t, out, "Round submitted")
	require.NoFileExists(t, file)
}

// An omitted level matches DefaultConfig.
func TestAbsentLevelIsDebug(t *testing.T) {
	level, err := parseLevel("")
	require.NoError(t, err)
	require.Equal(t, slog.LevelDebug, level)

	defaulted, err := parseLevel(DefaultConfig().Level)
	require.NoError(t, err)
	require.Equal(t, level, defaulted)
}

// Legacy PANIC and DPANIC settings are accepted as FATAL.
func TestZapLevelsAboveErrorAreFatal(t *testing.T) {
	for _, name := range []string{"FATAL", "PANIC", "DPANIC", "panic"} {
		level, err := parseLevel(name)
		require.NoError(t, err, name)
		require.Equal(t, LevelFatal, level, name)
	}
}

// lazyValue records whether anything rendered it.
type lazyValue struct{ rendered *bool }

func (v lazyValue) String() string {
	*v.rendered = true

	return "expensive"
}

// Disabled log calls must not format their arguments.
func TestDisabledLevelDoesNotFormat(t *testing.T) {
	var rendered bool

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Debugf("%s", lazyValue{&rendered})
		Debug(lazyValue{&rendered})
		Logger().Debugf("%s", lazyValue{&rendered})
		Logger().Debug(lazyValue{&rendered})
	})

	require.False(t, rendered, "String() ran for a line that was never written")
	require.NotContains(t, out, "expensive")
}

// Newlines in messages and fields must not split a console entry.
func TestConsoleKeepsOneLinePerRecord(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Errorw("Query failed", "error", "first\nsecond")
		Infof("tail: %s", "a\r\nb")
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Len(t, lines, 2, "one physical line per record, got %q", out)
	require.Contains(t, lines[0], `error="first\nsecond"`)
	require.Contains(t, lines[1], `tail: a\r\nb`)
}

// Caller fields must not overwrite metadata, even when their value types match.
func TestJSONKeepsItsOwnKeysWhenFieldsCollide(t *testing.T) {
	for name, field := range map[string][]any{
		"other kinds": {"time", "upstream-time", "level", "high", "source", "cache", "msg", 7},
		"same kinds":  {"time", time.Unix(0, 0), "level", "high", "source", "cache", "msg", "upstream-message"},
	} {
		t.Run(name, func(t *testing.T) {
			out := captureStdout(t, func() {
				Set(Config{Level: "INFO", Format: FormatJSON})
				Infow("Response", field...)
			})

			var line map[string]any
			require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))

			require.Regexp(t, isoTimestamp, line["time"], "the record keeps its own timestamp")
			require.Equal(t, "info", line["level"])
			require.Equal(t, "Response", line["msg"], "the searchable message is the record's own")
			require.Contains(t, line, "logged_time", "the caller's field is kept, renamed")
			require.Equal(t, "high", line["logged_level"])
			require.Equal(t, "cache", line["logged_source"])
			require.Contains(t, line, "logged_msg")
		})
	}
}

// JSON timestamps use the same format in metadata, fields and groups.
func TestJSONTimestampsAreMilliseconds(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		Infow("Batch indexed", "observed_at", time.Unix(0, 123456789), slog.Group("rpc", "asked_at", time.Unix(0, 987654321)))
	})

	var line map[string]any
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))

	require.Regexp(t, isoTimestamp, line["time"])
	require.Equal(t, "1970-01-01T00:00:00.123Z", line["observed_at"])

	rpc, ok := line["rpc"].(map[string]any)
	require.True(t, ok, "the group is an object")
	require.Equal(t, "1970-01-01T00:00:00.987Z", rpc["asked_at"])
}

// JSON durations use milliseconds, including fractional values.
func TestJSONDurationsAreMilliseconds(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		Infow("Tip read", "took", 1500*time.Millisecond, "parse", 240*time.Microsecond)
	})

	var line map[string]any
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))

	require.InDelta(t, 1500.0, line["took"], 0.0001)
	require.InDelta(t, 0.24, line["parse"], 0.0001, "a sub-millisecond duration keeps its detail")
}

// Fields attached once through With are renamed once, not on every line.
func TestWithRenamesReservedKeys(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		With("time", time.Unix(0, 0), "chain_id", 14).Infow("Batch indexed")
	})

	var line map[string]any
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))

	require.Regexp(t, isoTimestamp, line["time"])
	require.Contains(t, line, "logged_time")
	require.EqualValues(t, 14, line["chain_id"])
}

// brokenError panics when its Error method is called.
type brokenError struct{}

func (brokenError) Error() string { panic("Error called") }

// Console logging must handle typed nil errors and panicking Error methods.
func TestConsoleSurvivesBrokenErrorValues(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	var missing *os.PathError

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Errorw("Read failed", "error", missing)
		Errorw("Submit failed", "error", brokenError{})
		Errorw("Query failed", "error", errors.New("nonce too low"))
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Len(t, lines, 3)
	require.Contains(t, lines[0], "error=<nil>")
	require.Contains(t, lines[1], "PANIC", "the panic is reported in the field, not raised")
	require.Contains(t, lines[2], `error="nonce too low"`)
}

// Unnamed groups share the top-level namespace, so reserved keys must be renamed.
func TestUnnamedGroupFieldsAreRenamed(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		With(slog.Group("", "time", "1970")).Infow("Round submitted",
			slog.Group("", "msg", "upstream-message", "chain_id", 14))
	})

	var line map[string]any
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))

	require.Regexp(t, isoTimestamp, line["time"])
	require.Equal(t, "Round submitted", line["msg"])
	require.Equal(t, "upstream-message", line["logged_msg"])
	require.Equal(t, "1970", line["logged_time"])
	require.EqualValues(t, 14, line["chain_id"], "an ordinary field in the group is untouched")
}

// Unnamed groups must not add a prefix to console fields.
func TestConsoleInlinesUnnamedGroups(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Infow("Round submitted", slog.Group("", "chain_id", 14), slog.Group("rpc", "target", "node:9650"))
	})

	require.Contains(t, out, "\tchain_id=14 rpc.target=node:9650\n")
}

// groupValuer returns reserved field names and counts LogValue calls.
type groupValuer struct{ resolved *int }

func (v groupValuer) LogValue() slog.Value {
	*v.resolved++

	return slog.GroupValue(slog.String("msg", "upstream-message"), slog.String("level", "high"))
}

// Resolve LogValuer fields before checking reserved names, and only once.
func TestLogValuerCannotTakeMetadataKeys(t *testing.T) {
	for name, log := range map[string]func(v groupValuer){
		"direct": func(v groupValuer) { Infow("Round submitted", slog.Any("", v)) },
		"with":   func(v groupValuer) { With(slog.Any("", v)).Infow("Round submitted") },
	} {
		t.Run(name, func(t *testing.T) {
			resolved := 0

			out := captureStdout(t, func() {
				Set(Config{Level: "INFO", Format: FormatJSON})
				log(groupValuer{&resolved})
			})

			var line map[string]any
			require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))

			require.Equal(t, "Round submitted", line["msg"])
			require.Equal(t, "info", line["level"])
			require.Equal(t, "upstream-message", line["logged_msg"])
			require.Equal(t, "high", line["logged_level"])
			require.Equal(t, 1, resolved, "the handler gets the resolved value, so LogValue runs once")
		})
	}
}

// Control characters in field names must be escaped too.
func TestConsoleKeysStayOnOneLine(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Infow("Query failed", "first\nsecond", 1)
	})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Len(t, lines, 1, "one physical line per record, got %q", out)
	require.Contains(t, lines[0], `first\nsecond=1`)
}

// Panic calls must panic even when the configured level filters other messages.
func TestPanicAlwaysPanics(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "FATAL"})
		require.PanicsWithValue(t, "invariant", func() { Panic("invariant") })
		require.PanicsWithValue(t, "invariant 1", func() { Panicf("invariant %d", 1) })
		require.PanicsWithValue(t, "invariant", func() { Logger().Panicw("invariant", "worker", 3) })
	})

	require.Contains(t, out, "\tFATAL\t")
	require.Contains(t, out, "worker=3")
}

func TestNopPanicsLikeTheRealLogger(t *testing.T) {
	require.PanicsWithValue(t, "boom", func() { Nop{}.Panic("boom") })
	require.PanicsWithValue(t, "boom 42", func() { Nop{}.Panicf("boom %d", 42) })
	require.PanicsWithValue(t, "boom", func() { Nop{}.Panicw("boom", "k", 1) })
}

// fatalCalls lists the Fatal variants that must exit, including those on Nop.
var fatalCalls = map[string]func(){
	"nop-fatal":  func() { Nop{}.Fatal("boom") },
	"nop-fatalf": func() { Nop{}.Fatalf("boom %d", 42) },
	"nop-fatalw": func() { Nop{}.Fatalw("boom", "k", 1) },
	"fatal":      func() { Fatal("boom") },
	"fatalf":     func() { Fatalf("boom %d", 42) },
	"fatalw":     func() { Fatalw("boom", "k", 1) },
}

const fatalCallEnv = "LOGGER_TEST_FATAL_CALL"

// TestFatalExits runs each Fatal call in a subprocess and checks its exit status.
// The environment variable selects the call when the test binary restarts.
func TestFatalExits(t *testing.T) {
	if name := os.Getenv(fatalCallEnv); name != "" {
		fatalCalls[name]()
		// The call returned. Exit 0 rather than fail, since a failed test
		// also exits 1 and the parent could not tell the two apart.
		os.Exit(0)
	}

	for name := range fatalCalls {
		t.Run(name, func(t *testing.T) {
			//nolint:gosec // the command is this test binary, not user input
			cmd := exec.Command(os.Args[0], "-test.run=TestFatalExits")
			cmd.Env = append(os.Environ(), fatalCallEnv+"="+name)

			var exit *exec.ExitError
			require.ErrorAs(t, cmd.Run(), &exit)
			require.Equal(t, 1, exit.ExitCode())
		})
	}
}

// Console fields are key=value, with a group flattened into dotted keys and
// a value quoted only when bare text would be ambiguous.
func TestConsoleFieldsAreKeyValue(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		With("chain_id", 14).Infow("Tip read", slog.Group("rpc", "target", "node:9650", "took", 1500*time.Millisecond), "at", time.Unix(0, 0))
	})

	require.Contains(t, out, "\tTip read\tchain_id=14 rpc.target=node:9650 rpc.took=1.5s at=1970-01-01T00:00:00.000Z\n")
}

func TestConsoleQuotesValuesThatNeedIt(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Infow("Startup check",
			"check", "rpc node",
			"empty", "",
			"equals", "a=b",
			"plain", "ok",
		)
	})

	require.Contains(t, out, `check="rpc node" empty="" equals="a=b" plain=ok`)
}

// The custom FATAL level must appear as "fatal" in JSON.
func TestFatalLevelName(t *testing.T) {
	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		Logger().log(LevelFatal, "Indexer stopped", "reason", "behind the chain")
	})

	var line map[string]any
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &line))
	require.Equal(t, "fatal", line["level"])
	require.Equal(t, "Indexer stopped", line["msg"])
}

// A FATAL threshold must allow FATAL entries.
func TestFatalIsNotFilteredOutAtItsOwnLevel(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "FATAL"})
		Errorw("A unit of work failed")
		Logger().log(LevelFatal, "Indexer stopped")
	})

	require.NotContains(t, out, "A unit of work failed")
	require.Contains(t, out, "Indexer stopped")
}

// Console output uses color by default, including when redirected.
func TestConsoleIsColouredByDefault(t *testing.T) {
	t.Setenv("NO_COLOR", "")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Errorw("History drop failed", "dependency", "database")
	})

	require.Contains(t, out, "\x1b[")
}

func TestNoColorTurnsItOff(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Errorw("History drop failed", "dependency", "database")
	})

	require.NotContains(t, out, "\x1b[")
}

// The custom FATAL level uses its configured console color.
func TestFatalIsColoured(t *testing.T) {
	t.Setenv("NO_COLOR", "")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO"})
		Logger().log(LevelFatal, "Indexer stopped")
	})

	require.Contains(t, out, levelColor[LevelFatal]+"FATAL"+ansiReset)
}

// JSON output must not contain ANSI color codes.
func TestJSONIsNeverColoured(t *testing.T) {
	t.Setenv("NO_COLOR", "")

	out := captureStdout(t, func() {
		Set(Config{Level: "INFO", Format: FormatJSON})
		Errorw("History drop failed", "dependency", "database")
		Logger().log(LevelFatal, "Indexer stopped")
	})

	require.NotContains(t, out, "\x1b[")
	require.Contains(t, out, `"level":"fatal"`)
}

// An empty NO_COLOR value leaves color enabled.
func TestEmptyNoColorLeavesColourOn(t *testing.T) {
	t.Setenv("NO_COLOR", "")

	require.True(t, useColor())
}
