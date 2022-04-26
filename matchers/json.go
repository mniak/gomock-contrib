package matchers

import "encoding/json"

type jsonMatcher struct {
	expected []byte
}

func JSON(j string) jsonMatcher {
	return jsonMatcher{
		expected: []byte(j),
	}
}

func (m jsonMatcher) Matches(x interface{}) bool {
	switch val := x.(type) {
	case string:
		return matchBytes(m.expected, []byte(val))
	case []byte:
		return matchBytes(m.expected, val)
	default:
		return false
	}
}

func matchBytes(expected, data []byte) bool {
	var expectedmap map[string]any
	err := json.Unmarshal(expected, &expectedmap)
	if err != nil {
		return false
	}

	var datamap map[string]any
	err = json.Unmarshal(data, &datamap)
	if err != nil {
		return false
	}

	for ek, ev := range expectedmap {
		dv, has := datamap[ek]
		if !has ||
			ev == nil && dv != nil ||
			dv == nil && ev != nil {
			return false
		}
	}
	return true
}

func (m jsonMatcher) String() string {
	return "shoud be a valid JSON string"
}
