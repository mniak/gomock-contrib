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

func (m typedMatcher[T]) Matches(x any) bool {
	switch val := x.(type) {
	case T:
		return m.typedMatcher.Matches(val)
	case *T:
		return m.typedMatcher.Matches(*val)
	default:
		return false
	}
}

func (m typedMatcher[T]) String() string {
	if m.message != "" {
		return m.message
	}
	return m.typedMatcher.String()
}

func (m typedMatcher[T]) Got(actual any) string {
	switch actual := actual.(type) {
	case T:
		if gs, ok := m.typedMatcher.(typedmatchers.GotFormatter[T]); ok {
			return gs.Got(actual)
		}
	case *T:
		if gs, ok := m.typedMatcher.(typedmatchers.GotFormatter[T]); ok {
			return gs.Got(*actual)
		}
	}
	return fmt.Sprintf("%+v (%T)", actual, actual)
}
