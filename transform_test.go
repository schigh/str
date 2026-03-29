package str

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"ascii", "hello", "olleh"},
		{"multibyte", "你好", "好你"},
		{"empty", "", ""},
		{"single", "a", "a"},
		{"palindrome", "racecar", "racecar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.expected {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCollapseWhitespace(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"spaces", "hello   world", "hello world"},
		{"tabs", "hello\t\tworld", "hello world"},
		{"newlines", "hello\n\nworld", "hello world"},
		{"mixed", "hello \t\n world", "hello world"},
		{"leading trailing", "  hello  ", "hello"},
		{"empty", "", ""},
		{"only whitespace", "   \t\n  ", ""},
		{"no extra whitespace", "hello world", "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CollapseWhitespace(tt.input)
			if got != tt.expected {
				t.Errorf("CollapseWhitespace(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	got := TrimSpace("  hello  ")
	if got != "hello" {
		t.Errorf("TrimSpace(%q) = %q, want %q", "  hello  ", got, "hello")
	}
}

func TestSlugifyASCII(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"normal", "Hello World", "hello-world"},
		{"special chars", "Hello, World!", "hello-world"},
		{"consecutive hyphens", "Hello---World", "hello-world"},
		{"all punctuation", "!@#$%", ""},
		{"empty", "", ""},
		{"unicode stripped", "café", "caf"},
		{"mixed", "The Quick Brown Fox!", "the-quick-brown-fox"},
		{"underscores", "hello_world", "hello-world"},
		{"leading trailing punct", "---hello---", "hello"},
		{"numbers", "version 2.0", "version-2-0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SlugifyASCII(tt.input)
			if got != tt.expected {
				t.Errorf("SlugifyASCII(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
