package str

// Substring returns a function that extracts a substring from the input.
// Start and length are in runes. Negative start counts from the end (Python-style).
// Negative length returns empty string. Out-of-bounds indices clamp to string boundaries.
func Substring(start, length int) func(string) string {
	return func(s string) string {
		if length <= 0 {
			return ""
		}
		runes := []rune(s)
		n := len(runes)
		if n == 0 {
			return ""
		}

		idx := start
		if idx < 0 {
			idx += n
		}
		if idx < 0 {
			idx = 0
		}
		if idx >= n {
			return ""
		}

		end := min(idx+length, n)
		return string(runes[idx:end])
	}
}
