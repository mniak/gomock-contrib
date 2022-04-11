package matchers

type inlineMatcher struct {
	message  string
	function func(x any) bool
}

func Inline(message string, function func(x any) bool) *inlineMatcher {
	return &inlineMatcher{
		message:  message,
		function: function,
	}
}

func (m *inlineMatcher) Matches(x any) bool {
	return m.function(x)
}

func (m *inlineMatcher) String() string {
	return m.message
}
