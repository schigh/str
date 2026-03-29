package str

import "testing"

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		maxLen   int
		ellipsis string
		input    string
		expected string
	}{
		{"normal truncation", 8, "...", "hello world", "hello..."},
		{"input fits", 20, "...", "hello", "hello"},
		{"maxLen zero", 0, "...", "hello", ""},
		{"ellipsis >= maxLen", 2, "...", "hello", ".."},
		{"empty ellipsis", 5, "", "hello world", "hello"},
		{"rune counting", 5, "...", "你好世界abcde", "你好..."},
		{"empty input", 5, "...", "", ""},
		{"exact length", 5, "...", "hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Truncate(tt.maxLen, tt.ellipsis)(tt.input)
			if got != tt.expected {
				t.Errorf("Truncate(%d, %q)(%q) = %q, want %q", tt.maxLen, tt.ellipsis, tt.input, got, tt.expected)
			}
		})
	}
}
