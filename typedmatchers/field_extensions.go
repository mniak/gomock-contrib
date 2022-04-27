package typedmatchers

import "github.com/golang/mock/gomock"

func FieldEqual[T any, F any](fieldName string, expectedValue F) fieldMatcher[T, any] {
	return FieldGeneric[T](fieldName, gomock.Eq(expectedValue))
}

func FieldInlineJSON[T any, J any](fieldName string, function func(x J) bool) any {
	return Field[T, string](fieldName, InlineJSON(function))
}
