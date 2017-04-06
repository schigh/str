package str

import (
	"fmt"
	"hash/crc64"
)

// Wrapper for crc32 polynomial options
type CRC64Options struct {
	PolynomialType uint64
}

const (
	CRC64TypeECMA uint64 = crc64.ECMA
	CRC64TypeISO  uint64 = crc64.ISO
)

var defaultCRC64Options *CRC64Options

func init() {
	defaultCRC64Options = &CRC64Options{
		PolynomialType: CRC64TypeECMA,
	}
}

// SetCRC64Options - set the default crc64 options for the application
func SetCRC64Options(options *CRC64Options) {
	defaultCRC64Options = options
}

// CRC64 - perform the cyclic redundancy checksum with default options
func CRC64(in string) string {
	return CRC64WithOptions(in, defaultCRC64Options)
}

// CRC64WithOptions - perform the cyclic redundancy checksum with supplied options
func CRC64WithOptions(in string, options *CRC64Options) string {
	table := crc64.MakeTable(options.PolynomialType)
	return fmt.Sprintf(`%d`, crc64.Checksum([]byte(in), table))
}
