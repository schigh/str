package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestHMACTestString = "This is a string of tests. 世界"
	HMACResultMD5      = "63014a2296a106d0ca704dc7976f7e3e"
	HMACResultSHA1     = "a93d9e728669c00069b759feee8602a4105c20c9"
	HMACResultSHA256   = "e93a940a97065132c6a7192f803940ba52f166292813c58644627e6d0ea7db04"
	HMACResultsSHA384  = "ad0556957fab2be8777dd0da60b8f509e4dc7291704eb844b0879cec84fc1794a8da1316c5a5f4f131efb20d55393ae1"
	HMACResultsSHA512  = "7a971944106120a110466d20d55c37dad6fbe88ef5e97e40bfa537744435250fe0b0f80f90d3cb8cd1d50f5b5de63118eacf77b506eed27ff90e0ec9febf667b"
)

var TestHMACKeyData []byte

func init() {
	TestHMACKeyData = []byte(`This is a random string that is not meant to represent anything secure`)
}

func TestHMAC(t *testing.T) {
	options := &HMACOptions{
		KeyData:    TestHMACKeyData,
		DigestType: HMACDigestTypeSHA256,
	}
	SetHMACOptions(options)
	assert.Equal(t, HMACResultSHA256, HMAC(TestHMACTestString))
}

func TestHMACMD5(t *testing.T) {
	options := &HMACOptions{
		KeyData:    TestHMACKeyData,
		DigestType: HMACDigestTypeMD5,
	}
	SetHMACOptions(options)
	assert.Equal(t, HMACResultMD5, HMAC(TestHMACTestString))
}

func TestHMACSHA1(t *testing.T) {
	options := &HMACOptions{
		KeyData:    TestHMACKeyData,
		DigestType: HMACDigestTypeSHA1,
	}
	SetHMACOptions(options)
	assert.Equal(t, HMACResultSHA1, HMAC(TestHMACTestString))
}

func TestHMACSHA384(t *testing.T) {
	options := &HMACOptions{
		KeyData:    TestHMACKeyData,
		DigestType: HMACDigestTypeSHA384,
	}
	SetHMACOptions(options)
	assert.Equal(t, HMACResultsSHA384, HMAC(TestHMACTestString))
}

func TestHMACSHA512(t *testing.T) {
	options := &HMACOptions{
		KeyData:    TestHMACKeyData,
		DigestType: HMACDigestTypeSHA512,
	}
	SetHMACOptions(options)
	assert.Equal(t, HMACResultsSHA512, HMAC(TestHMACTestString))
}
