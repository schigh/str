package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	assert.Equal(t, "d5013a8864bd274b38a5207a018ef15aa8d09f12abb33162182740d89ed47bb6", SHA256("This is a string of tests. 世界"))
}

func ExampleSHA256() {
	fmt.Println(SHA256("This is a string of tests. 世界"))
	// Output: d5013a8864bd274b38a5207a018ef15aa8d09f12abb33162182740d89ed47bb6
}
