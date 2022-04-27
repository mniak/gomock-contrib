package matchers

type jsonMatcher struct {
	expected []byte
}

func JSON(j string) jsonMatcher {
	return jsonMatcher{
		expected: []byte(j),
	}
}

func (m jsonMatcher) Matches(x interface{}) bool {
	switch actual := x.(type) {
	case string:
		return matchBytes(m.expected, []byte(actual))
	case []byte:
		return matchBytes(m.expected, actual)
	default:
		return false
	}
}

func (m jsonMatcher) String() string {
	return "shoud be a valid JSON string"
}
