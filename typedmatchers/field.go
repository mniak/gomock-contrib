package typedmatchers

type fieldMatcher[T any, F any] struct {
	selector func(x T) F
	matcher  Matcher[F]
}

func MatchField[T any, F any](fieldSelector func(x T) F, matcher Matcher[F]) fieldMatcher[T, F] {
	return fieldMatcher[T, F]{
		selector: fieldSelector,
		matcher:  matcher,
	}
}

func (m fieldMatcher[T, F]) Matches(x T) bool {
	fieldValue := m.selector(x)
	return m.matcher.Matches(fieldValue)
}

func (m fieldMatcher[T, F]) String() string {
	return m.String()
}

// func FieldEqual[T any, F any](fieldSelector func(x T) F, expectedValue F) fieldMatcher[T, F] {
// 	return MatchField(fieldSelector, gomock.Eq)
// }
