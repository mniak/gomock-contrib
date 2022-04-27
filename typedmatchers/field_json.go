package typedmatchers

func FieldInlineJSON[T any, J any](fieldSelector func(x T) string, function func(x J) bool) any {
	return MatchField[T, string](fieldSelector, InlineJSON(function))
}
