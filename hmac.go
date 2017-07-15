package str

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"net"
	"regexp"
)

// HMACDigestType - The hash type for the HMAC function
type HMACDigestType int

const (
	HMACDigestTypeSHA256 HMACDigestType = iota // use sha256 digest
	HMACDigestTypeMD5                          // use md5 digest
	HMACDigestTypeSHA1                         // use sha1 digest
	HMACDigestTypeSHA384                       // use sha384 digest
	HMACDigestTypeSHA512                       // use sha512 digest
)

// HMACOptions - A wrapper for HMAC key data
type HMACOptions struct {
	KeyData    []byte         // the raw key bytes
	DigestType HMACDigestType // digest algorithm used
}

var defaultHMACOptions *HMACOptions // the default HMAC options

func init() {
	// This just takes the hardware address of the first network interface
	// matching the pattern (ex: 00:af:19:3c:2e:80)
	// and uses the SHA256 hash of that as the default key.  It is NOT
	// a secure implementation, and it is recommended that the user
	// seed the str implementation with their own key data
	re := regexp.MustCompile(`(([a-f0-9]{2}:){5}[a-f0-9]{2}$)`)
	var machineName string
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, inter := range interfaces {
			_s := inter.HardwareAddr.String()
			if re.MatchString(_s) {
				machineName = SHA256(_s)
				break
			}
		}
	}

	defaultHMACOptions = &HMACOptions{
		KeyData:    []byte(machineName),
		DigestType: HMACDigestTypeSHA256,
	}
}

// SetHMACOptions - Set the default HMAC options
func SetHMACOptions(options *HMACOptions) {
	defaultHMACOptions = options
}

// HMAC - get the HMAC hash of a string.  This uses the
// default options, which can be set during a call to
// SetHMACOptions(...)
func HMAC(in string) string {
	return HMACWithOptions(in, defaultHMACOptions)
}

// HMACWithOptions - get the HMAC hash of a string
// using user-supplied options
func HMACWithOptions(in string, options *HMACOptions) string {
	// digest function
	var hasher func() hash.Hash
	switch options.DigestType {
	case HMACDigestTypeMD5:
		hasher = md5.New
	case HMACDigestTypeSHA1:
		hasher = sha1.New
	case HMACDigestTypeSHA256:
		hasher = sha256.New
	case HMACDigestTypeSHA384:
		hasher = sha512.New384
	case HMACDigestTypeSHA512:
		hasher = sha512.New
	default:
		return ""
	}

	mac := hmac.New(hasher, options.KeyData)
	mac.Write([]byte(in))
	sum := mac.Sum(nil)
	bytesOut := make([]byte, getHexBufferSize(options.DigestType))
	hex.Encode(bytesOut, sum)

	return string(bytesOut)
}

// getHexBufferSize retrieves the buffer size
// for hex encoding the provided digest type
func getHexBufferSize(kind HMACDigestType) int {
	switch kind {
	case HMACDigestTypeMD5:
		return md5.Size * 2
	case HMACDigestTypeSHA1:
		return sha1.Size * 2
	case HMACDigestTypeSHA256:
		return sha256.Size * 2
	case HMACDigestTypeSHA384:
		return sha512.Size384 * 2
	case HMACDigestTypeSHA512:
		return sha512.Size * 2
	}

	return 0
}
