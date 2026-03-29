package str

import (
	"sort"
	"strings"
)

// Match represents a fuzzy match result.
type Match struct {
	Score float64 // 0.0 to 1.0 (JaroWinkler similarity)
	Index int     // position in the haystack
	Value string  // the matched string
}

// FuzzyMatch scores each haystack entry against the needle using JaroWinkler.
// Returns entries with Score >= 0.7, sorted by Score descending.
// Comparison is case-insensitive. Empty needle or haystack returns nil.
func FuzzyMatch(needle string, haystack []string) []Match {
	return fuzzyMatch(needle, haystack, 0.7)
}

// FuzzyMatchAll is like FuzzyMatch but returns all entries with Score > 0.
func FuzzyMatchAll(needle string, haystack []string) []Match {
	return fuzzyMatch(needle, haystack, 0.001)
}

func fuzzyMatch(needle string, haystack []string, threshold float64) []Match {
	if needle == "" || len(haystack) == 0 {
		return nil
	}
	low := strings.ToLower(needle)
	var matches []Match
	for i, h := range haystack {
		score := JaroWinkler(low, strings.ToLower(h))
		if score >= threshold {
			matches = append(matches, Match{
				Score: score,
				Index: i,
				Value: h,
			})
		}
	}
	if len(matches) == 0 {
		return nil
	}
	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].Score > matches[j].Score
	})
	return matches
}
