package str

import "testing"

func TestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		length   int
		input    string
		expected string
	}{
		{"normal", 0, 3, "1234567890", "123"},
		{"offset", 1, 3, "1234567890", "234"},
		{"negative start", -3, 2, "abcde", "cd"},
		{"negative start full", -3, 3, "abcde", "cde"},
		{"negative length", 0, -1, "hello", ""},
		{"zero length", 0, 0, "hello", ""},
		{"out of bounds start", 20, 3, "hello", ""},
		{"length exceeds", 3, 100, "hello", "lo"},
		{"empty input", 0, 3, "", ""},
		{"negative out of bounds", -10, 3, "hello", "hel"},
		{"multibyte", 0, 2, "你好世界", "你好"},
		{"negative start single", -1, 3, "1234567890", "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Substring(tt.start, tt.length)(tt.input)
			if got != tt.expected {
				t.Errorf("Substring(%d, %d)(%q) = %q, want %q", tt.start, tt.length, tt.input, got, tt.expected)
			}
		})
	}
}
