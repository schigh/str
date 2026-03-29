package str

import (
	"strings"
	"testing"
)

func TestPipeMap(t *testing.T) {
	t.Run("trim every line", func(t *testing.T) {
		fn := PipeMap(Lines, "\n", TrimSpace)
		got := fn("  hello  \n  world  ")
		if got != "hello\nworld" {
			t.Errorf("got %q, want %q", got, "hello\nworld")
		}
	})

	t.Run("uppercase every word", func(t *testing.T) {
		fn := PipeMap(Words, " ", strings.ToUpper)
		got := fn("hello world")
		if got != "HELLO WORLD" {
			t.Errorf("got %q, want %q", got, "HELLO WORLD")
		}
	})

	t.Run("nil split is passthrough", func(t *testing.T) {
		fn := PipeMap(nil, "\n", strings.ToUpper)
		if got := fn("hello"); got != "hello" {
			t.Errorf("got %q, want %q", got, "hello")
		}
	})

	t.Run("nil fn is passthrough", func(t *testing.T) {
		fn := PipeMap(Lines, "\n", nil)
		if got := fn("hello"); got != "hello" {
			t.Errorf("got %q, want %q", got, "hello")
		}
	})

	t.Run("empty string", func(t *testing.T) {
		fn := PipeMap(Lines, "\n", TrimSpace)
		if got := fn(""); got != "" {
			t.Errorf("got %q, want empty", got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		fn := PipeMap(Lines, "\n", strings.ToUpper)
		if got := fn("hello"); got != "HELLO" {
			t.Errorf("got %q, want %q", got, "HELLO")
		}
	})

	t.Run("composes in Pipe", func(t *testing.T) {
		fn := Pipe(
			TrimSpace,
			PipeMap(Lines, "\n", TrimSpace),
		)
		got := fn("  hello  \n  world  ")
		if got != "hello\nworld" {
			t.Errorf("got %q, want %q", got, "hello\nworld")
		}
	})
}

func TestLines(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		got := Lines("a\nb\nc")
		if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
			t.Errorf("got %v", got)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		got := Lines("")
		if len(got) != 1 || got[0] != "" {
			t.Errorf("got %v, want [\"\"]", got)
		}
	})

	t.Run("no newlines", func(t *testing.T) {
		got := Lines("hello")
		if len(got) != 1 || got[0] != "hello" {
			t.Errorf("got %v", got)
		}
	})

	t.Run("trailing newline", func(t *testing.T) {
		got := Lines("a\nb\n")
		if len(got) != 3 || got[2] != "" {
			t.Errorf("got %v, want [a, b, \"\"]", got)
		}
	})

	t.Run("cr preserved", func(t *testing.T) {
		got := Lines("a\r\nb")
		if len(got) != 2 || got[0] != "a\r" {
			t.Errorf("got %v, want [\"a\\r\", \"b\"]", got)
		}
	})
}

func TestWords(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		got := Words("hello world")
		if len(got) != 2 || got[0] != "hello" || got[1] != "world" {
			t.Errorf("got %v", got)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		got := Words("")
		if len(got) != 0 {
			t.Errorf("got %v, want empty", got)
		}
	})

	t.Run("multiple spaces", func(t *testing.T) {
		got := Words("hello   world")
		if len(got) != 2 {
			t.Errorf("got %v, want 2 elements", got)
		}
	})

	t.Run("single word", func(t *testing.T) {
		got := Words("hello")
		if len(got) != 1 || got[0] != "hello" {
			t.Errorf("got %v", got)
		}
	})
}
