package str

// Pipe composes multiple string transformers into a single function.
// Functions are applied left to right. With zero functions, Pipe returns
// an identity function that returns the input unchanged.
func Pipe(fns ...func(string) string) func(string) string {
	return func(s string) string {
		for _, fn := range fns {
			s = fn(s)
		}
		return s
	}
}

// PipeErr composes multiple fallible string transformers into a single function.
// Functions are applied left to right. Short-circuits on the first non-nil error;
// the error is returned unwrapped from the failing function. With zero functions,
// PipeErr returns an identity function that returns the input and nil.
func PipeErr(fns ...func(string) (string, error)) func(string) (string, error) {
	return func(s string) (string, error) {
		for _, fn := range fns {
			var err error
			s, err = fn(s)
			if err != nil {
				return s, err
			}
		}
		return s, nil
	}
}
