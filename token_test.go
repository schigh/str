package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	opts := &TokenOptions{
		Length:  10,
		Charset: "abcdef",
		Prefix:  "test_",
	}
	SetTokenOptions(opts)
	token := Token()
	assert.Regexp(t, "^test_[a-f]{5}$", token, nil)
}
