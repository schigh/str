package str

import (
	"slices"
	"testing"
)

func TestNaturalSortKey(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"empty", "", ""},
		{"no numbers", "hello", "hello"},
		{"single digit", "file2", "file00000000000000000002"},
		{"multi digit", "file10", "file00000000000000000010"},
		{"leading zeros", "file007", "file00000000000000000007"},
		{"multiple segments", "v1.2.10", "v00000000000000000001.00000000000000000002.00000000000000000010"},
		{"case folding", "File2", "file00000000000000000002"},
		{"over 20 digits", "x123456789012345678901y", "x123456789012345678901y"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NaturalSortKey(tt.input)
			if got != tt.expected {
				t.Errorf("NaturalSortKey(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNaturalSortOrder(t *testing.T) {
	input := []string{"file10", "file2", "file1", "file20", "file3"}
	expected := []string{"file1", "file2", "file3", "file10", "file20"}

	slices.SortFunc(input, func(a, b string) int {
		ka, kb := NaturalSortKey(a), NaturalSortKey(b)
		if ka < kb {
			return -1
		}
		if ka > kb {
			return 1
		}
		return 0
	})

	for i, got := range input {
		if got != expected[i] {
			t.Errorf("index %d: got %q, want %q", i, got, expected[i])
		}
	}
}
