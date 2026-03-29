package str

import (
	"errors"
	"strings"
	"testing"
)

func TestPipe(t *testing.T) {
	t.Run("zero functions", func(t *testing.T) {
		fn := Pipe()
		if got := fn("hello"); got != "hello" {
			t.Errorf("Pipe()(%q) = %q, want %q", "hello", got, "hello")
		}
	})

	t.Run("single function", func(t *testing.T) {
		fn := Pipe(strings.ToUpper)
		if got := fn("hello"); got != "HELLO" {
			t.Errorf("got %q, want %q", got, "HELLO")
		}
	})

	t.Run("multiple functions in order", func(t *testing.T) {
		fn := Pipe(TrimSpace, ToSnakeCase)
		if got := fn("  HelloWorld  "); got != "hello_world" {
			t.Errorf("got %q, want %q", got, "hello_world")
		}
	})

	t.Run("composition is reusable", func(t *testing.T) {
		fn := Pipe(strings.ToUpper)
		if fn("a") != "A" || fn("b") != "B" {
			t.Error("pipe should be reusable")
		}
	})
}

func TestPipeIf(t *testing.T) {
	isLong := func(s string) bool { return len(s) > 5 }

	t.Run("predicate true applies fn", func(t *testing.T) {
		fn := PipeIf(isLong, strings.ToUpper)
		if got := fn("hello world"); got != "HELLO WORLD" {
			t.Errorf("got %q, want %q", got, "HELLO WORLD")
		}
	})

	t.Run("predicate false passes through", func(t *testing.T) {
		fn := PipeIf(isLong, strings.ToUpper)
		if got := fn("hi"); got != "hi" {
			t.Errorf("got %q, want %q", got, "hi")
		}
	})

	t.Run("nil predicate is no-op", func(t *testing.T) {
		fn := PipeIf(nil, strings.ToUpper)
		if got := fn("hello"); got != "hello" {
			t.Errorf("got %q, want %q", got, "hello")
		}
	})

	t.Run("nil fn is no-op", func(t *testing.T) {
		fn := PipeIf(isLong, nil)
		if got := fn("hello world"); got != "hello world" {
			t.Errorf("got %q, want %q", got, "hello world")
		}
	})

	t.Run("composes in Pipe", func(t *testing.T) {
		fn := Pipe(
			TrimSpace,
			PipeIf(isLong, strings.ToUpper),
		)
		if got := fn("  hello world  "); got != "HELLO WORLD" {
			t.Errorf("got %q, want %q", got, "HELLO WORLD")
		}
		if got := fn("  hi  "); got != "hi" {
			t.Errorf("got %q, want %q", got, "hi")
		}
	})

	t.Run("empty string", func(t *testing.T) {
		fn := PipeIf(isLong, strings.ToUpper)
		if got := fn(""); got != "" {
			t.Errorf("got %q, want empty", got)
		}
	})
}

func TestPipeUnless(t *testing.T) {
	isShort := func(s string) bool { return len(s) <= 5 }

	t.Run("predicate false applies fn", func(t *testing.T) {
		fn := PipeUnless(isShort, strings.ToUpper)
		if got := fn("hello world"); got != "HELLO WORLD" {
			t.Errorf("got %q, want %q", got, "HELLO WORLD")
		}
	})

	t.Run("predicate true passes through", func(t *testing.T) {
		fn := PipeUnless(isShort, strings.ToUpper)
		if got := fn("hi"); got != "hi" {
			t.Errorf("got %q, want %q", got, "hi")
		}
	})

	t.Run("nil predicate always applies fn", func(t *testing.T) {
		fn := PipeUnless(nil, strings.ToUpper)
		if got := fn("hello"); got != "HELLO" {
			t.Errorf("got %q, want %q", got, "HELLO")
		}
	})

	t.Run("nil fn is no-op", func(t *testing.T) {
		fn := PipeUnless(isShort, nil)
		if got := fn("hello world"); got != "hello world" {
			t.Errorf("got %q, want %q", got, "hello world")
		}
	})
}

func TestPipeErr(t *testing.T) {
	good := func(s string) (string, error) { return strings.ToUpper(s), nil }
	bad := func(s string) (string, error) { return "", errors.New("fail") }

	t.Run("zero functions", func(t *testing.T) {
		fn := PipeErr()
		got, err := fn("hello")
		if err != nil || got != "hello" {
			t.Errorf("got %q, err=%v, want %q, nil", got, err, "hello")
		}
	})

	t.Run("all succeed", func(t *testing.T) {
		fn := PipeErr(good, good)
		got, err := fn("hello")
		if err != nil || got != "HELLO" {
			t.Errorf("got %q, err=%v", got, err)
		}
	})

	t.Run("first function errors", func(t *testing.T) {
		fn := PipeErr(bad, good)
		_, err := fn("hello")
		if err == nil {
			t.Error("expected error")
		}
	})

	t.Run("middle function errors short-circuits", func(t *testing.T) {
		called := false
		third := func(s string) (string, error) {
			called = true
			return s, nil
		}
		fn := PipeErr(good, bad, third)
		_, err := fn("hello")
		if err == nil {
			t.Error("expected error")
		}
		if called {
			t.Error("third function should not have been called")
		}
	})
}
