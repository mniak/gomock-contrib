package matchers

import (
	"fmt"

	"github.com/mniak/gomock-contrib/typedmatchers"
)

type typedMatcher[T any] struct {
	message      string
	typedMatcher typedmatchers.Matcher[T]
}

func Typed[T any](matchers ...typedmatchers.Matcher[T]) typedMatcher[T] {
	if len(matchers) == 0 {
		message := fmt.Sprintf("expecting value of type %s", "")
		return typedMatcher[T]{
			message:      message,
			typedMatcher: typedmatchers.Noop[T](),
		}
	} else if len(matchers) == 1 {
		return typedMatcher[T]{
			typedMatcher: matchers[0],
		}
	} else {
		matcher := typedmatchers.All(matchers...)
		return typedMatcher[T]{
			typedMatcher: matcher,
		}
	}
}

func (m *typedMatcher[T]) Matches(x interface{}) bool {
	val, is := x.(T)
	if !is {
		return false
	}
	return m.typedMatcher.Matches(val)
}

func (m *typedMatcher[T]) String() string {
	if m.message != "" {
		return m.message
	}
	return m.typedMatcher.String()
}
