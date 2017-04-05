package str

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 - Cryptographic string hash function.  This function wraps the underlying
// crypto/sha256 functionality for exclusive use with strings.
// https://en.wikipedia.org/wiki/SHA-2
func SHA256(in string) string {
	sum := sha256.Sum256([]byte(in))
	sha256Bytes := sum[:]
	bytesOut := make([]byte, sha256.Size*2)
	hex.Encode(bytesOut, sha256Bytes)

	return string(bytesOut)
}
