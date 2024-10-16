package jour

import (
	"context"
)

// contextKey is how we find Logger in a context.Context.
type contextKey struct{}

// FromContext returns the Logger associated with ctx, or the default Logger.
func FromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(contextKey{}).(*Logger); ok {
		return l
	}
	return DefaultLogger
}

// NewContext returns a new context with the given Logger.
func NewContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}
