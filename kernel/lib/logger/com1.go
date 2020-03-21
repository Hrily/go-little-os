package logger

import (
	"kernel/lib/io/serial"
)

const (
	_info    = "[INFO]:  "
	_error   = "[ERROR]: "
	_debug   = "[DEBUG]: "
	_newline = "\n"
)

// comLogger is the serial com logger
type comLogger struct{}

// _comLogger is the singleton of logger
var _comLogger = comLogger{}

// COM returns a com logger
func COM() *comLogger {
	return &_comLogger
}

// Info logs an info msg
func (l *comLogger) Info(msg string) {
	l.log(_info, msg)
}

// Error logs an error msg
func (l *comLogger) Error(msg string) {
	l.log(_error, msg)
}

// Debug logs a debug msg
func (l *comLogger) Debug(msg string) {
	l.log(_debug, msg)
}

// logs logs message with given verbosity
func (l *comLogger) log(verbosity, msg string) {
	serial.COM1().Write(verbosity)
	serial.COM1().Write(msg)
	serial.COM1().Write(_newline)
}
