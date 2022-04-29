package matchers

import (
	"encoding/json"

	"github.com/mniak/gomock-contrib/internal/utils"
)

type jsonStringMatcher struct {
	expected map[string]any
}

func JSONString(j string) jsonStringMatcher {
	var expectedMap map[string]any
	err := json.Unmarshal([]byte(j), &expectedMap)
	if err != nil {
		panic("the string provided is not a valid JSON")
	}

	return jsonStringMatcher{
		expected: expectedMap,
	}
}

func (m jsonStringMatcher) Matches(x any) bool {
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
	return utils.MatchMaps(m.expected, actualmap)
}

func (m jsonStringMatcher) String() string {
	return "should be a valid JSON matching the specified map"
}
