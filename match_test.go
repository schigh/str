package str

import "testing"

func TestFuzzyMatch(t *testing.T) {
	hay := []string{"hello", "help", "world", "helm", "hero"}

	t.Run("normal match", func(t *testing.T) {
		matches := FuzzyMatch("helo", hay)
		if len(matches) == 0 {
			t.Fatal("expected matches")
		}
		if matches[0].Value != "hello" {
			t.Errorf("best match got %q, want %q", matches[0].Value, "hello")
		}
	})

	t.Run("exact match score 1.0", func(t *testing.T) {
		matches := FuzzyMatch("hello", hay)
		if len(matches) == 0 {
			t.Fatal("expected matches")
		}
		if matches[0].Score != 1.0 {
			t.Errorf("exact match score got %f, want 1.0", matches[0].Score)
		}
	})

	t.Run("case insensitive", func(t *testing.T) {
		matches := FuzzyMatch("HELLO", hay)
		if len(matches) == 0 {
			t.Fatal("expected matches")
		}
		if matches[0].Value != "hello" {
			t.Errorf("got %q, want %q", matches[0].Value, "hello")
		}
	})

	t.Run("empty needle", func(t *testing.T) {
		if FuzzyMatch("", hay) != nil {
			t.Error("expected nil for empty needle")
		}
	})

	t.Run("empty haystack", func(t *testing.T) {
		if FuzzyMatch("hello", nil) != nil {
			t.Error("expected nil for empty haystack")
		}
	})

	t.Run("no matches above threshold", func(t *testing.T) {
		matches := FuzzyMatch("zzzzz", hay)
		if matches != nil {
			t.Errorf("expected nil, got %d matches", len(matches))
		}
	})

	t.Run("sorted by score descending", func(t *testing.T) {
		matches := FuzzyMatch("hel", hay)
		for i := 1; i < len(matches); i++ {
			if matches[i].Score > matches[i-1].Score {
				t.Errorf("not sorted: %f > %f", matches[i].Score, matches[i-1].Score)
			}
		}
	})

	t.Run("index preserved", func(t *testing.T) {
		matches := FuzzyMatch("hello", hay)
		if matches[0].Index != 0 {
			t.Errorf("got index %d, want 0", matches[0].Index)
		}
	})
}

func TestFuzzyMatchAll(t *testing.T) {
	hay := []string{"hello", "xyz", "help"}

	t.Run("returns low-scoring entries", func(t *testing.T) {
		all := FuzzyMatchAll("hello", hay)
		filtered := FuzzyMatch("hello", hay)
		if len(all) < len(filtered) {
			t.Errorf("FuzzyMatchAll returned fewer results (%d) than FuzzyMatch (%d)", len(all), len(filtered))
		}
	})

	t.Run("empty needle", func(t *testing.T) {
		if FuzzyMatchAll("", hay) != nil {
			t.Error("expected nil")
		}
	})
}
