package str

import (
	"strings"
	"unicode/utf8"
)

// WrapOption configures the behavior of Wrap.
type WrapOption func(*wrapConfig)

type wrapBehavior int

const (
	behaviorBeforeWord wrapBehavior = iota
	behaviorAfterWord
	behaviorHardBreak
)

type wrapConfig struct {
	behavior  wrapBehavior
	lineBreak string
	indent    string
}

// WrapBeforeWord wraps before the word that would extend past the line width.
var WrapBeforeWord WrapOption = func(c *wrapConfig) { c.behavior = behaviorBeforeWord }

// WrapAfterWord wraps after the word that extends past the line width.
var WrapAfterWord WrapOption = func(c *wrapConfig) { c.behavior = behaviorAfterWord }

// WrapHardBreak breaks mid-word at exactly the line width.
var WrapHardBreak WrapOption = func(c *wrapConfig) { c.behavior = behaviorHardBreak }

// WithLineBreak sets the line break string (default "\n").
func WithLineBreak(lb string) WrapOption {
	return func(c *wrapConfig) { c.lineBreak = lb }
}

// WithIndent sets a prefix for continuation lines.
func WithIndent(indent string) WrapOption {
	return func(c *wrapConfig) { c.indent = indent }
}

// Wrap returns a function that wraps text at the specified width (in runes).
// Default behavior is WrapBeforeWord with "\n" line breaks.
// Width of 0 returns input unchanged.
func Wrap(width int, opts ...WrapOption) func(string) string {
	cfg := &wrapConfig{
		behavior:  behaviorBeforeWord,
		lineBreak: "\n",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return func(s string) string {
		if width <= 0 || s == "" {
			return s
		}
		switch cfg.behavior {
		case behaviorHardBreak:
			return wrapHard(s, width, cfg)
		case behaviorAfterWord:
			return wrapAfterWord(s, width, cfg)
		default:
			return wrapBeforeWord(s, width, cfg)
		}
	}
}

func wrapHard(s string, width int, cfg *wrapConfig) string {
	runes := []rune(s)
	var b strings.Builder
	for i := 0; i < len(runes); {
		if i > 0 {
			b.WriteString(cfg.lineBreak)
			b.WriteString(cfg.indent)
		}
		end := min(i+width, len(runes))
		b.WriteString(string(runes[i:end]))
		i = end
	}
	return b.String()
}

func wrapBeforeWord(s string, width int, cfg *wrapConfig) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	var b strings.Builder
	lineLen := 0
	for i, word := range words {
		wLen := utf8.RuneCountInString(word)
		if i == 0 {
			b.WriteString(word)
			lineLen = wLen
			continue
		}
		if lineLen+1+wLen > width {
			b.WriteString(cfg.lineBreak)
			b.WriteString(cfg.indent)
			b.WriteString(word)
			lineLen = utf8.RuneCountInString(cfg.indent) + wLen
		} else {
			b.WriteByte(' ')
			b.WriteString(word)
			lineLen += 1 + wLen
		}
	}
	return b.String()
}

func wrapAfterWord(s string, width int, cfg *wrapConfig) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	var b strings.Builder
	lineLen := 0
	for i, word := range words {
		wLen := utf8.RuneCountInString(word)
		if i == 0 {
			b.WriteString(word)
			lineLen = wLen
			continue
		}
		b.WriteByte(' ')
		b.WriteString(word)
		lineLen += 1 + wLen
		if lineLen >= width && i < len(words)-1 {
			b.WriteString(cfg.lineBreak)
			b.WriteString(cfg.indent)
			lineLen = utf8.RuneCountInString(cfg.indent)
		}
	}
	return b.String()
}
