package jour

import (
	"context"
	"io"
	"strings"

	"golang.org/x/term"
)

type ctlHandler struct {
	level    Level
	output   io.Writer
	attrs    []Attr
	attrsStr *string
	groups   []string
	fd       int
}

func newCtlHandler(w io.Writer, fd int, level Level) Handler {
	return &ctlHandler{
		output: w,
		fd:     fd,
		level:  level,
	}
}

func (c *ctlHandler) Enabled(_ context.Context, level Level) bool {
	return level >= c.level
}

func (c *ctlHandler) Handle(_ context.Context, r Record) error {
	if r.Level < c.level {
		return nil
	}

	attrs := make([]string, 0, r.NumAttrs()+1)
	r.Attrs(func(attr Attr) bool {
		value := formatValue(attr.Value)
		if len(c.groups) == 0 {
			attrs = append(attrs, attr.Key+"="+value)
		} else {
			attrs = append(attrs, strings.Join(append(c.groups, attr.Key), ".")+"="+value)
		}
		return true
	})

	if c.attrsStr == nil {
		attrs := make([]string, 0, len(c.attrs))
		for i := len(c.attrs) - 1; i >= 0; i-- {
			attr := c.attrs[i]
			attrs = append(attrs, attr.Key+"="+formatValue(attr.Value))
		}
		attrsStr := strings.Join(attrs, " ")
		c.attrsStr = &attrsStr
	}

	if c.attrsStr != nil {
		attrs = append(attrs, *c.attrsStr)
	}

	attrsStr := ""
	if len(attrs) != 0 {
		attrsStr = strings.Join(attrs, " ")
	}

	var termWidth int
	if c.fd != 0 {
		termWidth, _, _ = term.GetSize(c.fd)
	}
	log := formatLog(r.Message, attrsStr, r.Level, termWidth)
	_, err := io.WriteString(c.output, log)
	return err
}

func (c *ctlHandler) WithAttrs(attrs []Attr) Handler {
	newAttrs := make([]Attr, 0, len(c.attrs)+len(attrs))
	newAttrs = append(newAttrs, c.attrs...)
	if len(c.groups) == 0 {
		newAttrs = append(newAttrs, attrs...)
	} else {
		for _, attr := range attrs {
			newAttrs = append(newAttrs, Attr{
				Key:   strings.Join(append(c.groups, attr.Key), "."),
				Value: attr.Value,
			})
		}
	}
	return &ctlHandler{
		fd:     c.fd,
		level:  c.level,
		output: c.output,
		attrs:  newAttrs,
		groups: c.groups,
	}
}

func (c *ctlHandler) WithGroup(name string) Handler {
	newGroups := make([]string, 0, len(c.groups)+1)
	newGroups = append(newGroups, c.groups...)
	newGroups = append(newGroups, name)
	return &ctlHandler{
		fd:     c.fd,
		level:  c.level,
		output: c.output,
		attrs:  c.attrs,
		groups: newGroups,
	}
}
