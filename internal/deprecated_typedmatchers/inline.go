package typedmatchers

type inlineMatcher[T any] struct {
	message  string
	function func(x T) bool
}

func Inline[T any](message string, function func(x T) bool) inlineMatcher[T] {
	return inlineMatcher[T]{
		message:  message,
		function: function,
	}
}

func (m inlineMatcher[T]) Matches(x T) bool {
	return m.function(x)
}

func (m inlineMatcher[T]) String() string {
	return m.message
}
