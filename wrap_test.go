package str

import "testing"

func TestWrap(t *testing.T) {
	long := "This is a string that could very likely be broken up into more than one line"

	t.Run("before word default", func(t *testing.T) {
		got := Wrap(24)(long)
		// Each line should be <= 24 runes (words aren't split).
		if got == long {
			t.Error("expected wrapping")
		}
		if got == "" {
			t.Error("got empty string")
		}
	})

	t.Run("after word", func(t *testing.T) {
		got := Wrap(24, WrapAfterWord)(long)
		if got == long {
			t.Error("expected wrapping")
		}
	})

	t.Run("hard break", func(t *testing.T) {
		got := Wrap(24, WrapHardBreak)("abcdefghijklmnopqrstuvwxyz1234567890")
		expected := "abcdefghijklmnopqrstuvwx\nyz1234567890"
		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})

	t.Run("custom line break", func(t *testing.T) {
		got := Wrap(10, WrapHardBreak, WithLineBreak("<br>"))("abcdefghijklmnop")
		expected := "abcdefghij<br>klmnop"
		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})

	t.Run("with indent", func(t *testing.T) {
		got := Wrap(20, WithIndent("  "))("hello world this is a test")
		if got == "" {
			t.Error("got empty string")
		}
	})

	t.Run("width zero", func(t *testing.T) {
		got := Wrap(0)("hello")
		if got != "hello" {
			t.Errorf("got %q, want %q", got, "hello")
		}
	})

	t.Run("empty input", func(t *testing.T) {
		got := Wrap(10)("")
		if got != "" {
			t.Errorf("got %q, want empty", got)
		}
	})
}
