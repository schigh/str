package str

import (
	"math/rand"
	"time"
)

type TokenOptions struct {
	Length  uint
	Charset string
	Prefix  string
}

var defaultTokenOptions *TokenOptions

func init() {
	defaultTokenOptions = &TokenOptions{
		Length:  20,
		Charset: "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789",
	}
	rand.Seed(time.Now().UnixNano())
}

// SetTokenOptions - set the options for token generation
func SetTokenOptions(options *TokenOptions) {
	defaultTokenOptions = options
}

// Token - create pseudorandom token
func Token() string {
	return TokenWithOptions(defaultTokenOptions)
}

// TokenWithOptions - create pseudorandom token with user-supplied options
func TokenWithOptions(options *TokenOptions) string {
	cl := len(options.Charset)
	pl := len(options.Prefix)
	if options.Length < 1 || cl == 0 {
		return ""
	}
	if pl >= int(options.Length) {
		return options.Prefix
	}
	rounds := int(options.Length) - pl
	var token string
	for i := 0; i < rounds; i++ {
		guard := rand.Intn(cl)
		token += options.Charset[guard : guard+1]
	}

	return options.Prefix + token
}
