package str

// Truncate returns a function that truncates the input string to maxLen runes.
// If the input is longer than maxLen, it is truncated and the ellipsis is appended.
// maxLen includes the ellipsis length. Counts runes, not bytes.
func Truncate(maxLen int, ellipsis string) func(string) string {
	return func(s string) string {
		if maxLen <= 0 {
			return ""
		}
		runes := []rune(s)
		if len(runes) <= maxLen {
			return s
		}
		ellRunes := []rune(ellipsis)
		if len(ellRunes) >= maxLen {
			return string(ellRunes[:maxLen])
		}
		return string(runes[:maxLen-len(ellRunes)]) + ellipsis
	}
}
