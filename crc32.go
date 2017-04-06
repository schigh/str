package str

import (
	"fmt"
	"hash/crc32"
)

// Wrapper for crc32 polynomial options
type CRC32Options struct {
	PolynomialType uint32
}

const (
	CRC32TypeDefault uint32 = 0xEDB88320
	CRC32TypeC       uint32 = 0x82F63B78
	CRC32TypeK       uint32 = 0xEB31D82E
	CRC32TypeK2      uint32 = 0x992C1A4C
	CRC32TypeQ       uint32 = 0xD5828281
)

var defaultCRC32Options *CRC32Options

func init() {
	defaultCRC32Options = &CRC32Options{
		PolynomialType: CRC32TypeDefault,
	}
}

// SetCRC32Options - set the default crc32 options for the application
func SetCRC32Options(options *CRC32Options) {
	defaultCRC32Options = options
}

// CRC32 - perform the cyclic redundancy checksum with default options
func CRC32(in string) string {
	return CRC32WithOptions(in, defaultCRC32Options)
}

// CRC32WithOptions - perform the cyclic redundancy checksum with supplied options
func CRC32WithOptions(in string, options *CRC32Options) string {
	table := crc32.MakeTable(options.PolynomialType)
	return fmt.Sprintf(`%d`, crc32.Checksum([]byte(in), table))
}
