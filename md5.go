package str

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 - String hash function.  This function wraps the underlying
// crypto/md5 functionality for exclusive use with strings.
// Exercise caution when using MD5, as it is considered not to be
// cryptographically secure.
// https://en.wikipedia.org/wiki/MD5
func MD5(in string) string {
	sum := md5.Sum([]byte(in))
	md5Bytes := sum[:]
	bytesOut := make([]byte, md5.Size*2)
	hex.Encode(bytesOut, md5Bytes)

	return string(bytesOut)
}
