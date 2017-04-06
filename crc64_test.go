package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestCRC64TestString = "This is a string of tests. 世界"
	CRC64ResultECMA     = "12939983120264787714"
	CRC64ResultISO      = "11395191879162951138"
)

func TestCRC64(t *testing.T) {
	assert.Equal(t, CRC64ResultECMA, CRC64(TestCRC64TestString))
}

func TestCRC64_ISO(t *testing.T) {
	options := &CRC64Options{
		PolynomialType: CRC64TypeISO,
	}
	SetCRC64Options(options)
	assert.Equal(t, CRC64ResultISO, CRC64(TestCRC64TestString))
}
