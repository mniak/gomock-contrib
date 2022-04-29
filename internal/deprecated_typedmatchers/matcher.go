package typedmatchers

// A Matcher is a representation of a class of values.
// It is used to represent the valid or expected arguments to a mocked method.
type Matcher[T any] interface {
	// Matches returns whether x is a match.
	Matches(x T) bool

	// String describes what the matcher matches.
	String() string
}

// GotFormatter is used to better print failure messages. If a matcher
// implements GotFormatter, it will use the result from Got when printing
// the failure message.
type GotFormatter[T any] interface {
	// Got is invoked with the received value. The result is used when
	// printing the failure message.
	Got(got T) string
}
