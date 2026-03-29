package str

import "math"

// Levenshtein returns the edit distance between two strings.
// Uses a 2-row rolling array for O(min(m,n)) space complexity.
func Levenshtein(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)

	// Ensure ra is the shorter string for space optimization.
	if len(ra) > len(rb) {
		ra, rb = rb, ra
	}

	m := len(ra)
	n := len(rb)

	if m == 0 {
		return n
	}

	prev := make([]int, m+1)
	curr := make([]int, m+1)

	for i := 0; i <= m; i++ {
		prev[i] = i
	}

	for j := 1; j <= n; j++ {
		curr[0] = j
		for i := 1; i <= m; i++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			curr[i] = min(
				curr[i-1]+1,
				min(prev[i]+1, prev[i-1]+cost),
			)
		}
		prev, curr = curr, prev
	}

	return prev[m]
}

// JaroWinkler returns the Jaro-Winkler similarity between two strings.
// Returns a value between 0.0 (completely different) and 1.0 (identical).
// Uses the standard Winkler prefix bonus scaling factor p=0.1 with max prefix length 4.
func JaroWinkler(a, b string) float64 {
	ra := []rune(a)
	rb := []rune(b)

	if len(ra) == 0 && len(rb) == 0 {
		return 1.0
	}
	if len(ra) == 0 || len(rb) == 0 {
		return 0.0
	}

	jaro := jaroSimilarity(ra, rb)

	// Winkler prefix bonus
	prefixLen := 0
	maxPrefix := min(4, min(len(ra), len(rb)))
	for i := range maxPrefix {
		if ra[i] == rb[i] {
			prefixLen++
		} else {
			break
		}
	}

	const p = 0.1
	return jaro + float64(prefixLen)*p*(1.0-jaro)
}

func jaroSimilarity(a, b []rune) float64 {
	aLen := len(a)
	bLen := len(b)

	if aLen == 0 && bLen == 0 {
		return 1.0
	}

	matchDist := max(0, max(aLen, bLen)/2-1)

	aMatched := make([]bool, aLen)
	bMatched := make([]bool, bLen)

	matches := 0
	transpositions := 0

	for i := range aLen {
		lo := max(0, i-matchDist)
		hi := min(i+matchDist+1, bLen)
		for j := lo; j < hi; j++ {
			if bMatched[j] || a[i] != b[j] {
				continue
			}
			aMatched[i] = true
			bMatched[j] = true
			matches++
			break
		}
	}

	if matches == 0 {
		return 0.0
	}

	k := 0
	for i := range aLen {
		if !aMatched[i] {
			continue
		}
		for !bMatched[k] {
			k++
		}
		if a[i] != b[k] {
			transpositions++
		}
		k++
	}

	m := float64(matches)
	return (m/float64(aLen) + m/float64(bLen) + (m-math.Floor(float64(transpositions)/2.0))/m) / 3.0
}
