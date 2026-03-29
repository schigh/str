package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CaseOption configures case conversion behavior.
type CaseOption func(*caseConfig)

type caseConfig struct {
	acronyms []string
}

// WithAcronyms replaces the default acronym list entirely. Users who want
// defaults plus custom acronyms must include both.
func WithAcronyms(words ...string) CaseOption {
	sorted := sortAcronymsByLength(words)
	return func(c *caseConfig) { c.acronyms = sorted }
}

// sorted by length descending for longest-prefix matching
var defaultAcronyms = []string{
	"HTTPS", "HTML", "HTTP", "JSON", "UUID",
	"API", "CSS", "DNS", "EOF", "SQL",
	"SSH", "SSL", "TCP", "TLS", "UDP",
	"URL", "XML",
	"ID", "IP",
}

func sortAcronymsByLength(words []string) []string {
	upper := make([]string, len(words))
	for i, w := range words {
		upper[i] = strings.ToUpper(w)
	}
	// Simple insertion sort by length descending (lists are small).
	for i := 1; i < len(upper); i++ {
		for j := i; j > 0 && len(upper[j]) > len(upper[j-1]); j-- {
			upper[j], upper[j-1] = upper[j-1], upper[j]
		}
	}
	return upper
}

func applyCaseOpts(opts []CaseOption) *caseConfig {
	cfg := &caseConfig{acronyms: defaultAcronyms}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// English title case skip list: articles, conjunctions, prepositions, and
// other common short function words.
var titleCaseSkipWords = map[string]bool{
	"a": true, "an": true, "the": true,
	"and": true, "but": true, "or": true, "nor": true,
	"for": true, "yet": true, "so": true,
	"at": true, "by": true, "in": true, "of": true,
	"on": true, "to": true, "up": true,
	"as": true, "is": true, "if": true,
}

// splitWords breaks s into lowercase words using a multi-pass algorithm:
//
//  1. Split on delimiters (spaces, underscores, hyphens) to get segments.
//  2. Within each segment, scan rune-by-rune for transitions:
//     - lowercase-to-uppercase starts a new word
//     - digit-to-letter or letter-to-digit starts a new word
//     - an uppercase run is checked against known acronyms (longest prefix first);
//       a matching prefix becomes one word, then scanning continues on the remainder
//     - a non-acronym uppercase run followed by a lowercase letter splits so the
//       last uppercase letter begins the next word (e.g. "XML" in "XMLParser" stays
//       together, but unknown runs like "ABCdef" split as "AB" + "Cdef")
//  3. Non-letter, non-digit runes (punctuation) are stripped.
//  4. Every word is lowercased before returning.
func splitWords(s string) []string {
	return splitWordsWithAcronyms(s, defaultAcronyms)
}

func splitWordsWithAcronyms(s string, acronyms []string) []string {
	if s == "" {
		return nil
	}

	segments := splitOnDelimiters(s)

	var words []string
	for _, seg := range segments {
		words = append(words, splitSegment(seg, acronyms)...)
	}
	return words
}

// Compatibility: old call in splitWords already passes acronyms through splitWordsWithAcronyms.

func splitOnDelimiters(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '_' || r == '-'
	})
}

func splitSegment(seg string, acronyms []string) []string {
	runes := []rune(seg)
	var words []string
	var current []rune

	flush := func() {
		if len(current) > 0 {
			words = append(words, strings.ToLower(string(current)))
			current = nil
		}
	}

	i := 0
	for i < len(runes) {
		r := runes[i]

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			// Strip punctuation.
			i++
			continue
		}

		if unicode.IsDigit(r) {
			// If current word is letters, start a new word for digits.
			if len(current) > 0 && unicode.IsLetter(current[len(current)-1]) {
				flush()
			}
			current = append(current, r)
			i++
			continue
		}

		// r is a letter.
		if len(current) > 0 && unicode.IsDigit(current[len(current)-1]) && unicode.IsUpper(r) {
			// Transition from digit to uppercase letter.
			flush()
		}

		if unicode.IsLower(r) {
			current = append(current, r)
			i++
			continue
		}

		// r is uppercase. Collect the full uppercase run.
		upperStart := i
		for i < len(runes) && unicode.IsUpper(runes[i]) {
			i++
		}
		upperRun := runes[upperStart:i]

		// If current word has lowercase content, flush before processing the uppercase run.
		if len(current) > 0 {
			flush()
		}

		// Try to consume the uppercase run using acronym matching.
		j := 0
		for j < len(upperRun) {
			remaining := string(upperRun[j:])
			matched := false

			for _, acr := range acronyms {
				if strings.HasPrefix(remaining, acr) {
					flush()
					words = append(words, strings.ToLower(acr))
					j += len([]rune(acr))
					matched = true
					break
				}
			}
			if matched {
				continue
			}

			// No acronym match. If this uppercase letter is followed by more uppercase
			// letters or is at the end, it could be part of a non-acronym run.
			// If the next char after the uppercase run is lowercase, the last uppercase
			// letter starts a new camelCase word.
			if j == len(upperRun)-1 && i < len(runes) && unicode.IsLower(runes[i]) {
				// Last uppercase letter begins a new word with the following lowercase.
				flush()
				current = append(current, upperRun[j])
				j++
			} else if j < len(upperRun)-1 {
				// Middle of a non-acronym uppercase run. Check if the tail transitions
				// into lowercase (meaning the last uppercase letter starts a new word).
				endOfRun := len(upperRun)
				if i < len(runes) && unicode.IsLower(runes[i]) {
					// The last uppercase letter of the run belongs with the lowercase.
					endOfRun = len(upperRun) - 1
				}
				// Emit remaining uppercase letters up to endOfRun as a single word.
				if endOfRun > j {
					flush()
					current = append(current, upperRun[j:endOfRun]...)
					j = endOfRun
				}
			} else {
				current = append(current, upperRun[j])
				j++
			}
		}
	}

	if len(current) > 0 {
		words = append(words, strings.ToLower(string(current)))
	}
	return words
}

