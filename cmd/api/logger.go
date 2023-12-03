package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger is a wrapper around a client logger.
type logger struct {
	client *zap.Logger
}

// info logs a message with info level.
func (l *logger) info(msg string) {
	l.client.Info(msg)
}

// warn logs a message with warn level.
func (l *logger) warn(msg string) {
	l.client.Warn(msg)
}

// error logs a message with error level.
func (l *logger) error(msg string) {
	l.client.Error(msg)
}

// sync flushes any buffered log entries.
func (l *logger) sync() error {
	return l.client.Sync()
}

// newLogger creates a new logger.
func newLogger() (*logger, error) {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	if _, err := os.Stat("./temp/api"); os.IsNotExist(err) {
		os.MkdirAll("./temp/api", 0755)
	}

	logFile, err := os.OpenFile("./temp/api/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	errorFile, err := os.OpenFile("./temp/api/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	fileInfos := zapcore.Lock(logFile)
	fileErrors := zapcore.Lock(errorFile)

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, fileInfos, lowPriority),
		zapcore.NewCore(jsonEncoder, fileErrors, highPriority),
	)

	return &logger{
		client: zap.New(core),
	}, nil
}
