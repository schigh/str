package str

import "testing"

func TestPadLeft(t *testing.T) {
	tests := []struct {
		name     string
		pad      string
		width    int
		input    string
		expected string
	}{
		{"single char", "0", 8, "1234", "00001234"},
		{"multi char exact", "ab", 8, "1234", "abab1234"},
		{"multi char truncated", "abc", 8, "1234", "abca1234"},
		{"empty pad", "", 8, "1234", "1234"},
		{"width equals input", "0", 4, "1234", "1234"},
		{"width less than input", "0", 2, "1234", "1234"},
		{"width zero", "0", 0, "1234", "1234"},
		{"empty input", "0", 4, "", "0000"},
		{"multibyte input", "0", 6, "你好", "0000你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PadLeft(tt.pad, tt.width)(tt.input)
			if got != tt.expected {
				t.Errorf("PadLeft(%q, %d)(%q) = %q, want %q", tt.pad, tt.width, tt.input, got, tt.expected)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		name     string
		pad      string
		width    int
		input    string
		expected string
	}{
		{"single char", "0", 8, "1234", "12340000"},
		{"multi char", "ab", 8, "1234", "1234abab"},
		{"empty pad", "", 8, "1234", "1234"},
		{"no padding needed", "0", 4, "1234", "1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PadRight(tt.pad, tt.width)(tt.input)
			if got != tt.expected {
				t.Errorf("PadRight(%q, %d)(%q) = %q, want %q", tt.pad, tt.width, tt.input, got, tt.expected)
			}
		})
	}
}
