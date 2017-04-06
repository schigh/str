package str

// Substring - gets a substring of a string based on a start index and a length.
// If the start index is negative, the substring will start [index] number
// of characters from the tail of the string
// If the requested length exceeds the adjusted slice length, then the tail
// of the string starting at [start] is returned.
// If [start] is out of range from either front or back, then an empty
// string is returned
func Substring(in string, start int, length uint) string {
	if length == 0 {
		return ""
	}
	size := len(in)
	if start < 0 {
		frontIndex := size + start
		if frontIndex < 0 {
			return ""
		}
		start = frontIndex
	}

	if start >= size {
		return ""
	}

	rearIndex := start + int(length)

	if rearIndex >= size {
		return in[start:]
	}

	return in[start:rearIndex]
}
