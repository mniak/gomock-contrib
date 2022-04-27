package typedmatchers

import "fmt"

type fieldMatcher[T any, F any] struct {
	fieldName string
	selector  func(x T) F
	matcher   Matcher[F]
}

func Field[T any, F any](fieldName string, fieldSelector func(x T) F, matcher Matcher[F]) fieldMatcher[T, F] {
	return fieldMatcher[T, F]{
		fieldName: fieldName,
		selector:  fieldSelector,
		matcher:   matcher,
	}
}

func FieldGeneric[T any](fieldSelector func(x T) any, matcher Matcher[any]) fieldMatcher[T, any] {
	return fieldMatcher[T, any]{
		selector: fieldSelector,
		matcher:  matcher,
	}
}

func (m fieldMatcher[T, F]) Matches(x T) bool {
	fieldValue := m.selector(x)
	return m.matcher.Matches(fieldValue)
}

func (m fieldMatcher[T, F]) String() string {
	return fmt.Sprintf("field %s %s", m.fieldName, m.matcher.String())
}

func (m fieldMatcher[T, F]) Got(actual T) string {
	fieldValue := m.selector(actual)
	return fmt.Sprintf("field %s: %s", m.fieldName, formatGottenArg(m.matcher, fieldValue))
}
