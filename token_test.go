package str

import (
	"strings"
	"testing"
)

func TestRandomToken(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		tok := RandomToken()
		if len(tok) != 20 {
			t.Errorf("got length %d, want 20", len(tok))
		}
		for _, c := range tok {
			if !strings.ContainsRune(defaultCharset, c) {
				t.Errorf("char %c not in default charset", c)
			}
		}
	})

	t.Run("custom length", func(t *testing.T) {
		tok := RandomToken(WithLength(32))
		if len(tok) != 32 {
			t.Errorf("got length %d, want 32", len(tok))
		}
	})

	t.Run("length zero", func(t *testing.T) {
		tok := RandomToken(WithLength(0))
		if tok != "" {
			t.Errorf("got %q, want empty", tok)
		}
	})

	t.Run("binary charset", func(t *testing.T) {
		tok := RandomToken(WithLength(10), WithCharset("01"))
		if len(tok) != 10 {
			t.Errorf("got length %d, want 10", len(tok))
		}
		for _, c := range tok {
			if c != '0' && c != '1' {
				t.Errorf("got char %c, want 0 or 1", c)
			}
		}
	})

	t.Run("empty charset uses default", func(t *testing.T) {
		tok := RandomToken(WithLength(5), WithCharset(""))
		if len(tok) != 5 {
			t.Errorf("got length %d, want 5", len(tok))
		}
	})

	t.Run("charset with duplicates", func(t *testing.T) {
		tok := RandomToken(WithLength(100), WithCharset("aaab"))
		for _, c := range tok {
			if c != 'a' && c != 'b' {
				t.Errorf("got char %c, want a or b", c)
			}
		}
	})

	t.Run("with prefix", func(t *testing.T) {
		tok := RandomToken(WithPrefix("tok_"), WithLength(6))
		if !strings.HasPrefix(tok, "tok_") {
			t.Errorf("got %q, want prefix tok_", tok)
		}
		if len(tok) != 10 { // "tok_" (4) + 6
			t.Errorf("got length %d, want 10", len(tok))
		}
	})

	t.Run("prefix with zero length", func(t *testing.T) {
		tok := RandomToken(WithPrefix("tok_"), WithLength(0))
		if tok != "tok_" {
			t.Errorf("got %q, want %q", tok, "tok_")
		}
	})

	t.Run("uniqueness", func(t *testing.T) {
		a := RandomToken()
		b := RandomToken()
		if a == b {
			t.Error("two tokens should not be identical")
		}
	})
}
