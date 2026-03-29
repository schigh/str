package str

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"already snake", "already_snake", "already_snake"},
		{"camelCase", "helloWorld", "hello_world"},
		{"PascalCase", "HelloWorld", "hello_world"},
		{"acronym followed by word", "HTMLParser", "html_parser"},
		{"multiple acronyms", "getHTTPSURL", "get_https_url"},
		{"digit boundary", "Get2ndItem", "get_2nd_item"},
		{"mixed delimiters", "Mixed-kebab_snake", "mixed_kebab_snake"},
		{"acronym at start", "IDField", "id_field"},
		{"acronym at end", "UserID", "user_id"},
		{"multiple acronyms in sentence", "getURLForAPI", "get_url_for_api"},
		{"all uppercase", "ALLCAPS", "allcaps"},
		{"all lowercase", "alllower", "alllower"},
		{"single char", "A", "a"},
		{"unicode letters", "GroßEltern", "groß_eltern"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSnakeCase(tt.input)
			if got != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"empty", "", ""},
		{"from snake", "hello_world", "helloWorld"},
		{"from pascal", "HelloWorld", "helloWorld"},
		{"acronym", "html_parser", "htmlParser"},
		{"single word", "hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCamelCase(tt.input)
			if got != tt.expected {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"empty", "", ""},
		{"from snake", "hello_world", "HelloWorld"},
		{"from camel", "helloWorld", "HelloWorld"},
		{"single word", "hello", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToPascalCase(tt.input)
			if got != tt.expected {
				t.Errorf("ToPascalCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"empty", "", ""},
		{"from camel", "helloWorld", "hello-world"},
		{"from snake", "hello_world", "hello-world"},
		{"acronym", "HTMLParser", "html-parser"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToKebabCase(tt.input)
			if got != tt.expected {
				t.Errorf("ToKebabCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"empty", "", ""},
		{"from snake", "hello_world", "Hello World"},
		{"from camel", "helloWorld", "Hello World"},
		{"single word", "hello", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToTitleCase(tt.input)
			if got != tt.expected {
				t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToScreamingSnake(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"empty", "", ""},
		{"from camel", "helloWorld", "HELLO_WORLD"},
		{"from snake", "hello_world", "HELLO_WORLD"},
		{"acronym", "HTMLParser", "HTML_PARSER"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToScreamingSnake(tt.input)
			if got != tt.expected {
				t.Errorf("ToScreamingSnake(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func FuzzSplitWords(f *testing.F) {
	f.Add("HTMLParser")
	f.Add("getHTTPSURL")
	f.Add("Get2ndItem")
	f.Add("already_snake")
	f.Add("Mixed-kebab_snake")
	f.Add("IDField")
	f.Add("UserID")
	f.Add("")
	f.Add("GroßEltern")
	f.Fuzz(func(t *testing.T, s string) {
		// Should never panic.
		_ = splitWords(s)
	})
}
