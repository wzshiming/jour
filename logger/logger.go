package logger

import (
	"io"
	"log/slog"

	"github.com/wzshiming/jour"
)

// DefaultLogger is the default logger
var DefaultLogger = NewLoggerWithHandler(jour.DefaultHandler)

// NewLogger returns a new Logger that writes to w.
func NewLogger(w io.Writer, level jour.Level) *Logger {
	return NewLoggerWithHandler(jour.NewHandler(w, level))
}

// NewLoggerWithHandler returns a new Logger with give Handler
func NewLoggerWithHandler(handler jour.Handler) *Logger {
	return slog.New(handler)
}

// Logger is a wrapper around jour.Handler.
type Logger = slog.Logger
