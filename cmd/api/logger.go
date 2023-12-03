package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper around a client logger.
type Logger struct {
	Client *zap.Logger
}

// Info logs a message with info level.
func (l *Logger) Info(msg string) {
	l.Client.Info(msg)
}

// Warn logs a message with warn level.
func (l *Logger) Warn(msg string) {
	l.Client.Warn(msg)
}

// Error logs a message with error level.
func (l *Logger) Error(msg string) {
	l.Client.Error(msg)
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() error {
	return l.Client.Sync()
}

// newProdLogger creates a new logger for production.
func newProdLogger() (*Logger, error) {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	if _, err := os.Stat("./tmp/api"); os.IsNotExist(err) {
		os.MkdirAll("./tmp/api", 0755)
	}

	infoFile, err := os.OpenFile("./tmp/api/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	errorFile, err := os.OpenFile("./tmp/api/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	infoWriter := zapcore.Lock(infoFile)
	errorWriter := zapcore.Lock(errorFile)

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, infoWriter, lowPriority),
		zapcore.NewCore(jsonEncoder, errorWriter, highPriority),
	)

	return &Logger{
		Client: zap.New(core),
	}, nil
}

func newDevLogger() (*Logger, error) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Logger{
		Client: zapLogger,
	}, nil
}

// newLogger creates a new logger.
func newLogger(isProd bool) (*Logger, error) {
	if isProd {
		return newProdLogger()
	}

	return newDevLogger()
}
