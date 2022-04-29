package typedmatchers

var (
	_ Matcher[string]      = All[string]()
	_ Matcher[bool]        = All[bool]()
	_ GotFormatter[string] = All[string]()
	_ GotFormatter[bool]   = All[bool]()
)
