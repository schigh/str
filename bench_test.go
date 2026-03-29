package str

import "testing"

func BenchmarkPipe(b *testing.B) {
	fn := Pipe(TrimSpace, CollapseWhitespace, ToSnakeCase)
	input := "  Hello World  "
	for b.Loop() {
		fn(input)
	}
}

func BenchmarkToSnakeCase(b *testing.B) {
	for b.Loop() {
		ToSnakeCase("getHTTPSURLForAPI")
	}
}

func BenchmarkToSnakeCaseSimple(b *testing.B) {
	for b.Loop() {
		ToSnakeCase("helloWorld")
	}
}

func BenchmarkPadLeft(b *testing.B) {
	fn := PadLeft("0", 20)
	for b.Loop() {
		fn("1234")
	}
}

func BenchmarkTruncate(b *testing.B) {
	fn := Truncate(10, "...")
	input := "this is a longer string that needs truncation"
	for b.Loop() {
		fn(input)
	}
}

func BenchmarkReverse(b *testing.B) {
	for b.Loop() {
		Reverse("hello world 你好世界")
	}
}

func BenchmarkLevenshteinShort(b *testing.B) {
	for b.Loop() {
		Levenshtein("kitten", "sitting")
	}
}

func BenchmarkLevenshteinLong(b *testing.B) {
	a := "the quick brown fox jumps over the lazy dog"
	c := "the quack brown fix jumps over the lazy cat"
	for b.Loop() {
		Levenshtein(a, c)
	}
}

func BenchmarkJaroWinkler(b *testing.B) {
	for b.Loop() {
		JaroWinkler("martha", "marhta")
	}
}

func BenchmarkSlugifyASCII(b *testing.B) {
	for b.Loop() {
		SlugifyASCII("Hello, World! This is a Test.")
	}
}

func BenchmarkCollapseWhitespace(b *testing.B) {
	for b.Loop() {
		CollapseWhitespace("hello   \t\n  world   \t  foo")
	}
}
