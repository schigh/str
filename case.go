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
//     a matching prefix becomes one word, then scanning continues on the remainder
//     - a non-acronym uppercase run followed by a lowercase letter splits so the
//     last uppercase letter begins the next word (e.g. "XML" in "XMLParser" stays
//     together, but unknown runs like "ABCdef" split as "AB" + "Cdef")
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

type segmentSplitter struct {
	runes    []rune
	acronyms []string
	words    []string
	current  []rune
	pos      int
}

func splitSegment(seg string, acronyms []string) []string {
	s := &segmentSplitter{
		runes:    []rune(seg),
		acronyms: acronyms,
	}
	s.split()
	return s.words
}

func (s *segmentSplitter) flush() {
	if len(s.current) > 0 {
		s.words = append(s.words, strings.ToLower(string(s.current)))
		s.current = nil
	}
}

func (s *segmentSplitter) split() {
	for s.pos < len(s.runes) {
		r := s.runes[s.pos]

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			s.pos++
			continue
		}

		if unicode.IsDigit(r) {
			s.handleDigit(r)
			continue
		}

		s.handleLetter(r)
	}

	if len(s.current) > 0 {
		s.words = append(s.words, strings.ToLower(string(s.current)))
	}
}

func (s *segmentSplitter) handleDigit(r rune) {
	if len(s.current) > 0 && unicode.IsLetter(s.current[len(s.current)-1]) {
		s.flush()
	}
	s.current = append(s.current, r)
	s.pos++
}

func (s *segmentSplitter) handleLetter(r rune) {
	if len(s.current) > 0 && unicode.IsDigit(s.current[len(s.current)-1]) && unicode.IsUpper(r) {
		s.flush()
	}

	if unicode.IsLower(r) {
		s.current = append(s.current, r)
		s.pos++
		return
	}

	s.processUpperRun()
}

func (s *segmentSplitter) processUpperRun() {
	start := s.pos
	for s.pos < len(s.runes) && unicode.IsUpper(s.runes[s.pos]) {
		s.pos++
	}
	upperRun := s.runes[start:s.pos]

	if len(s.current) > 0 {
		s.flush()
	}

	followedByLower := s.pos < len(s.runes) && unicode.IsLower(s.runes[s.pos])

	j := 0
	for j < len(upperRun) {
		if n := s.matchAcronym(upperRun[j:]); n > 0 {
			s.flush()
			s.words = append(s.words, strings.ToLower(string(upperRun[j:j+n])))
			j += n
			continue
		}

		s.handleNonAcronymUpper(upperRun, j, followedByLower)
		break
	}
}

func (s *segmentSplitter) matchAcronym(run []rune) int {
	remaining := string(run)
	for _, acr := range s.acronyms {
		if strings.HasPrefix(remaining, acr) {
			return len([]rune(acr))
		}
	}
	return 0
}

func (s *segmentSplitter) handleNonAcronymUpper(upperRun []rune, j int, followedByLower bool) {
	if j == len(upperRun)-1 && followedByLower {
		s.flush()
		s.current = append(s.current, upperRun[j])
		return
	}

	endOfRun := len(upperRun)
	if followedByLower {
		endOfRun = len(upperRun) - 1
	}

	if endOfRun > j {
		s.flush()
		s.current = append(s.current, upperRun[j:endOfRun]...)
	}

	if endOfRun < len(upperRun) {
		s.flush()
		s.current = append(s.current, upperRun[endOfRun:]...)
	}
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
