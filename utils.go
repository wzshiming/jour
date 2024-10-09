package jour

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/wzshiming/ctc"

	"github.com/wzshiming/jour/internal/format"
	"github.com/wzshiming/jour/internal/monospace"
)

type color struct {
	renderer string
	width    int
}

func newColour(c ctc.Color, msg string) color {
	return color{
		renderer: strings.Join([]string{c.String(), msg, ctc.Reset.String()}, ""),
		width:    monospace.String(msg),
	}
}

var levelColor = map[string]color{
	LevelError.String(): newColour(ctc.ForegroundRed, LevelError.String()),
	LevelWarn.String():  newColour(ctc.ForegroundYellow, LevelWarn.String()),
	LevelDebug.String(): newColour(ctc.ForegroundCyan, LevelDebug.String()),
}

func formatLog(msg string, attrs string, level Level, termWidth int) string {
	if attrs == "" {
		if level != LevelInfo {
			levelStr := level.String()
			c, ok := levelColor[strings.SplitN(levelStr, "+", 2)[0]]
			if ok {
				msg = c.renderer + " " + msg
			}
		}
		return fmt.Sprintf("%s\n", msg)
	}

	msgWidth := monospace.String(msg)
	if level != LevelInfo {
		levelStr := level.String()
		c, ok := levelColor[strings.SplitN(levelStr, "+", 2)[0]]
		if ok {
			msg = c.renderer + " " + msg
			msgWidth += c.width + 1
		}
	}
	if termWidth > msgWidth {
		return fmt.Sprintf("%s %*s\n", msg, termWidth-msgWidth-1, attrs)
	}

	return fmt.Sprintf("%s %s\n", msg, attrs)
}

func formatValue(val slog.Value) string {
	switch val.Kind() {
	case slog.KindString:
		return format.QuoteIfNeed(val.String())
	case slog.KindDuration:
		return format.HumanDuration(val.Duration())
	default:
		switch x := val.Any().(type) {
		case error:
			return format.QuoteIfNeed(x.Error())
		case fmt.Stringer:
			return format.QuoteIfNeed(x.String())
		default:
			v, err := json.Marshal(x)
			if err == nil {
				return string(v)
			}
			return format.QuoteIfNeed(val.String())
		}
	}
}
