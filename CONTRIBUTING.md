# Contributing to str

## Design Language

Every transformation function in this library has the signature `func(string) string`. This is the library's identity. It makes every function composable with `Pipe`.

## API Design Rules

Three calling patterns, one decision tree:

1. **Direct** (`func(string) string`) — no configuration needed.
   - Examples: `ToSnakeCase`, `Reverse`, `CollapseWhitespace`

2. **Partial application** (`func(params) func(string) string`) — 1-2 required parameters.
   - Examples: `PadLeft(pad, width)`, `Truncate(maxLen, ellipsis)`

3. **Functional options** (`func(required, opts...) func(string) string`) — required param + optional modifiers.
   - Examples: `Wrap(width, opts...)`

**Decision tree for new functions:**
- No config? -> Direct
- All params required, 1-2 params? -> Partial application
- Required + optional behavior? -> Functional options
- Not a string transformation? -> Standalone, document clearly

## Naming Conventions

- Transformers: verb form (`Truncate`, `Reverse`, `Wrap`)
- Case conversions: `To{Format}` (`ToSnakeCase`, `ToCamelCase`)
- Options: `With{Thing}` for modifiers (`WithLineBreak`), bare name for strategies (`WrapAfterWord`)

## Constraints

- **Zero dependencies.** Stdlib only. No test dependencies either.
- **No panics.** Every function handles degenerate input gracefully.
- **Rune-based counting.** All length/width/index parameters count runes, not bytes.
- **Thread safe.** No shared mutable state.

## Testing

- Table-driven tests using stdlib `testing` only.
- Every function needs tests for edge cases (empty string, zero values, multibyte input).
- Every exported function needs a godoc `Example` test.
- Run `go test -race ./...` before submitting.

## File Organization

One concern per file. Test files mirror source files 1:1. New functions go in the file that matches their category, or get a new file if they don't fit.
