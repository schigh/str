package str

import (
	"strings"
	"unicode"
)

// Reverse returns the input string with runes in reverse order.
// This is rune-aware but not grapheme-cluster-aware; combining characters
// may reorder.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CollapseWhitespace collapses all runs of Unicode whitespace (spaces, tabs,
// newlines, vertical tabs, etc.) to a single ASCII space. Uses unicode.IsSpace
// for classification. Leading and trailing whitespace is also collapsed to a
// single space if present, then trimmed.
func CollapseWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	inSpace := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			if !inSpace {
				b.WriteByte(' ')
				inSpace = true
			}
		} else {
			b.WriteRune(r)
			inSpace = false
		}
	}
	return strings.TrimSpace(b.String())
}

// TrimSpace removes leading and trailing whitespace from the input string.
// This is a re-export of strings.TrimSpace for pipeline convenience.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// SlugifyASCII generates a URL-safe ASCII slug from the input string.
// Lowercases, replaces spaces and punctuation with hyphens, collapses
// consecutive hyphens, and strips all non-ASCII characters.
func SlugifyASCII(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	prevHyphen := false
	for _, r := range s {
		if r > 127 {
			continue
		}
		action := classifySlugRune(r)
		switch action {
		case slugKeep:
			b.WriteRune(r)
			prevHyphen = false
		case slugLower:
			b.WriteRune(r + 32)
			prevHyphen = false
		case slugHyphen:
			if !prevHyphen && b.Len() > 0 {
				b.WriteByte('-')
				prevHyphen = true
			}
		}
	}
	result := b.String()
	if len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}
	return result
}

type slugAction int

const (
	slugSkip   slugAction = iota
	slugKeep              // lowercase letter or digit
	slugLower             // uppercase letter to lowercase
	slugHyphen            // replace with hyphen
)

func classifySlugRune(r rune) slugAction {
	switch {
	case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
		return slugKeep
	case r >= 'A' && r <= 'Z':
		return slugLower
	case r == ' ' || r == '-' || r == '_' || unicode.IsPunct(r) || unicode.IsSymbol(r):
		return slugHyphen
	default:
		return slugSkip
	}
}
