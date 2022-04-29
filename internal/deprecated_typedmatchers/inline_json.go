package typedmatchers

import "encoding/json"

type inlineJsonMatcher[T any] struct {
	function func(x T) bool
}

func InlineJSON[T any](function func(x T) bool) inlineJsonMatcher[T] {
	return inlineJsonMatcher[T]{
		function: function,
	}
}

func (m inlineJsonMatcher[T]) Matches(x string) bool {
	var obj T
	err := json.Unmarshal([]byte(x), &obj)
	if err != nil {
		return false
	}
	if m.function == nil {
		return true
	}
	return m.function(obj)
}

func (m inlineJsonMatcher[T]) String() string {
	return "should be a valid JSON string"
}
