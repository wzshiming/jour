package jour

import (
	"context"
)

type noopHandler struct{}

var _ Handler = noopHandler{}

func (noopHandler) Enabled(_ context.Context, _ Level) bool {
	return false
}

func (noopHandler) Handle(_ context.Context, _ Record) error {
	return nil
}

func (h noopHandler) WithAttrs(attrs []Attr) Handler {
	return h
}

func (h noopHandler) WithGroup(name string) Handler {
	return h
}
