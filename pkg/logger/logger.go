package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// LogLevel definition
type LogLevel string

// LogLevel constants
const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

var (
	once      sync.Once
	singleton *zap.Logger
)

// Init the instance of logger
func Init(debug bool) {
	once.Do(func() {
		var err error
		if debug {
			singleton, err = zap.NewDevelopment()
		} else {
			singleton, err = zap.NewProduction()
		}
		if err != nil {
			panic(err)
		}
	})
}

// Instance the logger instance
func Instance() *zap.Logger {
	return singleton
}

// SetLevel to set the new level for logger
func SetLevel(logLevel string) {
	switch LogLevel(logLevel) {
	case DebugLevel:
		singleton.Core().Enabled(zap.DebugLevel)
	case InfoLevel:
		singleton.Core().Enabled(zap.InfoLevel)
	case WarnLevel:
		singleton.Core().Enabled(zap.WarnLevel)
	case ErrorLevel:
		singleton.Core().Enabled(zap.ErrorLevel)
	default:
		singleton.Core().Enabled(zap.InfoLevel)
	}
}

// Printf is converted to an Info call because Zap does not have Printf()
func Printf(s string, i ...interface{}) {
	singleton.Info(fmt.Sprintf(s, i...))
}

// Debug logs as the debug level
func Debug(i ...interface{}) {
	singleton.Debug(fmt.Sprint(i...))
}

// Debugf logs as the debug level
func Debugf(s string, i ...interface{}) {
	singleton.Debug(fmt.Sprintf(s, i...))
}

// Info logs as the info level
func Info(i ...interface{}) {
	singleton.Info(fmt.Sprint(i...))
}

// Infof logs as the info level
func Infof(s string, i ...interface{}) {
	singleton.Info(fmt.Sprintf(s, i...))
}

// Warn logs as the WARN level
func Warn(i ...interface{}) {
	singleton.Warn(fmt.Sprint(i...))
}

// Warnf logs as the WARN level
func Warnf(s string, i ...interface{}) {
	singleton.Warn(fmt.Sprintf(s, i...))
}

// Error logs as the ERROR level
func Error(i ...interface{}) {
	singleton.Error(fmt.Sprint(i...))
}

// Errorf logs as the ERROR level
func Errorf(s string, i ...interface{}) {
	singleton.Error(fmt.Sprintf(s, i...))
}
