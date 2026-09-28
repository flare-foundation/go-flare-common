package logger

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

// ANSI colors for levels.
var levelColor = map[slog.Level]string{
	slog.LevelDebug: "\x1b[36m", // cyan
	slog.LevelInfo:  "\x1b[32m", // green
	slog.LevelWarn:  "\x1b[33m", // yellow
	slog.LevelError: "\x1b[31m", // red
	LevelFatal:      "\x1b[35m", // magenta
}

const ansiReset = "\x1b[0m"

// consoleHandler writes one tab-separated line per record:
//
//	2026-09-28T13:42:00.001Z	INFO	core/engine.go:441	Batch indexed	chain_id=14 from_block=41200000
//
// Fields use key=value notation. Values are quoted when they contain separators
// or control characters.
type consoleHandler struct {
	level  slog.Leveler
	out    io.Writer
	mu     *sync.Mutex
	color  bool
	fields string // preformatted, from WithAttrs
	group  string // dotted prefix, from WithGroup
}

func newConsoleHandler(out io.Writer, level slog.Leveler, color bool) *consoleHandler {
	return &consoleHandler{level: level, out: out, mu: new(sync.Mutex), color: color}
}

func (h *consoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

func (h *consoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	c := *h
	for _, a := range attrs {
		c.fields = appendField(c.fields, h.group, a)
	}

	return &c
}

func (h *consoleHandler) WithGroup(name string) slog.Handler {
	c := *h
	c.group = h.group + name + "."

	return &c
}

func (h *consoleHandler) Handle(_ context.Context, r slog.Record) error {
	var b strings.Builder

	b.WriteString(r.Time.UTC().Format(timeLayout))
	b.WriteByte('\t')
	b.WriteString(h.levelString(r.Level))
	b.WriteByte('\t')
	if r.PC != 0 {
		frame, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()
		b.WriteString(sourceRef(frame.File, frame.Line))
		b.WriteByte('\t')
	}
	b.WriteString(oneLine(r.Message))

	fields := h.fields
	r.Attrs(func(a slog.Attr) bool {
		fields = appendField(fields, h.group, a)

		return true
	})
	if fields != "" {
		b.WriteByte('\t')
		b.WriteString(fields)
	}
	b.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := io.WriteString(h.out, b.String())

	return err
}

// levelString returns the level name, with color if enabled.
func (h *consoleHandler) levelString(l slog.Level) string {
	name := levelName(l)
	if c, coloured := levelColor[l]; coloured && h.color {
		return c + name + ansiReset
	}

	return name
}

// appendField adds a as key=value, flattening a group into dotted keys.
func appendField(fields, prefix string, a slog.Attr) string {
	a.Value = a.Value.Resolve()
	if a.Equal(slog.Attr{}) {
		return fields
	}

	if a.Value.Kind() == slog.KindGroup {
		// slog inlines a group with no name, so it adds no prefix either.
		inner := prefix
		if a.Key != "" {
			inner = prefix + a.Key + "."
		}

		for _, g := range a.Value.Group() {
			fields = appendField(fields, inner, g)
		}

		return fields
	}

	if fields != "" {
		fields += " "
	}

	return fields + oneLine(prefix+a.Key) + "=" + value(a.Value)
}

// value renders v, quoting it only when bare text would be ambiguous.
func value(v slog.Value) string {
	var s string
	switch v.Kind() {
	case slog.KindDuration:
		s = v.Duration().String()
	case slog.KindTime:
		s = v.Time().UTC().Format(timeLayout)
	default:
		// Value.String uses fmt, which handles nil pointers and recovers from
		// panics in Error or String methods.
		s = v.String()
	}

	if needsQuote(s) {
		return strconv.Quote(s)
	}

	return s
}

// needsQuote checks for empty values, separators and control characters.
func needsQuote(s string) bool {
	if s == "" {
		return true
	}

	return strings.ContainsFunc(s, func(r rune) bool {
		return r == ' ' || r == '"' || r == '=' || unicode.IsControl(r)
	})
}

// oneLine escapes control characters so messages and keys stay on one line.
func oneLine(s string) string {
	if !strings.ContainsFunc(s, unicode.IsControl) {
		return s
	}

	quoted := strconv.Quote(s)

	return quoted[1 : len(quoted)-1]
}
