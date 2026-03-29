package str

// Pluralize returns a transformer that selects between the input (singular)
// and the given plural form based on count. Returns singular when count == 1,
// plural otherwise (including count == 0).
func Pluralize(count int, plural string) func(string) string {
	return func(singular string) string {
		if count == 1 {
			return singular
		}
		return plural
	}
}
