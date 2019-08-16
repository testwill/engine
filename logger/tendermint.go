package logger

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tendermint/tendermint/libs/log"
)

// logrusLogger is common interface for logrus.Logger and logrus.Entry.
type logrusLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	WithFields(fields logrus.Fields) *logrus.Entry
}

// tendermintLogger wraps logrus logger into tendermint logger.
type tendermintLogger struct {
	logrusLogger
}

// Debug logs a message at level Debug on the standard logger.
func (l *tendermintLogger) Debug(msg string, keyvals ...interface{}) {
	if len(keyvals) == 0 {
		l.logrusLogger.Debug(msg)
	} else {
		l.With(keyvals...).(*tendermintLogger).logrusLogger.Debug(msg)
	}
}

// Info logs a message at level Info on the standard logger.
func (l *tendermintLogger) Info(msg string, keyvals ...interface{}) {
	if len(keyvals) == 0 {
		l.logrusLogger.Info(msg)
	} else {
		l.With(keyvals...).(*tendermintLogger).logrusLogger.Info(msg)
	}
}

// Error logs a message at level Error on the standard logger.
func (l *tendermintLogger) Error(msg string, keyvals ...interface{}) {
	if len(keyvals) == 0 {
		l.logrusLogger.Error(msg)
	} else {
		l.With(keyvals...).(*tendermintLogger).logrusLogger.Error(msg)
	}
}

// With creates an entry from the logger and adds a key values to it.
func (l *tendermintLogger) With(keyvals ...interface{}) log.Logger {
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		} else {
			fields[fmt.Sprint(keyvals[i])] = errors.New("(MISSING)")
		}
	}
	return &tendermintLogger{
		logrusLogger: l.logrusLogger.WithFields(fields),
	}
}

// TendermintLogger returns standard logrus logger
// wrapped into tendermint logger interface.
func TendermintLogger() log.Logger {
	return &tendermintLogger{
		logrusLogger: logrus.StandardLogger(),
	}
}