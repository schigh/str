package str

import "strings"

// PipeMap returns a transformer that splits the input using split, applies fn
// to each part, and rejoins with the join string. Nil split or nil fn returns
// the input unchanged.
func PipeMap(split func(string) []string, join string, fn func(string) string) func(string) string {
	return func(s string) string {
		if split == nil || fn == nil {
			return s
		}
		parts := split(s)
		if len(parts) == 0 {
			return ""
		}
		for i, p := range parts {
			parts[i] = fn(p)
		}
		return strings.Join(parts, join)
	}
}

// Lines splits a string on newlines. Empty string returns []string{""}.
// Splits on "\n" only; "\r" is preserved if present.
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// Words splits a string on whitespace, matching strings.Fields behavior.
// Empty string returns nil.
func Words(s string) []string {
	return strings.Fields(s)
}
