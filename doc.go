// Package str is a pipeline-first string toolkit for Go.
//
// Every transformation function has the signature func(string) string,
// making them composable with [Pipe]. Functions that need configuration
// use partial application to return func(string) string.
//
// All length, width, and index parameters count runes, not bytes.
// All functions are safe for concurrent use.
// No function panics for any input.
//
// # Compose
//
// [Pipe] and [PipeErr] compose multiple transformations into a single function.
//
// # Transform
//
// Case conversion, reversal, whitespace normalization, and slug generation.
//
// # Format
//
// Padding, truncation, substring extraction, and word wrapping.
//
// # Measure
//
// String distance and similarity algorithms.
package str
