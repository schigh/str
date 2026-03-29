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

// PipeIf returns a transformer that applies fn only when the predicate returns true.
// If the predicate returns false, the input passes through unchanged.
// Nil predicate is treated as always-false (no-op). Nil fn is a no-op.
func PipeIf(predicate func(string) bool, fn func(string) string) func(string) string {
	return func(s string) string {
		if predicate == nil || fn == nil {
			return s
		}
		if predicate(s) {
			return fn(s)
		}
		return s
	}
}

// PipeUnless returns a transformer that applies fn only when the predicate returns false.
// If the predicate returns true, the input passes through unchanged.
// Nil predicate is treated as always-false, so fn always applies. Nil fn is a no-op.
func PipeUnless(predicate func(string) bool, fn func(string) string) func(string) string {
	return func(s string) string {
		if fn == nil {
			return s
		}
		if predicate == nil || !predicate(s) {
			return fn(s)
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
