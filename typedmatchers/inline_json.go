package typedmatchers

import "encoding/json"

type jsonMatcher[T any] struct {
	function func(x T) bool
}

func InlineJSON[T any](function func(x T) bool) jsonMatcher[T] {
	return jsonMatcher[T]{
		function: function,
	}
}

func (m jsonMatcher[T]) Matches(x string) bool {
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

func (m jsonMatcher[T]) String() string {
	return ""
}
