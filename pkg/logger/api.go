package logger

import (
	"fmt"
	"log/slog"
	"os"
)

// Provides wrappers for structured, printf and print logging.

func (l *Log) Debugw(msg string, keysAndValues ...any) {
	l.log(slog.LevelDebug, msg, keysAndValues...)
}

func (l *Log) Infow(msg string, keysAndValues ...any) {
	l.log(slog.LevelInfo, msg, keysAndValues...)
}

func (l *Log) Warnw(msg string, keysAndValues ...any) {
	l.log(slog.LevelWarn, msg, keysAndValues...)
}

func (l *Log) Errorw(msg string, keysAndValues ...any) {
	l.log(slog.LevelError, msg, keysAndValues...)
}

// Fatalw logs msg and fields at FATAL level and exits with status 1.
// Deferred functions do not run.
func (l *Log) Fatalw(msg string, keysAndValues ...any) {
	l.log(LevelFatal, msg, keysAndValues...)
	os.Exit(1)
}

// Panicw logs msg and fields at FATAL level.
// It then panics with the message, allowing deferred functions to run.
func (l *Log) Panicw(msg string, keysAndValues ...any) {
	l.log(LevelFatal, msg, keysAndValues...)
	panic(msg)
}

func (l *Log) Debugf(format string, args ...any) {
	l.logf(slog.LevelDebug, format, args...)
}

func (l *Log) Infof(format string, args ...any) {
	l.logf(slog.LevelInfo, format, args...)
}

func (l *Log) Warnf(format string, args ...any) {
	l.logf(slog.LevelWarn, format, args...)
}

func (l *Log) Errorf(format string, args ...any) {
	l.logf(slog.LevelError, format, args...)
}

// Fatalf logs a formatted message at FATAL level and exits with status 1.
// Deferred functions do not run.
func (l *Log) Fatalf(format string, args ...any) {
	l.logf(LevelFatal, format, args...)
	os.Exit(1)
}

// Panicf logs a formatted message at FATAL level.
// It then panics with the message, allowing deferred functions to run.
func (l *Log) Panicf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.log(LevelFatal, msg)
	panic(msg)
}

func (l *Log) Debug(args ...any) { l.logPrint(slog.LevelDebug, args...) }

func (l *Log) Info(args ...any) { l.logPrint(slog.LevelInfo, args...) }

func (l *Log) Warn(args ...any) { l.logPrint(slog.LevelWarn, args...) }

func (l *Log) Error(args ...any) { l.logPrint(slog.LevelError, args...) }

// Fatal logs its arguments at FATAL level using fmt.Sprint, then exits with status 1.
// Deferred functions do not run.
func (l *Log) Fatal(args ...any) {
	l.logPrint(LevelFatal, args...)
	os.Exit(1)
}

// Panic logs its arguments at FATAL level using fmt.Sprint.
// It then panics with the message, allowing deferred functions to run.
func (l *Log) Panic(args ...any) {
	msg := fmt.Sprint(args...)
	l.log(LevelFatal, msg)
	panic(msg)
}

// The package functions write through the logger set by Set.

func Debugw(msg string, keysAndValues ...any) { Logger().log(slog.LevelDebug, msg, keysAndValues...) }

func Infow(msg string, keysAndValues ...any) { Logger().log(slog.LevelInfo, msg, keysAndValues...) }

func Warnw(msg string, keysAndValues ...any) { Logger().log(slog.LevelWarn, msg, keysAndValues...) }

func Errorw(msg string, keysAndValues ...any) { Logger().log(slog.LevelError, msg, keysAndValues...) }

// Fatalw logs msg and fields at FATAL level and exits with status 1.
// Deferred functions do not run.
func Fatalw(msg string, keysAndValues ...any) {
	Logger().log(LevelFatal, msg, keysAndValues...)
	os.Exit(1)
}

// Panicw logs msg and fields at FATAL level.
// It then panics with the message, allowing deferred functions to run.
func Panicw(msg string, keysAndValues ...any) {
	Logger().log(LevelFatal, msg, keysAndValues...)
	panic(msg)
}

func Debugf(format string, args ...any) { Logger().logf(slog.LevelDebug, format, args...) }

func Infof(format string, args ...any) { Logger().logf(slog.LevelInfo, format, args...) }

func Warnf(format string, args ...any) { Logger().logf(slog.LevelWarn, format, args...) }

func Errorf(format string, args ...any) { Logger().logf(slog.LevelError, format, args...) }

// Fatalf logs a formatted message at FATAL level and exits with status 1.
// Deferred functions do not run.
func Fatalf(format string, args ...any) {
	Logger().logf(LevelFatal, format, args...)
	os.Exit(1)
}

// Panicf logs a formatted message at FATAL level.
// It then panics with the message, allowing deferred functions to run.
func Panicf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	Logger().log(LevelFatal, msg)
	panic(msg)
}

func Debug(args ...any) { Logger().logPrint(slog.LevelDebug, args...) }

func Info(args ...any) { Logger().logPrint(slog.LevelInfo, args...) }

func Warn(args ...any) { Logger().logPrint(slog.LevelWarn, args...) }

func Error(args ...any) { Logger().logPrint(slog.LevelError, args...) }

// Fatal logs its arguments at FATAL level using fmt.Sprint, then exits with status 1.
// Deferred functions do not run.
func Fatal(args ...any) {
	Logger().logPrint(LevelFatal, args...)
	os.Exit(1)
}

// Panic logs its arguments at FATAL level using fmt.Sprint.
// It then panics with the message, allowing deferred functions to run.
func Panic(args ...any) {
	msg := fmt.Sprint(args...)
	Logger().log(LevelFatal, msg)
	panic(msg)
}
