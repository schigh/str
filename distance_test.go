package str

import (
	"math"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected int
	}{
		{"identical", "hello", "hello", 0},
		{"completely different", "abc", "xyz", 3},
		{"one empty", "", "hello", 5},
		{"both empty", "", "", 0},
		{"insertion", "hello", "helloo", 1},
		{"deletion", "hello", "hell", 1},
		{"substitution", "hello", "hallo", 1},
		{"unicode", "你好", "你坏", 1},
		{"swap order", "abc", "xyz", Levenshtein("xyz", "abc")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Levenshtein(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Levenshtein(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestJaroWinkler(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		min  float64
		max  float64
	}{
		{"identical", "hello", "hello", 1.0, 1.0},
		{"both empty", "", "", 1.0, 1.0},
		{"one empty", "", "hello", 0.0, 0.0},
		{"similar", "martha", "marhta", 0.96, 1.0},
		{"completely different", "abc", "xyz", 0.0, 0.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JaroWinkler(tt.a, tt.b)
			if got < tt.min-0.001 || got > tt.max+0.001 {
				t.Errorf("JaroWinkler(%q, %q) = %f, want [%f, %f]", tt.a, tt.b, got, tt.min, tt.max)
			}
		})
	}

	// Symmetry test.
	t.Run("symmetric", func(t *testing.T) {
		a := JaroWinkler("hello", "world")
		b := JaroWinkler("world", "hello")
		if math.Abs(a-b) > 0.001 {
			t.Errorf("not symmetric: %f vs %f", a, b)
		}
	})
}
