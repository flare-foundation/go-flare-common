// Package logger provides console and JSON logging for Flare's Go services,
// using log/slog.
//
// Logs are written to standard output only. The container runtime should
// handle retention and rotation. The default is console output at DEBUG level,
// use Set to change the format or minimum level.
//
// Always use structured logging, including for DEBUG messages. Use the
// functions ending in w, keep the message fixed, and put variable values
// in fields so entries can be searched and filtered consistently:
//
//	logger.Infow("Round submitted", "voting_round", 12345, "protocol_id", 100)
//
// Structured functions accept key-value pairs or slog.Attr values after the
// message. Top-level fields named time, level, msg or source are prefixed with
// logged_ to keep them separate from log metadata.
//
// Use With to include the same fields in several entries:
//
//	log := logger.With("voting_round", 12345)
//	log.Infow("Round submitted")
//
// Printf functions (ending in f) and print functions (without a suffix) are
// retained for compatibility. Do not use them in new or updated logging code.
//
// Fatalw logs and exits with status 1. Panicw logs and panics, allowing
// deferred functions to run.
package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// Supported output formats.
const (
	// FormatConsole writes time, level, source and message separated by tabs,
	// followed by key=value fields. Set NO_COLOR to a nonempty value to disable color.
	FormatConsole = "console"
	// FormatJSON writes one JSON object per entry, with lowercase levels,
	// UTC timestamps at millisecond precision and durations in milliseconds.
	FormatJSON = "json"
)

// LevelFatal is the log level used by both Fatal and Panic calls.
const LevelFatal = slog.LevelError + 4

// timeLayout is ISO 8601 in UTC with millisecond precision.
const timeLayout = "2006-01-02T15:04:05.000Z07:00"

// loggedPrefix is added to field names that conflict with log metadata.
const loggedPrefix = "logged_"

// Config sets the minimum log level and output format.
// Deprecated fields allow older TOML configurations to load but have no effect.
type Config struct {
	Level  string `toml:"level"`  // DEBUG (default), INFO, WARN, ERROR or FATAL
	Format string `toml:"format"` // FormatConsole (default) or FormatJSON

	File        string `toml:"file"`          // Deprecated: ignored; logs go to standard output.
	MaxFileSize int    `toml:"max_file_size"` // Deprecated: ignored; file rotation is no longer supported.
	MaxBackups  int    `toml:"max_backups"`   // Deprecated: ignored; file rotation is no longer supported.
	MaxAgeDays  int    `toml:"max_age_days"`  // Deprecated: ignored; file rotation is no longer supported.
	Console     bool   `toml:"console"`       // Deprecated: ignored; use Format to select the output format.
}

// DefaultConfig returns a configuration with DEBUG level and console output.
func DefaultConfig() Config {
	return Config{Level: "DEBUG", Format: FormatConsole}
}

// Log is a logger with a configured output and optional shared fields.
// It is safe to use from multiple goroutines.
type Log struct {
	s *slog.Logger
}

var global atomic.Pointer[Log]

func init() {
	l, _ := newLog(DefaultConfig())
	global.Store(l)
}

// newLog returns a logger for cfg. An unknown level or format is an error.
func newLog(cfg Config) (*Log, error) {
	level, err := parseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	var h slog.Handler
	switch cfg.Format {
	case FormatJSON:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level, AddSource: true, ReplaceAttr: replaceAttr})
	case FormatConsole, "":
		h = newConsoleHandler(os.Stdout, level, useColor())
	default:
		return nil, fmt.Errorf("unknown logger format %q, want %q or %q", cfg.Format, FormatConsole, FormatJSON)
	}

	return &Log{s: slog.New(h)}, nil
}

// Set replaces the package logger with one configured by cfg. If cfg is
// invalid, it keeps the current logger and logs an ERROR message if that level
// is enabled. Existing Log values keep their configuration.
//
// Set is safe to call while other goroutines are logging.
func Set(cfg Config) {
	// Attribute configuration messages to the caller of Set.
	const skip = 3 // runtime.Callers, record, Set

	l, err := newLog(cfg)
	if err != nil {
		Logger().record(skip, slog.LevelError, "Invalid logger configuration, keeping the previous one", "error", err)
		return
	}

	global.Store(l)

	if keys := retiredKeys(cfg); len(keys) > 0 {
		l.record(skip, slog.LevelInfo, "Ignoring logger keys that no longer do anything", "keys", strings.Join(keys, " "))
	}
}

// retiredKeys lists deprecated settings with nonzero values.
func retiredKeys(cfg Config) []string {
	var keys []string
	if cfg.File != "" {
		keys = append(keys, "file")
	}
	if cfg.MaxFileSize != 0 {
		keys = append(keys, "max_file_size")
	}
	if cfg.MaxBackups != 0 {
		keys = append(keys, "max_backups")
	}
	if cfg.MaxAgeDays != 0 {
		keys = append(keys, "max_age_days")
	}
	if cfg.Console {
		keys = append(keys, "console")
	}

	return keys
}

// Logger returns the current package logger. A later call to Set does not
// change the returned logger.
func Logger() *Log {
	return global.Load()
}

// Deprecated: this function does nothing. Logs go directly to standard output
// and do not need flushing.
func SyncFileLogger() {}

// With returns a copy of the current package logger with fields added to every
// entry. Arguments are key-value pairs or slog.Attr values. Top-level fields
// named time, level, msg or source get a logged_ prefix, including fields
// from unnamed groups and slog.LogValuer values.
func With(keysAndValues ...any) *Log {
	return Logger().With(keysAndValues...)
}

