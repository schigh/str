package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func strpadtests() []string {
	return []string{
		"0",
		"",
		"1234567890",
		"ababababababab",
		"aaa",
	}
}

func ExamplePad() {
	fmt.Println(Pad("123", "0", 6, PadLeft))
	// Output: 000123
}

func TestPadLeftSingle(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"000000",
		"000000",
		"1234567890",
		"ababababababab",
		"000aaa",
	}
	for idx, test := range tests {
		pad := Pad(test, "0", 6, PadLeft)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestPadLeftMulti(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"xoxoxoxoxoxo0",
		"xoxoxoxoxoxo",
		"xo1234567890",
		"ababababababab",
		"xoxoxoxoxoaaa",
	}
	for idx, test := range tests {
		pad := Pad(test, "xo", 12, PadLeft)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestPadRightSingle(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"000000",
		"000000",
		"1234567890",
		"ababababababab",
		"aaa000",
	}
	for idx, test := range tests {
		pad := Pad(test, "0", 6, PadRight)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestPadRightMulti(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"0xoxoxoxoxoxo",
		"xoxoxoxoxoxo",
		"1234567890xo",
		"ababababababab",
		"aaaxoxoxoxoxo",
	}
	for idx, test := range tests {
		pad := Pad(test, "xo", 12, PadRight)
		assert.Equal(t, pad, correct[idx])
	}
}
