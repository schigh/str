package str

import (
	"math"
	"strings"
)

type WrapBehavior uint

const (
	WrapBeforeWord WrapBehavior = iota
	WrapAfterWord
	WrapLiteral
)

// The user options for the wrap
type WrapOptions struct {
	Width     uint
	LineBreak string
	Behavior  WrapBehavior
}

var defaultWrapOptions *WrapOptions

func init() {
	defaultWrapOptions = &WrapOptions{
		Width:     76,
		LineBreak: "\n",
		Behavior:  WrapBeforeWord,
	}
}

// SetWrapOptions - set the global wordwrap options
func SetWrapOptions(options *WrapOptions) {
	defaultWrapOptions = options
}

// Wrap - Perform word wrap on a string
func Wrap(in string) string {
	return WrapWithOptions(in, defaultWrapOptions)
}

// WrapWithOptions - Perform word wrap with user options
func WrapWithOptions(in string, options *WrapOptions) string {
	fields := strings.Fields(in)
	if options.Behavior == WrapLiteral {
		return wrapLiteral(strings.Join(fields, ` `), options.Width, options.LineBreak)
	} else if options.Behavior == WrapAfterWord {
		return wrapAfterWord(fields, options.Width, options.LineBreak)
	}

	return wrapBeforeWord(fields, options.Width, options.LineBreak)
}

func wrapLiteral(in string, width uint, lineBreak string) string {
	lines := uint(math.Ceil(float64(len(in)) / float64(width)))
	var out string
	var i uint
	for i = 0; i < lines; i++ {
		start := i * width
		remain := in[start:]
		if len(remain) <= int(width) {
			out += remain
			continue
		}
		out += in[start:start+width] + lineBreak
	}

	return out
}

func wrapBeforeWord(words []string, width uint, lineBreak string) string {
	var out, line string
	for _, word := range words {
		_line := line + word
		if len(_line) > int(width) {
			out += strings.TrimSpace(line) + lineBreak
			line = word + " "
			continue
		}
		line += word + " "
	}
	out += strings.TrimSpace(line)
	return out
}

func wrapAfterWord(words []string, width uint, lineBreak string) string {
	var out, line string
	for _, word := range words {
		if len(line)+1 > int(width) {
			out += strings.TrimSpace(line) + lineBreak
			line = word + " "
			continue
		}
		line += word + " "
	}
	out += strings.TrimSpace(line)
	return out
}
