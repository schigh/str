[![CI](https://github.com/schigh/str/actions/workflows/ci.yml/badge.svg)](https://github.com/schigh/str/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/schigh/str.svg)](https://pkg.go.dev/github.com/schigh/str)
[![Go Report Card](https://goreportcard.com/badge/github.com/schigh/str)](https://goreportcard.com/report/github.com/schigh/str)

# str

Pipeline-first string toolkit for Go. Compose transformations, not calls.

```go
normalize := str.Pipe(
    str.TrimSpace,
    str.CollapseWhitespace,
    str.ToSnakeCase,
    str.Truncate(64, "..."),
)
slug := normalize("  HelloWorld  ") // "hello_world"
```

Every function is `func(string) string`. Compose them with `Pipe`. Zero dependencies.

## Install

```
go get github.com/schigh/str@latest
```

## Compose

**Pipe** composes multiple transformers into a single function, applied left to right.

```go
format := str.Pipe(str.TrimSpace, str.PadLeft("0", 8))
format("  42  ") // "00000042"
```

**PipeErr** composes fallible transformers. Short-circuits on the first error.

**PipeIf** applies a transformation only when a predicate is true. Otherwise passes through.

```go
isLong := func(s string) bool { return len([]rune(s)) > 100 }
clean := str.Pipe(
    str.TrimSpace,
    str.PipeIf(isLong, str.Truncate(100, "...")),
)
```

**PipeUnless** is the inverse: applies when the predicate is false.

**PipeMap** splits a string, transforms each part, and rejoins.

```go
trimLines := str.PipeMap(str.Lines, "\n", str.TrimSpace)
trimLines("  hello  \n  world  ") // "hello\nworld"
```

Built-in splitters: `Lines` (split on `\n`) and `Words` (split on whitespace).

## Transform

**ToSnakeCase** / **ToCamelCase** / **ToPascalCase** / **ToKebabCase** / **ToTitleCase** / **ToScreamingSnake**

```go
str.ToSnakeCase("HTMLParser")   // "html_parser"
str.ToCamelCase("hello_world")  // "helloWorld"
str.ToPascalCase("hello_world") // "HelloWorld"
str.ToKebabCase("HelloWorld")   // "hello-world"
str.ToTitleCase("hello_world")  // "Hello World"
str.ToScreamingSnake("hello")   // "HELLO"
```

Handles acronyms (ID, URL, HTTP, HTML, API, JSON, etc.), digit boundaries, and Unicode letters.

**Reverse** reverses runes in a string.

```go
str.Reverse("hello") // "olleh"
```

**CollapseWhitespace** collapses all Unicode whitespace to a single space.

```go
str.CollapseWhitespace("hello   \t\n  world") // "hello world"
```

**SlugifyASCII** generates URL-safe ASCII slugs.

```go
str.SlugifyASCII("Hello, World!") // "hello-world"
```

## Format

**PadLeft** / **PadRight** pad to a target width (runes). Multi-char pads are truncated to hit exact width.

```go
str.PadLeft("0", 8)("1234")   // "00001234"
str.PadRight(".", 10)("hello") // "hello....."
```

**Truncate** truncates to a max length (runes) with an ellipsis. Length includes the ellipsis.

```go
str.Truncate(8, "...")("hello world") // "hello..."
```

**Substring** extracts a substring with negative indexing support.

```go
str.Substring(0, 3)("hello")  // "hel"
str.Substring(-3, 2)("abcde") // "cd"
```

**Wrap** wraps text at a given width with configurable behavior.

```go
str.Wrap(40)("long text...")                           // wraps before words (default)
str.Wrap(40, str.WrapAfterWord)("long text...")        // wraps after words
str.Wrap(76, str.WrapHardBreak)("base64encoded...")    // hard break at width
```

## Measure

**Levenshtein** returns the edit distance between two strings.

```go
str.Levenshtein("kitten", "sitting") // 3
```

**JaroWinkler** returns a similarity score between 0.0 and 1.0.

```go
str.JaroWinkler("martha", "marhta") // ~0.96
```

## Pipeline Patterns

```go
// Normalize user input for storage
var normalizeInput = str.Pipe(
    str.TrimSpace,
    str.CollapseWhitespace,
)

// Generate URL slugs
var slugify = str.Pipe(
    str.TrimSpace,
    str.CollapseWhitespace,
    str.SlugifyASCII,
)

// Format log fields
var formatField = str.Pipe(
    str.TrimSpace,
    str.Truncate(50, "..."),
    str.PadRight(" ", 50),
)
```

## Roadmap

Phase 2 will add: random token generation, fuzzy matching, transliteration, grapheme-aware reverse, Unicode-aware slugify, conditional pipe composition, and customizable acronym lists.

## License

MIT
