package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	assert.Equal(t, "1d340856fb75496176c4b968925452e7", MD5("This is a string of tests. 世界"))
}
