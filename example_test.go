package str_test

import (
	"fmt"

	"github.com/schigh/str"
)

func ExamplePipe() {
	normalize := str.Pipe(
		str.TrimSpace,
		str.CollapseWhitespace,
		str.ToSnakeCase,
	)
	fmt.Println(normalize("  HelloWorld  "))
	// Output: hello_world
}

func ExamplePipeErr() {
	upper := func(s string) (string, error) { return str.Reverse(s), nil }
	fn := str.PipeErr(upper)
	result, _ := fn("hello")
	fmt.Println(result)
	// Output: olleh
}

func ExampleToSnakeCase() {
	fmt.Println(str.ToSnakeCase("HTMLParser"))
	fmt.Println(str.ToSnakeCase("getHTTPSURL"))
	fmt.Println(str.ToSnakeCase("UserID"))
	// Output:
	// html_parser
	// get_https_url
	// user_id
}

func ExampleToSnakeCase_pipe() {
	slugify := str.Pipe(str.TrimSpace, str.ToSnakeCase)
	fmt.Println(slugify("  Hello World  "))
	// Output: hello_world
}

func ExampleToCamelCase() {
	fmt.Println(str.ToCamelCase("hello_world"))
	fmt.Println(str.ToCamelCase("HTMLParser"))
	// Output:
	// helloWorld
	// htmlParser
}

func ExampleToPascalCase() {
	fmt.Println(str.ToPascalCase("hello_world"))
	// Output: HelloWorld
}

func ExampleToKebabCase() {
	fmt.Println(str.ToKebabCase("HelloWorld"))
	// Output: hello-world
}

func ExampleToTitleCase() {
	fmt.Println(str.ToTitleCase("hello_world"))
	// Output: Hello World
}

func ExampleToScreamingSnake() {
	fmt.Println(str.ToScreamingSnake("helloWorld"))
	// Output: HELLO_WORLD
}

func ExamplePadLeft() {
	fmt.Println(str.PadLeft("0", 8)("1234"))
	// Output: 00001234
}

func ExamplePadLeft_pipe() {
	format := str.Pipe(str.TrimSpace, str.PadLeft("0", 8))
	fmt.Println(format("  42  "))
	// Output: 00000042
}

func ExamplePadRight() {
	fmt.Println(str.PadRight(".", 10)("hello"))
	// Output: hello.....
}

func ExampleTruncate() {
	fmt.Println(str.Truncate(8, "...")("hello world"))
	// Output: hello...
}

func ExampleTruncate_pipe() {
	shorten := str.Pipe(str.TrimSpace, str.Truncate(10, "..."))
	fmt.Println(shorten("  a]very long string  "))
	// Output: a]very ...
}

func ExampleSubstring() {
	fmt.Println(str.Substring(0, 3)("hello"))
	fmt.Println(str.Substring(-3, 2)("abcde"))
	// Output:
	// hel
	// cd
}

func ExampleWrap() {
	fmt.Println(str.Wrap(10, str.WrapHardBreak)("abcdefghijklmno"))
	// Output:
	// abcdefghij
	// klmno
}

func ExampleReverse() {
	fmt.Println(str.Reverse("hello"))
	// Output: olleh
}

func ExampleReverse_pipe() {
	rev := str.Pipe(str.TrimSpace, str.Reverse)
	fmt.Println(rev("  hello  "))
	// Output: olleh
}

func ExampleCollapseWhitespace() {
	fmt.Println(str.CollapseWhitespace("hello   \t\n  world"))
	// Output: hello world
}

func ExampleTrimSpace() {
	fmt.Println(str.TrimSpace("  hello  "))
	// Output: hello
}

func ExampleSlugifyASCII() {
	fmt.Println(str.SlugifyASCII("Hello, World!"))
	// Output: hello-world
}

func ExampleSlugifyASCII_pipe() {
	slugify := str.Pipe(str.TrimSpace, str.CollapseWhitespace, str.SlugifyASCII)
	fmt.Println(slugify("  Hello,  World!  "))
	// Output: hello-world
}

func ExamplePipeIf() {
	isLong := func(s string) bool { return len([]rune(s)) > 10 }
	shorten := str.Pipe(
		str.TrimSpace,
		str.PipeIf(isLong, str.Truncate(10, "...")),
	)
	fmt.Println(shorten("  hello  "))
	fmt.Println(shorten("  this is a long string  "))
	// Output:
	// hello
	// this is...
}

func ExamplePipeUnless() {
	isEmpty := func(s string) bool { return s == "" }
	fn := str.PipeUnless(isEmpty, str.SlugifyASCII)
	fmt.Println(fn("Hello World"))
	fmt.Println(fn(""))
	// Output:
	// hello-world
	//
}

func ExamplePipeMap() {
	trimLines := str.PipeMap(str.Lines, "\n", str.TrimSpace)
	fmt.Println(trimLines("  hello  \n  world  "))
	// Output:
	// hello
	// world
}

func ExamplePipeMap_words() {
	shoutWords := str.PipeMap(str.Words, " ", str.ToScreamingSnake)
	fmt.Println(shoutWords("hello world"))
	// Output: HELLO WORLD
}

func ExampleLevenshtein() {
	fmt.Println(str.Levenshtein("kitten", "sitting"))
	// Output: 3
}

func ExampleJaroWinkler() {
	fmt.Printf("%.2f\n", str.JaroWinkler("hello", "hello"))
	// Output: 1.00
}
