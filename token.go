package str

import (
	"crypto/rand"
	"math/big"
)

// TokenOption configures RandomToken behavior.
type TokenOption func(*tokenConfig)

type tokenConfig struct {
	length  int
	charset string
	prefix  string
}

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// WithLength sets the number of random characters (default 20).
// Prefix is not included in this count.
func WithLength(n int) TokenOption {
	return func(c *tokenConfig) { c.length = n }
}

// WithCharset sets the character set to sample from (default alphanumeric).
// Duplicate characters are deduplicated.
func WithCharset(cs string) TokenOption {
	return func(c *tokenConfig) { c.charset = cs }
}

// WithPrefix sets a prefix prepended to the random characters.
func WithPrefix(p string) TokenOption {
	return func(c *tokenConfig) { c.prefix = p }
}

// RandomToken generates a cryptographically random string.
// Uses crypto/rand for secure randomness. Panics if the OS entropy source fails.
func RandomToken(opts ...TokenOption) string {
	cfg := &tokenConfig{
		length:  20,
		charset: defaultCharset,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.charset == "" {
		cfg.charset = defaultCharset
	}
	cs := dedup(cfg.charset)
	if cfg.length <= 0 {
		return cfg.prefix
	}

	buf := make([]byte, 0, cfg.length)
	max := big.NewInt(int64(len(cs)))
	for range cfg.length {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic("str: crypto/rand failed: " + err.Error())
		}
		buf = append(buf, cs[n.Int64()])
	}
	return cfg.prefix + string(buf)
}

func dedup(s string) []byte {
	seen := make(map[byte]bool, len(s))
	out := make([]byte, 0, len(s))
	for i := range len(s) {
		b := s[i]
		if !seen[b] {
			seen[b] = true
			out = append(out, b)
		}
	}
	return out
}
