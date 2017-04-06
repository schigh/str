package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestCRC32TestString = "This is a string of tests. 世界"
	CRC32ResultDefault  = "3425437894"
	CRC32ResultC        = "2760337661"
	CRC32ResultK        = "1921250769"
	CRC32ResultK2       = "1662776799"
	CRC32ResultQ        = "3464169987"
)

func TestCRC32(t *testing.T) {
	assert.Equal(t, CRC32ResultDefault, CRC32(TestCRC32TestString))
}

func TestCRC322_C(t *testing.T) {
	options := &CRC32Options{
		PolynomialType: CRC32TypeC,
	}
	SetCRC32Options(options)
	assert.Equal(t, CRC32ResultC, CRC32(TestCRC32TestString))
}

func TestCRC322_K(t *testing.T) {
	options := &CRC32Options{
		PolynomialType: CRC32TypeK,
	}
	SetCRC32Options(options)
	assert.Equal(t, CRC32ResultK, CRC32(TestCRC32TestString))
}

func TestCRC322_K2(t *testing.T) {
	options := &CRC32Options{
		PolynomialType: CRC32TypeK2,
	}
	SetCRC32Options(options)
	assert.Equal(t, CRC32ResultK2, CRC32(TestCRC32TestString))
}

func TestCRC322_Q(t *testing.T) {
	options := &CRC32Options{
		PolynomialType: CRC32TypeQ,
	}
	SetCRC32Options(options)
	assert.Equal(t, CRC32ResultQ, CRC32(TestCRC32TestString))
}
