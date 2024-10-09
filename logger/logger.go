package logger

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"time"

	"github.com/wzshiming/jour"
)

// DefaultLogger is the default logger
var DefaultLogger = NewLoggerWithHandler(jour.DefaultHandler, jour.LevelInfo)

// NewLogger returns a new Logger that writes to w.
func NewLogger(w io.Writer, level jour.Level) *Logger {
	return NewLoggerWithHandler(jour.NewHandler(w, level), level)
}

// NewLoggerWithHandler returns a new Logger with give Handler
func NewLoggerWithHandler(handler jour.Handler, level jour.Level) *Logger {
	return &Logger{handler, level}
}

// Logger is a wrapper around jour.Handler.
type Logger struct {
	handler jour.Handler
	level   jour.Level // Level specifies a level of verbosity for V logs.
}

// Log logs a message with the given level.
func (l *Logger) Log(ctx context.Context, level jour.Level, msg string, args ...any) {
	l.log(ctx, level, msg, args...)
}

// LogAttr logs a message with the given level.
func (l *Logger) LogAttr(ctx context.Context, level jour.Level, msg string, attrs ...jour.Attr) {
	l.logAttr(ctx, level, msg, attrs...)
}

// Debug calls [Logger.Debug] on the Logger.
func (l *Logger) Debug(msg string, args ...any) {
	l.log(context.Background(), jour.LevelDebug, msg, args...)
}

// DebugContext calls [Logger.DebugContext] on the Logger.
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, jour.LevelDebug, msg, args...)
}

// Info calls [Logger.Info] on the Logger.
func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), jour.LevelInfo, msg, args...)
}

// InfoContext calls [Logger.InfoContext] on the Logger.
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, jour.LevelInfo, msg, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), jour.LevelWarn, msg, args...)
}

// WarnContext calls [Logger.WarnContext] on the Logger.
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, jour.LevelWarn, msg, args...)
}

// Error logs an error message.
func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), jour.LevelError, msg, args...)
}

// ErrorContext calls [Logger.ErrorContext] on the Logger.
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, jour.LevelError, msg, args...)
}

// log is the low-level logging method for methods that take ...any.
// It must always be called directly by an exported logging method
// or function, because it uses a fixed call depth to obtain the pc.
// copied from slog.Logger
func (l *Logger) log(ctx context.Context, level jour.Level, msg string, args ...any) {
	if !l.handler.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)

	var (
		attr  jour.Attr
		attrs []jour.Attr
	)
	for len(args) > 0 {
		attr, args = argsToAttr(args)
		attrs = append(attrs, attr)
	}
	r.AddAttrs(attrs...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.handler.Handle(ctx, r)
}

func (l *Logger) logAttr(ctx context.Context, level jour.Level, msg string, attrs ...jour.Attr) {
	if !l.handler.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.AddAttrs(attrs...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.handler.Handle(ctx, r)
}

// With returns a new Logger that includes the given arguments.
func (l *Logger) With(args ...any) *Logger {
	var (
		attr  jour.Attr
		attrs []jour.Attr
	)
	for len(args) > 0 {
		attr, args = argsToAttr(args)
		attrs = append(attrs, attr)
	}
	return NewLoggerWithHandler(l.handler.WithAttrs(attrs), l.level)
}

// WithGroup returns a new Logger that starts a group. The keys of all
// attributes added to the Logger will be qualified by the given name.
func (l *Logger) WithGroup(name string) *Logger {
	return NewLoggerWithHandler(l.handler.WithGroup(name), l.level)
}

// Level returns the level of the Logger
func (l *Logger) Level() jour.Level {
	return l.level
}

// Handler returns the handler of the Logger
func (l *Logger) Handler() jour.Handler {
	return l.handler
}

const badKey = "!BADKEY"

// argsToAttr turns a prefix of the nonempty args slice into an Attr
// and returns the unconsumed portion of the slice.
// If args[0] is an Attr, it returns it.
// If args[0] is a string, it treats the first two elements as
// a key-value pair.
// Otherwise, it treats args[0] as a value with a missing key.
// copied from slog.Logger
func argsToAttr(args []any) (jour.Attr, []any) {
	switch x := args[0].(type) {
	case string:
		if len(args) == 1 {
			return slog.String(badKey, x), nil
		}
		return slog.Any(x, args[1]), args[2:]
	case jour.Attr:
		return x, args[1:]
	case error:
		return slog.String("err", x.Error()), args[1:]
	default:
		return slog.Any(badKey, x), args[1:]
	}
}
