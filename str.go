package str

import (
	"math"
	"strings"
)

type Padding int

const (
	PadLeft  Padding = iota
	PadRight
)

// Padding - left-pads or right-pads a string to a specified length
// using a variable-length pad.  This function will always pad until the
// length is reached, and in the case of pads with multiple characters,
// will overshoot the length if necessary.
func Pad(in string, pad string, length int, direction Padding) string {
	inLen := len(in)
	if inLen >= length {
		return in
	}

	numPads := int(math.Ceil(float64(length-inLen) / float64(len(pad))))
	switch direction {
	case PadLeft:
		return strings.Repeat(pad, numPads) + in
	case PadRight:
		return in + strings.Repeat(pad, numPads)
	default:
		return in
	}
}

// Substring - gets a substring of a string based on a start index and a length.
// If the start index is negative, the substring will start [index] number
// of characters from the tail of the string
// If the requested length exceeds the adjusted slice length, then the tail
// of the string starting at [start] is returned.
// If [start] is out of range from either front or back, then an empty
// string is returned
func Substring(in string, start int, length uint) string {
	size := len(in)
	if length == 0 {
		return ""
	}
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
