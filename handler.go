package jour

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"golang.org/x/term"
)

type (
	// Level is the logging level.
	Level = slog.Level
	// Handler is the logging handler.
	Handler = slog.Handler
	// Attr is a key-value pair.
	Attr = slog.Attr
	// Record holds information about a log event
	Record = slog.Record
)

// DefaultHandler is the default handler
var DefaultHandler = NewHandler(os.Stderr, LevelInfo)

// The following is Level definitions copied from slog.
const (
	LevelDebug Level = slog.LevelDebug
	LevelInfo  Level = slog.LevelInfo
	LevelWarn  Level = slog.LevelWarn
	LevelError Level = slog.LevelError
)

func NewHandler(w io.Writer, level Level) Handler {
	if w == nil {
		return noopHandler{}
	}

	if file, ok := w.(*os.File); ok {
		fd := int(file.Fd())
		if term.IsTerminal(fd) {
			return newCtlHandler(w, fd, level)
		}
	}

	handler := &slog.HandlerOptions{
		AddSource:   true,
		Level:       level,
		ReplaceAttr: jsonReplaceAttr,
	}
	return slog.NewJSONHandler(w, handler)
}

func jsonReplaceAttr(groups []string, a Attr) Attr {
	if a.Value.Kind() == slog.KindDuration {
		if t, ok := a.Value.Any().(time.Duration); ok {
			return slog.Any(a.Key, durationFormat{
				Nanosecond: int64(t),
				Human:      t.String(),
			})
		}
	}
	if a.Value.Kind() == slog.KindAny {
		if t, ok := a.Value.Any().(fmt.Stringer); ok {
			return slog.Attr{
				Key:   a.Key,
				Value: slog.StringValue(t.String()),
			}
		}
	}
	return a
}

// durationFormat is the format used to print time.Duration in both nanosecond and string.
type durationFormat struct {
	Nanosecond int64  `json:"nanosecond"`
	Human      string `json:"human"`
}
