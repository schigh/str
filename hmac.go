package str

import (
	"net"
	"regexp"
)

// HMACOptions - A wrapper for HMAC key data
type HMACOptions struct {
	KeyData []byte
}

var defaultHMACOptions *HMACOptions

func init() {
	// This just takes the hardware address of the first network interface
	// matching the pattern (ex: 00:af:19:3c:2e:80)
	// and uses the SHA256 hash of that as the default key.  It is NOT
	// a secure implementation, as it is recommended that the user
	// seed the str implementation with their own key data
	re := regexp.MustCompile(`(([a-f0-9]{2}:){5}[a-f0-9]{2}$)`)
	machineName := ""
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
		KeyData: []byte(machineName),
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
	return ""
}
