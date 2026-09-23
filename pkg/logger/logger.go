// Package logger is the shared logger of Flare's Go services, built on zap.
//
// Messages an operator searches for are fixed, and their identifiers travel
// as fields:
//
//	logger.Infow("round submitted", "voting_round", 12345, "protocol_id", 100)
//
// Fields shared by every line of a unit of work are attached once:
//
//	log := logger.With("voting_round", 12345)
//	log.Infow("round submitted")
package logger

import (
	"io"
	"os"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Formats of the log output.
const (
	// FormatConsole is one readable line per entry, coloured when standard
	// output is a terminal.
	FormatConsole = "console"
	// FormatJSON is one JSON object per entry, for log shippers.
	FormatJSON = "json"
)

// timeLayout is ISO 8601 in UTC with millisecond precision.
const timeLayout = "2006-01-02T15:04:05.000Z07:00"

// state is one configured logger: base carries no caller skip and is what
// Logger and With hand out, pkg skips this package's wrapper functions.
type state struct {
	base *zap.SugaredLogger
	pkg  *zap.SugaredLogger
}

var current atomic.Pointer[state]

func init() {
	current.Store(createState(DefaultConfig()))
}

func pkg() *zap.SugaredLogger {
	return current.Load().pkg
}

// Config holds the logger configuration.
type Config struct {
	Level  string `toml:"level"`  // valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (zap)
	Format string `toml:"format"` // FormatConsole (default) or FormatJSON

	// File writes a copy of the log to a rotated file.
	//
	// Deprecated: services log to standard output only and leave retention
	// to the runtime. Setting File logs a warning. MaxFileSize, MaxBackups
	// and MaxAgeDays size the rotation and are deprecated with it.
	File        string `toml:"file"`
	MaxFileSize int    `toml:"max_file_size"` // megabytes; 0 → lumberjack default (100 MB)
	MaxBackups  int    `toml:"max_backups"`   // rotated files retained on disk; 0 → defaultMaxBackups
	MaxAgeDays  int    `toml:"max_age_days"`  // max age of rotated files; 0 → defaultMaxAgeDays

	// Console used to switch standard output on and off.
	//
	// Deprecated: standard output is always written. The field is ignored.
	Console bool `toml:"console"`
}

const (
	// Defensive caps applied when Config.MaxBackups / MaxAgeDays are unset; without
	// them lumberjack retains every rotated file forever, leaking disk on long-running services.
	defaultMaxBackups = 10
	defaultMaxAgeDays = 30
)

// DefaultConfig returns the default logger configuration.
//
//	Level: "DEBUG"
//	Format: FormatConsole
func DefaultConfig() Config {
	return Config{
		Level:  "DEBUG",
		Format: FormatConsole,
	}
}

// Logger returns the configured logger for direct use. Lines it writes
// report their caller correctly.
func Logger() *zap.SugaredLogger {
	return current.Load().base
}

// With returns a logger that adds the given fields to every line. Attach the
// fields of a unit of work once, here, rather than at every call.
func With(keysAndValues ...any) *zap.SugaredLogger {
	return Logger().With(keysAndValues...)
}

// Set configures logger according to Config. Safe to call concurrently
// with logging calls.
func Set(cfg Config) {
	current.Store(createState(cfg))

	if cfg.File != "" {
		Warnw(
			"Deprecated logger file output is set: services log to standard output only, remove the file, max_file_size, max_backups and max_age_days options",
			"file", cfg.File,
		)
	}
}

func createState(config Config) *state {
	// Resolve level first so AtomicLevel is correctly populated from the start;
	// otherwise the cores enable Info-and-above until SetLevel runs.
	level, err := zapcore.ParseLevel(config.Level)
	parseErr := err
	if err != nil {
		// Fall back to DEBUG (the DefaultConfig level) rather than the
		// zero value of zapcore.Level (which is InfoLevel). Silently
		// downgrading to INFO would drop messages the operator likely
		// intended to see; DEBUG keeps everything visible — including
		// the subsequent parse-error log.
		level = zapcore.DebugLevel
	}
	atom := zap.NewAtomicLevelAt(level)

	format := config.Format
	if format == "" {
		format = FormatConsole
	}
	formatKnown := format == FormatConsole || format == FormatJSON
	if !formatKnown {
		format = FormatConsole
	}

	cores := []zapcore.Core{createStdoutCore(format, atom)}
	if len(config.File) > 0 {
		cores = append(cores, createFileLoggerCore(config, atom))
	}

	// Stack traces are kept for panics only: an ERROR or FATAL line is one
	// event, its fields say what failed, and a shipper reads it as one line.
	core := zapcore.NewTee(cores...)
	base := zap.New(core,
		zap.AddStacktrace(zap.PanicLevel),
		zap.AddCaller(),
	).Sugar()

	s := &state{
		base: base,
		pkg:  base.WithOptions(zap.AddCallerSkip(1)),
	}

	if parseErr != nil {
		s.base.Errorw("Invalid logger level, falling back to DEBUG", "level", config.Level)
	}
	if !formatKnown {
		s.base.Errorw("Invalid logger format, falling back to console", "format", config.Format)
	}
	return s
}

// SyncFileLogger flushes buffered log entries. Standard output is
// unbuffered, so this matters only when the deprecated file output is set.
// It is called before the process exits on a FATAL or PANIC line.
func SyncFileLogger() {
	_ = Logger().Sync()
}

func createFileLoggerCore(config Config, atom zap.AtomicLevel) zapcore.Core {
	maxBackups := config.MaxBackups
	if maxBackups == 0 {
		maxBackups = defaultMaxBackups
	}
	maxAge := config.MaxAgeDays
	if maxAge == 0 {
		maxAge = defaultMaxAgeDays
	}
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.File,
		MaxSize:    config.MaxFileSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	})
	encoderCfg := encoderConfig()
	encoderCfg.EncodeLevel = plainLevelEncoder
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		w,
		atom,
	)
}

func encoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	// Timestamps are written in UTC whatever the host's zone.
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(timeLayout))
	}

	return cfg
}

// isTerminal reports whether f is a character device, which is what a
// terminal is and a pipe or a file is not.
func isTerminal(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}

type noSyncWriter struct {
	io.Writer
}

func (n noSyncWriter) Sync() error {
	return nil
}

// createStdoutCore writes to standard output in the given format. The
// console format is coloured only when standard output is a terminal, so
// container logs stay plain text.
func createStdoutCore(format string, atom zap.AtomicLevel) zapcore.Core {
	cfg := encoderConfig()
	out := noSyncWriter{os.Stdout}

	if format == FormatJSON {
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder

		return zapcore.NewCore(zapcore.NewJSONEncoder(cfg), out, atom)
	}

	cfg.EncodeLevel = plainLevelEncoder
	if isTerminal(os.Stdout) {
		cfg.EncodeLevel = consoleColorLevelEncoder
	}

	return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), out, atom)
}

func consoleColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	s, ok := levelToCapitalColorString[l]
	if !ok {
		s = unknownLevelColor.Wrap(l.CapitalString())
	}
	enc.AppendString(s)
}

func plainLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(l.CapitalString())
}

// Debugw logs a message with key-value fields at DEBUG level.
func Debugw(msg string, keysAndValues ...any) {
	pkg().Debugw(msg, keysAndValues...)
}

// Infow logs a message with key-value fields at INFO level.
func Infow(msg string, keysAndValues ...any) {
	pkg().Infow(msg, keysAndValues...)
}

// Warnw logs a message with key-value fields at WARN level.
func Warnw(msg string, keysAndValues ...any) {
	pkg().Warnw(msg, keysAndValues...)
}

// Errorw logs a message with key-value fields at ERROR level.
func Errorw(msg string, keysAndValues ...any) {
	pkg().Errorw(msg, keysAndValues...)
}

// Panicw logs a message with key-value fields at PANIC level and panics.
//
// Defers will be executed.
func Panicw(msg string, keysAndValues ...any) {
	SyncFileLogger()
	pkg().Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with key-value fields at FATAL level and calls os.Exit.
//
// Defers will not be executed.
func Fatalw(msg string, keysAndValues ...any) {
	SyncFileLogger()
	pkg().Fatalw(msg, keysAndValues...)
}

// Debugf formats the message and logs it at DEBUG level.
func Debugf(msg string, args ...any) {
	pkg().Debugf(msg, args...)
}

// Infof formats the message and logs it at INFO level.
func Infof(msg string, args ...any) {
	pkg().Infof(msg, args...)
}

// Warnf formats the message and logs it at WARN level.
func Warnf(msg string, args ...any) {
	pkg().Warnf(msg, args...)
}

// Errorf formats the message and logs it at ERROR level.
func Errorf(msg string, args ...any) {
	pkg().Errorf(msg, args...)
}

// Panicf formats the message and logs it at PANIC level and panics.
//
// Defers will be executed.
func Panicf(msg string, args ...any) {
	SyncFileLogger()
	pkg().Panicf(msg, args...)
}

// Fatalf formats the message and logs it at FATAL level and calls os.Exit.
//
// Defers will not be executed.
func Fatalf(msg string, args ...any) {
	SyncFileLogger()
	pkg().Fatalf(msg, args...)
}

// Debug logs arguments at DEBUG level.
func Debug(args ...any) {
	pkg().Debug(args...)
}

// Info logs arguments at INFO level.
func Info(args ...any) {
	pkg().Info(args...)
}

// Warn logs arguments at WARN level.
func Warn(args ...any) {
	pkg().Warn(args...)
}

// Error logs arguments at ERROR level.
func Error(args ...any) {
	pkg().Error(args...)
}

// Panic logs arguments at PANIC level and panics.
//
// Defers will be executed.
func Panic(args ...any) {
	SyncFileLogger()
	pkg().Panic(args...)
}

// Fatal logs arguments at FATAL level and calls os.Exit.
//
// Defers will not be executed.
func Fatal(args ...any) {
	SyncFileLogger()
	pkg().Fatal(args...)
}