// With returns a copy of l with additional fields, leaving l unchanged.
// Arguments and field names follow the same rules as [With].
func (l *Log) With(keysAndValues ...any) *Log {
	// Parse and rename shared fields once, when the logger is created.
	var parsed slog.Record
	parsed.Add(keysAndValues...)
	parsed = renameReserved(parsed)

	fields := make([]any, 0, parsed.NumAttrs())
	parsed.Attrs(func(a slog.Attr) bool {
		fields = append(fields, a)

		return true
	})

	return &Log{s: l.s.With(fields...)}
}

func parseLevel(s string) (slog.Level, error) {
	// Match DefaultConfig when the level is omitted.
	if s == "" {
		return slog.LevelDebug, nil
	}

	switch strings.ToUpper(s) {
	case "FATAL", "PANIC", "DPANIC": // Accept legacy zap levels as FATAL.
		return LevelFatal, nil
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(s)); err != nil {
		return 0, fmt.Errorf("unknown logger level %q, want DEBUG, INFO, WARN, ERROR or FATAL", s)
	}

	return level, nil
}

// Console output uses color unless NO_COLOR is nonempty, even when redirected.
func useColor() bool {
	return os.Getenv("NO_COLOR") == ""
}

// replaceAttr formats JSON levels and source locations, and converts time and
// duration fields to the shared log format. Reserved names are handled earlier.
func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if len(groups) == 0 {
		switch a.Key {
		case slog.LevelKey:
			if level, ok := a.Value.Any().(slog.Level); ok {
				return slog.String(a.Key, strings.ToLower(levelName(level)))
			}
		case slog.SourceKey:
			if src, ok := a.Value.Any().(*slog.Source); ok {
				return slog.String(a.Key, sourceRef(src.File, src.Line))
			}
		}
	}

	// Use UTC with millisecond precision for both metadata and caller fields.
	if a.Value.Kind() == slog.KindTime {
		return slog.String(a.Key, a.Value.Time().UTC().Format(timeLayout))
	}

	// Convert slog's nanoseconds to milliseconds to match the NestJS services.
	if a.Value.Kind() == slog.KindDuration {
		return slog.Float64(a.Key, float64(a.Value.Duration())/float64(time.Millisecond))
	}

	return a
}

// reservedKeys lists names used by log metadata. Rename conflicting caller
// fields before the handler processes them.
var reservedKeys = map[string]bool{
	slog.TimeKey:    true,
	slog.LevelKey:   true,
	slog.MessageKey: true,
	slog.SourceKey:  true,
}

// renameReserved resolves values and renames conflicting fields before they
// reach the handler. It returns the original record when no changes are needed.
func renameReserved(r slog.Record) slog.Record {
	rewrite := false
	r.Attrs(func(a slog.Attr) bool {
		rewrite = needsRewrite(a)

		return !rewrite
	})

	if !rewrite {
		return r
	}

	renamed := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	r.Attrs(func(a slog.Attr) bool {
		renamed.AddAttrs(rewriteAttr(a))

		return true
	})

	return renamed
}

// needsRewrite checks for reserved names and unresolved LogValuer values.
// A LogValuer may return an unnamed group containing reserved names.
func needsRewrite(a slog.Attr) bool {
	if a.Value.Kind() == slog.KindLogValuer {
		return true
	}

	// Unnamed groups are inlined, so their fields share the parent's namespace.
	if a.Key == "" && a.Value.Kind() == slog.KindGroup {
		return slices.ContainsFunc(a.Value.Group(), needsRewrite)
	}

	return reservedKeys[a.Key]
}

// rewriteAttr resolves a value and renames reserved keys, including those in
// unnamed groups. Passing the resolved value prevents the handler from calling
// LogValue again.
func rewriteAttr(a slog.Attr) slog.Attr {
	a.Value = a.Value.Resolve()

	if a.Key == "" && a.Value.Kind() == slog.KindGroup {
		group := a.Value.Group()
		inner := make([]slog.Attr, len(group))
		for i, field := range group {
			inner[i] = rewriteAttr(field)
		}

		return slog.Attr{Key: a.Key, Value: slog.GroupValue(inner...)}
	}

	if reservedKeys[a.Key] {
		a.Key = loggedPrefix + a.Key
	}

	return a
}

// sourceRef formats a source location as package/file.go:line.
func sourceRef(file string, line int) string {
	return filepath.Base(filepath.Dir(file)) + "/" + filepath.Base(file) + ":" + strconv.Itoa(line)
}

// levelName adds FATAL to slog's standard level names.
func levelName(level slog.Level) string {
	if level >= LevelFatal {
		return "FATAL"
	}

	return level.String()
}

// log records the file and line number of the user's logging call.
func (l *Log) log(level slog.Level, msg string, keysAndValues ...any) {
	// Skip runtime.Callers, record, log, and the public logging function.
	l.record(4, level, msg, keysAndValues...)
}

// record writes an entry if its level is enabled. skip selects the caller
// frame using runtime.Callers.
func (l *Log) record(skip int, level slog.Level, msg string, keysAndValues ...any) {
	ctx := context.Background()
	if !l.s.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(skip, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(keysAndValues...)
	_ = l.s.Handler().Handle(ctx, renameReserved(r))
}

// logf and logPrint skip argument formatting when the level is disabled.
func (l *Log) logf(level slog.Level, format string, args ...any) {
	if !l.s.Enabled(context.Background(), level) {
		return
	}

	// Skip runtime.Callers, record, logf, and the public logging function.
	l.record(4, level, fmt.Sprintf(format, args...))
}

func (l *Log) logPrint(level slog.Level, args ...any) {
	if !l.s.Enabled(context.Background(), level) {
		return
	}

	l.record(4, level, fmt.Sprint(args...))
}
