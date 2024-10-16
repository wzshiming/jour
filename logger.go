package jour

import (
	"io"
	"log/slog"
)

// DefaultLogger is the default logger
var DefaultLogger = NewLoggerWithHandler(DefaultHandler)

// NewLogger returns a new Logger that writes to w.
func NewLogger(w io.Writer, level Level) *Logger {
	return NewLoggerWithHandler(NewHandler(w, level))
}

// NewLoggerWithHandler returns a new Logger with give Handler
func NewLoggerWithHandler(handler Handler) *Logger {
	return slog.New(handler)
}

// Logger is a wrapper around jour.Handler.
type Logger = slog.Logger