func capitalize(word string) string {
	if word == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(word)
	return string(unicode.ToUpper(r)) + strings.ToLower(word[size:])
}

func ToSnakeCase(s string) string {
	return strings.Join(splitWords(s), "_")
}

func ToCamelCase(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(words[0])
	for _, w := range words[1:] {
		b.WriteString(capitalize(w))
	}
	return b.String()
}

func ToPascalCase(s string) string {
	words := splitWords(s)
	var b strings.Builder
	for _, w := range words {
		b.WriteString(capitalize(w))
	}
	return b.String()
}

func ToKebabCase(s string) string {
	return strings.Join(splitWords(s), "-")
}

func ToTitleCase(s string) string {
	words := splitWords(s)
	capitalized := make([]string, len(words))
	for i, w := range words {
		capitalized[i] = capitalize(w)
	}
	return strings.Join(capitalized, " ")
}

func ToScreamingSnake(s string) string {
	return strings.ToUpper(strings.Join(splitWords(s), "_"))
}

// ToSnakeCaseWith returns a snake_case transformer using custom options.
func ToSnakeCaseWith(opts ...CaseOption) func(string) string {
	cfg := applyCaseOpts(opts)
	return func(s string) string {
		return strings.Join(splitWordsWithAcronyms(s, cfg.acronyms), "_")
	}
}

// ToCamelCaseWith returns a camelCase transformer using custom options.
func ToCamelCaseWith(opts ...CaseOption) func(string) string {
	cfg := applyCaseOpts(opts)
	return func(s string) string {
		words := splitWordsWithAcronyms(s, cfg.acronyms)
		if len(words) == 0 {
			return ""
		}
		var b strings.Builder
		b.WriteString(words[0])
		for _, w := range words[1:] {
			b.WriteString(capitalize(w))
		}
		return b.String()
	}
}

// ToPascalCaseWith returns a PascalCase transformer using custom options.
func ToPascalCaseWith(opts ...CaseOption) func(string) string {
	cfg := applyCaseOpts(opts)
	return func(s string) string {
		words := splitWordsWithAcronyms(s, cfg.acronyms)
		var b strings.Builder
		for _, w := range words {
			b.WriteString(capitalize(w))
		}
		return b.String()
	}
}

// ToKebabCaseWith returns a kebab-case transformer using custom options.
func ToKebabCaseWith(opts ...CaseOption) func(string) string {
	cfg := applyCaseOpts(opts)
	return func(s string) string {
		return strings.Join(splitWordsWithAcronyms(s, cfg.acronyms), "-")
	}
}

// ToTitleCaseWith returns a Title Case transformer using custom options.
func ToTitleCaseWith(opts ...CaseOption) func(string) string {
	cfg := applyCaseOpts(opts)
	return func(s string) string {
		words := splitWordsWithAcronyms(s, cfg.acronyms)
		capitalized := make([]string, len(words))
		for i, w := range words {
			capitalized[i] = capitalize(w)
		}
		return strings.Join(capitalized, " ")
	}
}

// ToScreamingSnakeWith returns a SCREAMING_SNAKE transformer using custom options.
func ToScreamingSnakeWith(opts ...CaseOption) func(string) string {
	cfg := applyCaseOpts(opts)
	return func(s string) string {
		return strings.ToUpper(strings.Join(splitWordsWithAcronyms(s, cfg.acronyms), "_"))
	}
}

// ToTitleCaseEnglish returns English-aware title case. Capitalizes all words
// except short function words (articles, conjunctions, prepositions) when they
// appear mid-sentence. First and last word are always capitalized.
func ToTitleCaseEnglish(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}
	result := make([]string, len(words))
	for i, w := range words {
		if i == 0 || i == len(words)-1 || !titleCaseSkipWords[w] {
			result[i] = capitalize(w)
		} else {
			result[i] = w
		}
	}
	return strings.Join(result, " ")
}
