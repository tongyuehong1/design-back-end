/*
 * Revision History:
 *     Initial: 2018/05/24        Tong Yuehong
 */

package log

import (
	"encoding/json"
	"fmt"
)

type logger interface {
	Debug(message string, fields ...interface{})
	Info(message string, fields ...interface{})
	Warn(message string, fields ...interface{})
	Error(message string, fields ...interface{})
}

// LogFunc - .
type LogFunc func(string, ...interface{})

var globalLogger Logger

// Debug - default funcation about debug.
var Debug LogFunc = defaultDebugLog

// Info -
var Info LogFunc = defaultInfoLog

// Warn -
var Warn LogFunc = defaultWarnLog

// Error -
var Error LogFunc = defaultErrorLog

// InitGlobalLogger -
func InitGlobalLogger(logger Logger) {
	globalLogger = logger
	Debug = globalLogger.Debug
	Info = globalLogger.Info
	Warn = globalLogger.Warn
	Error = globalLogger.Error
}

// defaultLog manually encodes the log to STDOUT, providing a basic, default logging implementation
// before mlog is fully configured.
func defaultLog(level, msg string, fields ...interface{}) {
	log := struct {
		Level   string        `json:"level"`
		Message string        `json:"msg"`
		Fields  []interface{} `json:"fields,omitempty"`
	}{
		level,
		msg,
		fields,
	}

	if b, err := json.Marshal(log); err != nil {
		fmt.Printf(`{"level":"error","msg":"failed to encode log message"}%s`, "\n")
	} else {
		fmt.Printf("%s\n", b)
	}
}

func defaultDebugLog(msg string, fields ...interface{}) {
	defaultLog("debug", msg, fields...)
}

func defaultInfoLog(msg string, fields ...interface{}) {
	defaultLog("info", msg, fields...)
}

func defaultWarnLog(msg string, fields ...interface{}) {
	defaultLog("warn", msg, fields...)
}

func defaultErrorLog(msg string, fields ...interface{}) {
	defaultLog("error", msg, fields...)
}
