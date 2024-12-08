package logging

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

type SimpleLogger struct {
	debug bool
}

func NewLogger(debug bool) *SimpleLogger {
	return &SimpleLogger{debug: debug}
}

func (l *SimpleLogger) Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

func (l *SimpleLogger) Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}

func (l *SimpleLogger) Debug(msg string) {
	if l.debug {
		log.Printf("[DEBUG] %s", msg)
	}
}

func DefaultLogger() Logger {
	debug := os.Getenv("DEBUG") == "true"
	return NewLogger(debug)
}
