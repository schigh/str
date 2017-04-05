package str

import (
	"math"
	"strings"
)

type Padding int

const (
	PadLeft Padding = iota
	PadRight
)

// Pad - left-pads or right-pads a string to a specified length
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
