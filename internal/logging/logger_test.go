package logging

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-chi/chi/middleware"

	"github.com/stretchr/testify/assert"
)

type mockStream struct {
	messages []string
}

func (l *mockStream) Printf(format string, v ...interface{}) {
	l.messages = append(l.messages, fmt.Sprintf(format, v...))
}

func TestLoggerWritesCorrectPrefixes(t *testing.T) {
	tests := map[string]struct {
		run              func(*Logger)
		expectedMessages []string
	}{
		"Writes message with [INFO]": {
			func(l *Logger) {
				l.Info("HELLO")
			},
			[]string{"[] INFO: HELLO"},
		},
		"Writes message with [DEBUG]": {
			func(l *Logger) {
				l.Debug("HELLO")
			},
			[]string{"[] DEBUG: HELLO"},
		},
		"Writes message with [ERROR]": {
			func(l *Logger) {
				l.Error("HELLO")
			},
			[]string{"[] ERROR: HELLO"},
		},
	}

	for _, test := range tests {
		log := &mockStream{}
		l := New(log)

		test.run(l)

		assert.Equal(t, test.expectedMessages, log.messages)
	}
}

func TestLoggerWritesRequestIDToLog(t *testing.T) {
	tests := map[string]struct {
		getContext       func() context.Context
		expectedMessages []string
	}{
		"Renders empty string if request id is not set": {
			func() context.Context {
				return context.Background()
			},
			[]string{"[] INFO: my message"},
		},
		"Renders request id if set": {
			func() context.Context {
				return context.WithValue(context.Background(), middleware.RequestIDKey, "my-key")
			},
			[]string{"[my-key] INFO: my message"},
		},
	}

	for _, test := range tests {
		stream := &mockStream{}
		l := New(stream)

		l = l.WithContext(test.getContext())

		l.Info("my message")
		assert.Equal(t, test.expectedMessages, stream.messages)
	}
}
