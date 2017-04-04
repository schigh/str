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

func ExampleStrPad() {
	fmt.Println(StrPad("123", "0", 6, StrPadLeft))
	// Output: 000123
}

func ExampleSubStr() {
	fmt.Println(SubStr("This is a sentence", 5, 2))
	// Output: is
}

func TestStrPadLeftSingle(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"000000",
		"000000",
		"1234567890",
		"ababababababab",
		"000aaa",
	}
	for idx, test := range tests {
		pad := StrPad(test, "0", 6, StrPadLeft)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestStrPadLeftMulti(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"xoxoxoxoxoxo0",
		"xoxoxoxoxoxo",
		"xo1234567890",
		"ababababababab",
		"xoxoxoxoxoaaa",
	}
	for idx, test := range tests {
		pad := StrPad(test, "xo", 12, StrPadLeft)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestStrPadRightSingle(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"000000",
		"000000",
		"1234567890",
		"ababababababab",
		"aaa000",
	}
	for idx, test := range tests {
		pad := StrPad(test, "0", 6, StrPadRight)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestStrPadRightMulti(t *testing.T) {
	tests := strpadtests()
	correct := []string{
		"0xoxoxoxoxoxo",
		"xoxoxoxoxoxo",
		"1234567890xo",
		"ababababababab",
		"aaaxoxoxoxoxo",
	}
	for idx, test := range tests {
		pad := StrPad(test, "xo", 12, StrPadRight)
		assert.Equal(t, pad, correct[idx])
	}
}

func TestSubStr(t *testing.T) {
	testString := "This is a test string that you should use for tests"

	// forward in bounds
	str1 := SubStr(testString, 0, 4)
	assert.Equal(t, "This", str1)

	// backward in bounds
	str2 := SubStr(testString, -5, 4)
	assert.Equal(t, "test", str2)

	// forward in bounds
	str3 := SubStr(testString, 10, 11)
	assert.Equal(t, "test string", str3)

	// backward in bounds
	str4 := SubStr(testString, -13, 3)
	assert.Equal(t, "use", str4)

	// forward with too large length
	str5 := SubStr(testString, 38, 100)
	assert.Equal(t, "use for tests", str5)

	// backward with too large length
	str6 := SubStr(testString, -13, 100)
	assert.Equal(t, "use for tests", str6)

	// forward with invalid start
	str7 := SubStr(testString, 60, 1)
	assert.Equal(t, "", str7)

	// backward with invalid start
	str8 := SubStr(testString, -60, 1)
	assert.Equal(t, "", str8)

	// zero length
	str9 := SubStr(testString, 0, 0)
	assert.Equal(t, "", str9)
}
