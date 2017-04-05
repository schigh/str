package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA1(t *testing.T) {
	assert.Equal(t, "749b48fa078203b28acb7ee8d581d63efa23a3cf", SHA1("This is a string of tests. 世界"))
}

func ExampleSHA1() {
	fmt.Println(SHA1("This is a string of tests. 世界"))
	// Output: 749b48fa078203b28acb7ee8d581d63efa23a3cf
}
