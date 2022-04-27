package typedmatchers

import "github.com/golang/mock/gomock"

func FieldEqual[T any, F any](fieldSelector func(x T) F, expectedValue F) fieldMatcher[T, any] {
	return FieldGeneric(func(x T) any {
		return fieldSelector(x)
	}, gomock.Eq(expectedValue))
}

func FieldInlineJSON[T any, J any](fieldSelector func(x T) string, function func(x J) bool) any {
	return Field[T, string](fieldSelector, InlineJSON(function))
}
