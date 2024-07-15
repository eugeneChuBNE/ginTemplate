package logging

import (
	"log"
	"os"
)

// Logger is a custom logger struct.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new Logger instance.
func NewLogger(prefix string) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile),
	}
}

// Info logs informational messages.
func (l *Logger) Info(msg string) {
	l.SetPrefix("INFO: ")
	l.Println(msg)
}

// Warning logs warning messages.
func (l *Logger) Warning(msg string) {
	l.SetPrefix("WARNING: ")
	l.Println(msg)
}

// Error logs error messages.
func (l *Logger) Error(err error) {
	l.SetPrefix("ERROR: ")
	l.Println(err)
}
