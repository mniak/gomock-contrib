package typedmatchers

var (
	_ Matcher[string] = &noopMatcher[string]{}
	_ Matcher[int32]  = &noopMatcher[int32]{}
)
