package str

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(in string) string {
	sum := md5.Sum([]byte(in))
	md5Bytes := sum[:]
	bytesOut := make([]byte, md5.Size*2)
	hex.Encode(bytesOut, md5Bytes)

	return string(bytesOut)
}
