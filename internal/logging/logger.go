package logging

import (
	"context"

	"github.com/go-chi/chi/middleware"
)

// Logger provides a wrapper around logging Request IDs with different logging levels
type Logger struct {
	stream Stream
	ctx    *context.Context
}

func (l *Logger) Debug(msg string) {
	l.print(logLevelDebug, msg)
}

func (l *Logger) Error(msg string) {
	l.print(logLevelError, msg)
}

func (l *Logger) Info(msg string) {
	l.print(logLevelInfo, msg)
}

func (l *Logger) print(level logLevel, message string) {
	requestID := ""

	if l.ctx != nil {
		requestID = middleware.GetReqID(*l.ctx)
	}

	l.stream.Printf("[%s] %s: %s", requestID, level, message)
}

// New creates a new Logger instances with the output stream configured.
// This should only be called once within the scope of an application
func New(stream Stream) *Logger {
	return &Logger{
		stream: stream,
	}
}

// WithContext creates a new instance of Logger, but with the context set. This should be invoked
// within each request rather than being set as a field of a struct to ensure context safety
func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{
		stream: l.stream,
		ctx:    &ctx,
	}
}
