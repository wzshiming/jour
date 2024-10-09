package logger

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wzshiming/jour"
)

// ParseLevel parses a level string.
func ParseLevel(s string) (l jour.Level, err error) {
	name := s
	offsetStr := ""
	i := strings.IndexAny(s, "+-")
	if i > 0 {
		name = s[:i]
		offsetStr = s[i:]
	} else if i == 0 ||
		(name[0] >= '0' && name[0] <= '9') {
		name = "INFO"
		offsetStr = s
	}

	switch strings.ToUpper(name) {
	case "DEBUG":
		l = jour.LevelDebug
	case "INFO":
		l = jour.LevelInfo
	case "WARN":
		l = jour.LevelWarn
	case "ERROR":
		l = jour.LevelError
	default:
		return 0, fmt.Errorf("ParseLevel %q: invalid level name", s)
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return 0, fmt.Errorf("ParseLevel %q: invalid offset: %w", s, err)
		}
		l += jour.Level(offset)
	}

	return l, nil
}
