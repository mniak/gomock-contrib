package typedmatchers

// A Matcher is a representation of a class of values.
// It is used to represent the valid or expected arguments to a mocked method.
type Matcher[T any] interface {
	// Matches returns whether x is a match.
	Matches(x T) bool

	// String describes what the matcher matches.
	String() string
}
