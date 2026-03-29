package str

import (
	"strings"
	"unicode"
)

const naturalSortPadWidth = 20

// NaturalSortKey returns a string key that sorts naturally when compared
// lexicographically. Numeric segments (ASCII digits) are zero-padded to 20
// digits. Case is folded to lowercase. "file2" sorts before "file10".
func NaturalSortKey(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(s) * 2)
	i := 0
	runes := []rune(strings.ToLower(s))
	for i < len(runes) {
		r := runes[i]
		if r >= '0' && r <= '9' {
			// Collect the full digit run.
			start := i
			for i < len(runes) && runes[i] >= '0' && runes[i] <= '9' {
				i++
			}
			digits := string(runes[start:i])
			if len(digits) < naturalSortPadWidth {
				b.WriteString(strings.Repeat("0", naturalSortPadWidth-len(digits)))
			}
			b.WriteString(digits)
		} else {
			b.WriteRune(unicode.ToLower(r))
			i++
		}
	}
	return b.String()
}
