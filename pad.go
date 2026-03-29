package str

import "unicode/utf8"

// PadLeft returns a function that pads the input string on the left with the
// given pad string until it reaches the specified width (in runes).
// Multi-character pad strings are truncated on the final repetition to hit
// exactly width. Empty pad string or width <= input length returns input unchanged.
func PadLeft(padStr string, width int) func(string) string {
	return func(s string) string {
		return pad(s, padStr, width, true)
	}
}

// PadRight returns a function that pads the input string on the right with the
// given pad string until it reaches the specified width (in runes).
func PadRight(padStr string, width int) func(string) string {
	return func(s string) string {
		return pad(s, padStr, width, false)
	}
}

func pad(s, padStr string, width int, left bool) string {
	if padStr == "" || width <= 0 {
		return s
	}
	inputLen := utf8.RuneCountInString(s)
	if inputLen >= width {
		return s
	}
	needed := width - inputLen
	padRunes := []rune(padStr)
	buf := make([]rune, 0, needed)
	for len(buf) < needed {
		buf = append(buf, padRunes...)
	}
	buf = buf[:needed]
	if left {
		return string(buf) + s
	}
	return s + string(buf)
}
