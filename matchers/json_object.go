package matchers

import "encoding/json"

type jsonObjectMatcher struct {
	expected map[string]any
}

func JSONObject(expected map[string]any) jsonObjectMatcher {
	return jsonObjectMatcher{
		expected: expected,
	}
}

func (m jsonObjectMatcher) Matches(x interface{}) bool {
	var actualmap map[string]any
	switch actual := x.(type) {
	case string:
		err := json.Unmarshal([]byte(actual), &actualmap)
		if err != nil {
			return false
		}
	case []byte:
		err := json.Unmarshal(actual, &actualmap)
		if err != nil {
			return false
		}
	default:
		return false
	}
	return matchMaps(m.expected, actualmap)
}

func (m jsonObjectMatcher) String() string {
	return "shoud be a valid JSON matching the specified map"
}
