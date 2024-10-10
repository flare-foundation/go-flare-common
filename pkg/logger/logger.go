package logger

import (
	"errors"
	"log"
	"os"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	timeFormat = "[01-02|15:04:05.000]"
)

var (
	sugaredLogger *zap.SugaredLogger
)

func init() {
	sugaredLogger = createSugared(DefaultConfig())
}

type Config struct {
	Level       string `toml:"level"` // valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (zap)
	File        string `toml:"file"`
	MaxFileSize int    `toml:"max_file_size"` // In megabytes
	Console     bool   `toml:"console"`
}

func DefaultConfig() Config {
	return Config{
		Level:   "DEBUG",
		Console: true,
	}
}

// Set configures logger according to Config.
func Set(cfg Config) {
	createSugared(cfg)
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
	defer func() {
		err := logger.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) && !errors.Is(err, syscall.EBADF) {
			log.Print("Failed to sync logger", err)
		}
	}()

	sugaredLogger = logger.Sugar()

	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		sugaredLogger.Errorf("Wrong level %s", config.Level)
	}
	atom.SetLevel(level)
	return sugaredLogger
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

func createConsoleLoggerCore(atom zap.AtomicLevel) zapcore.Core {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeLevel = consoleColorLevelEncoder
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
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
func Debugf(msg string, args ...interface{}) {
	sugaredLogger.Debugf(msg, args...)
}

// Infof formats the message and logs it at INFO level.
func Infof(msg string, args ...interface{}) {
	sugaredLogger.Infof(msg, args...)
}

// Warnf formats the message and logs it at WARN level.
func Warnf(msg string, args ...interface{}) {
	sugaredLogger.Warnf(msg, args...)
}

// Errorf formats the message and logs it at ERROR level.
func Errorf(msg string, args ...interface{}) {
	sugaredLogger.Errorf(msg, args...)
}

// Panicf formats the message and logs it at PANIC level and panics.
//
// Defers will be executed.
func Panicf(msg string, args ...interface{}) {
	sugaredLogger.Panicf(msg, args...)
}

// Fatalf formats the message and logs it at FATAL level and calls os.Exit.
//
// Defers will not be executed.
func Fatalf(msg string, args ...interface{}) {
	sugaredLogger.Fatalf(msg, args...)
}

// Debug logs arguments at DEBUG level.
func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

// Info logs arguments at INFO level.
func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

// Warn logs arguments at WARN level.
func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

// Error logs arguments at ERROR level.
func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

// Panic logs arguments at PANIC level and panics.
//
// Defers will be executed.
func Panic(args ...interface{}) {
	sugaredLogger.Panic(args...)
}

// Fatal logs arguments at FATAL level  and calls os.Exit.
//
// Defers will not be executed.
func Fatal(args ...interface{}) {
	sugaredLogger.Fatal(args...)
}
