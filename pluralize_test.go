package str

import "testing"

func TestPluralize(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		plural   string
		singular string
		expected string
	}{
		{"count 1 singular", 1, "items", "item", "item"},
		{"count 0 plural", 0, "items", "item", "items"},
		{"count 5 plural", 5, "items", "item", "items"},
		{"count -1 plural", -1, "items", "item", "items"},
		{"empty singular", 1, "items", "", ""},
		{"empty plural", 0, "", "item", ""},
		{"both empty", 1, "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pluralize(tt.count, tt.plural)(tt.singular)
			if got != tt.expected {
				t.Errorf("Pluralize(%d, %q)(%q) = %q, want %q", tt.count, tt.plural, tt.singular, got, tt.expected)
			}
		})
	}
}

func TestPluralizeInPipe(t *testing.T) {
	fn := Pipe(TrimSpace, Pluralize(3, "foxes"))
	got := fn("  fox  ")
	if got != "foxes" {
		t.Errorf("got %q, want %q", got, "foxes")
	}
}
