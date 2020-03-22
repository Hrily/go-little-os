package logger

import (
	"kernel/lib/io/serial"
)

const (
	Info     Verbosity = "[INFO]:  "
	Error    Verbosity = "[ERROR]: "
	Debug    Verbosity = "[DEBUG]: "
	_newline           = "\n"
	_space             = " "
	_zero              = 0x30
	_a                 = 0x61
)

// comLogger is the serial com logger
type comLogger struct{}

type Verbosity string

// _comLogger is the singleton of logger
var _comLogger = comLogger{}

// COM returns a com logger
func COM() *comLogger {
	return &_comLogger
}

// Info logs an info msg
func (l *comLogger) Info(msg string) {
	l.Log(Info, msg)
}

// Error logs an error msg
func (l *comLogger) Error(msg string) {
	l.Log(Error, msg)
}

// Debug logs a debug msg
func (l *comLogger) Debug(msg string) {
	l.Log(Debug, msg)
}

// Log logs message with given verbosity
func (l *comLogger) Log(verbosity Verbosity, msg string) {
	serial.COM1().Write(string(verbosity))
	serial.COM1().Write(msg)
	serial.COM1().Write(_newline)
}

// LogUint logs an integer as hex along with a msg
func (l *comLogger) LogUint(verbosity Verbosity, msg string, n uint64) {
	serial.COM1().Write(string(verbosity))
	serial.COM1().Write(msg)
	serial.COM1().Write(_space)
	l.writeUint(n)
	serial.COM1().Write(_newline)
}

func (l *comLogger) writeUint(n uint64) {
	hexChars := [18]byte{
		'0', 'x',
		'0', '0', '0', '0',
		'0', '0', '0', '0',
		'0', '0', '0', '0',
		'0', '0', '0', '0',
	}
	i := len(hexChars) - 1
	for n > 0 {
		c := n & 0xf
		if c < 10 {
			hexChars[i] = byte(_zero + c)
		} else if c <= 36 {
			hexChars[i] = byte(_a + c - 10)
		}
		n = n >> 4
		i--
	}
	for i := 0; i < len(hexChars); i++ {
		serial.COM1().WriteB(hexChars[i])
	}
}
