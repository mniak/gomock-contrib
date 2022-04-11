package typedmatchers

import "strings"

func All[T any](matchers ...Matcher[T]) Matcher[T] {
	return allMatcher[T]{
		matchers: matchers,
	}
}

type allMatcher[T any] struct {
	matchers []Matcher[T]
}

func (am allMatcher[T]) Matches(x T) bool {
	for _, m := range am.matchers {
		if !m.Matches(x) {
			return false
		}
	}
	return true
}

func (am allMatcher[T]) String() string {
	ss := make([]string, 0, len(am.matchers))
	for _, matcher := range am.matchers {
		ss = append(ss, matcher.String())
	}
	return strings.Join(ss, "; ")
}
