package str

import (
	"math"
	"strings"
)

type StrPadding int

const (
	StrPadLeft StrPadding = iota
	StrPadRight
)

// StrPad - left-pads or right-pads a string to a specified length
// using a variable-length pad.  This function will always pad until the
// length is reached, and in the case of pads with multiple characters,
// will overshoot the length if necessary.
func StrPad(in string, pad string, length int, direction StrPadding) string {
	inLen := len(in)
	if inLen >= length {
		return in
	}

	numPads := int(math.Ceil(float64(length-inLen) / float64(len(pad))))
	switch direction {
	case StrPadLeft:
		return strings.Repeat(pad, numPads) + in
	case StrPadRight:
		return in + strings.Repeat(pad, numPads)
	default:
		return in
	}
}

// SubStr - gets a substring of a string based on a start index and a length.
// If the start index is negative, the substring will start [index] number
// of characters from the tail of the string
// If the requested length exceeds the adjusted slice length, then the tail
// of the string starting at [start] is returned.
// If [start] is out of range from either front or back, then an empty
// string is returned
func SubStr(in string, start int, length uint) string {
	strLen := len(in)
	if length == 0 {
		return ""
	}
	if start < 0 {
		frontIndex := strLen + start
		if frontIndex < 0 {
			return ""
		}
		start = frontIndex
	}

	if start >= strLen {
		return ""
	}

	rearIndex := start + int(length)

	if rearIndex >= strLen {
		return in[start:]
	}

	return in[start:rearIndex]
}
