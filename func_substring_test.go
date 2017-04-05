package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSubstring() {
	fmt.Println(Substring("This is a sentence", 5, 2))
	// Output: is
}

func TestSubstring(t *testing.T) {
	testString := "This is a test string that you should use for tests"

	// forward in bounds
	s1 := Substring(testString, 0, 4)
	assert.Equal(t, "This", s1)

	// backward in bounds
	s2 := Substring(testString, -5, 4)
	assert.Equal(t, "test", s2)

	// forward in bounds
	s3 := Substring(testString, 10, 11)
	assert.Equal(t, "test string", s3)

	// backward in bounds
	s4 := Substring(testString, -13, 3)
	assert.Equal(t, "use", s4)

	// forward with too large length
	s5 := Substring(testString, 38, 100)
	assert.Equal(t, "use for tests", s5)

	// backward with too large length
	s6 := Substring(testString, -13, 100)
	assert.Equal(t, "use for tests", s6)

	// forward with invalid start
	s7 := Substring(testString, 60, 1)
	assert.Equal(t, "", s7)

	// backward with invalid start
	s8 := Substring(testString, -60, 1)
	assert.Equal(t, "", s8)

	// zero length
	s9 := Substring(testString, 0, 0)
	assert.Equal(t, "", s9)
}
