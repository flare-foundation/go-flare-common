// Package logger provides a structured logging framework built on top of zap with console and file output support.
package logger

import (
	"io"
	"os"
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	timeFormat = "[01-02|15:04:05.000]"
)

var loggerPtr atomic.Pointer[zap.SugaredLogger]

func init() {
	loggerPtr.Store(createSugared(DefaultConfig()))
}

func current() *zap.SugaredLogger {
	return loggerPtr.Load()
}

// Config holds logger configuration for output level, file, and console settings.
type Config struct {
	Level       string `toml:"level"` // valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (zap)
	File        string `toml:"file"`
	MaxFileSize int    `toml:"max_file_size"` // In megabytes
	Console     bool   `toml:"console"`
}

// DefaultConfig returns the default logger configuration.
//
//	Level: "DEBUG"
//	Console: true
func DefaultConfig() Config {
	return Config{
		Level:   "DEBUG",
		Console: true,
	}
}

// Logger returns the global sugared logger instance.
func Logger() *zap.SugaredLogger {
	return current()
}

// Set configures logger according to Config. Safe to call concurrently
// with logging calls.
func Set(cfg Config) {
	loggerPtr.Store(createSugared(cfg))
}

func createSugared(config Config) *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()
	cores := make([]zapcore.Core, 0)
	if config.Console {
		cores = append(cores, createConsoleLoggerCore(atom))
	}
	if len(config.File) > 0 {
		cores = append(cores, createFileLoggerCore(config, atom))
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core,
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	sugared := logger.Sugar()

	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		// Fall back to DEBUG (the DefaultConfig level) rather than the
		// zero value of zapcore.Level (which is InfoLevel). Silently
		// downgrading to INFO would drop messages the operator likely
		// intended to see; DEBUG keeps everything visible — including
		// this very Errorf, which the surrounding atom.SetLevel below
		// must pass through. Reload-time Set() callers see the same
		// fallback path.
		level = zapcore.DebugLevel
		sugared.Errorf("invalid logger level %q; falling back to DEBUG", config.Level)
	}
	atom.SetLevel(level)
	return sugared
}

// SyncFileLogger synchronizes the file logger (but not the console logger). It is
// automatically called during fatal or panic log events. If you need to manually
// synchronize the logger at other points in your application, you can invoke this function as needed.
func SyncFileLogger() {
	current().Infof("Syncing file logger.")
	err := current().Sync()
	if err != nil {
		current().Infof("Failed to sync logger: %v", err)
	}
}

func createFileLoggerCore(config Config, atom zap.AtomicLevel) zapcore.Core {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: config.File,
		MaxSize:  config.MaxFileSize,
	})
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeLevel = fileLevelEncoder
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		w,
		atom,
	)
}

type noSyncWriter struct {
	io.Writer
}

func (n noSyncWriter) Sync() error {
	return nil
}

func createConsoleLoggerCore(atom zap.AtomicLevel) zapcore.Core {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeLevel = consoleColorLevelEncoder
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		noSyncWriter{os.Stdout},
		atom,
	)
}

func consoleColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	s, ok := levelToCapitalColorString[l]
	if !ok {
		s = unknownLevelColor.Wrap(l.CapitalString())
	}
	enc.AppendString(s)
}

func fileLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(l.CapitalString())
}

// Debugf formats the message and logs it at DEBUG level.
func Debugf(msg string, args ...any) {
	current().Debugf(msg, args...)
}

// Infof formats the message and logs it at INFO level.
func Infof(msg string, args ...any) {
	current().Infof(msg, args...)
}

// Warnf formats the message and logs it at WARN level.
func Warnf(msg string, args ...any) {
	current().Warnf(msg, args...)
}

// Errorf formats the message and logs it at ERROR level.
func Errorf(msg string, args ...any) {
	current().Errorf(msg, args...)
}

// Panicf formats the message and logs it at PANIC level and panics.
//
// Defers will be executed.
func Panicf(msg string, args ...any) {
	SyncFileLogger()
	current().Panicf(msg, args...)
}

// Fatalf formats the message and logs it at FATAL level and calls os.Exit.
//
// Defers will not be executed.
func Fatalf(msg string, args ...any) {
	SyncFileLogger()
	current().Fatalf(msg, args...)
}

// Debug logs arguments at DEBUG level.
func Debug(args ...any) {
	current().Debug(args...)
}

// Info logs arguments at INFO level.
func Info(args ...any) {
	current().Info(args...)
}

// Warn logs arguments at WARN level.
func Warn(args ...any) {
	current().Warn(args...)
}

// Error logs arguments at ERROR level.
func Error(args ...any) {
	current().Error(args...)
}

// Panic logs arguments at PANIC level and panics.
//
// Defers will be executed.
func Panic(args ...any) {
	SyncFileLogger()
	current().Panic(args...)
}

// Fatal logs arguments at FATAL level and calls os.Exit.
//
// Defers will not be executed.
func Fatal(args ...any) {
	SyncFileLogger()
	current().Fatal(args...)
}
