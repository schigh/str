package str

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 - String hash function.  This function wraps the underlying
// crypto/sha1 functionality for exclusive use with strings.
// Exercise caution when using SHA1, as it is considered not to be
// cryptographically secure.
// https://en.wikipedia.org/wiki/SHA-1
func SHA1(in string) string {
	sum := sha1.Sum([]byte(in))
	sha1Bytes := sum[:]
	bytesOut := make([]byte, sha1.Size*2)
	hex.Encode(bytesOut, sha1Bytes)

	return string(bytesOut)
}
