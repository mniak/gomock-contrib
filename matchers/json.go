package matchers

import (
	"encoding/json"
	"fmt"

	"github.com/mniak/gomock-contrib/internal/utils"
)

type jsonMatcher struct {
	expected map[string]any
}

func matchJSON(expected map[string]any) jsonMatcher {
	return jsonMatcher{
		expected: expected,
	}
}

func JSON(expected map[string]any) jsonMatcher {
	return matchJSON(expected)
}

func (m jsonMatcher) Matches(arg any) bool {
	var actualmap map[string]any
	switch actual := arg.(type) {
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

func (m jsonMatcher) String() string {
	pretty, err := json.MarshalIndent(m.expected, "", "  ")
	if err != nil {
		return fmt.Sprintf("shoud be JSON matching %+v", m.expected)
	}
	return fmt.Sprintf("shoud be JSON matching %s", pretty)
}

func (m jsonMatcher) Got(arg any) string {
	defaultMessage := fmt.Sprintf("%+v (%T)", arg, arg)
	var actualmap map[string]any
	switch actual := arg.(type) {
	case string:
		err := json.Unmarshal([]byte(actual), &actualmap)
		if err != nil {
			return defaultMessage
		}
	case []byte:
		err := json.Unmarshal(actual, &actualmap)
		if err != nil {
			return defaultMessage
		}
	default:
		return defaultMessage
	}

	pretty, err := json.MarshalIndent(actualmap, "", "  ")
	if err != nil {
		return defaultMessage
	}
	return string(pretty)
}

type isJSONObjectMatcher struct{}

func IsJSONObject() isJSONObjectMatcher {
	return isJSONObjectMatcher{}
}

func (m isJSONObjectMatcher) internalMatches(arg interface{}) (map[string]any, bool) {
	var actualmap map[string]any
	switch actual := arg.(type) {
	case string:
		err := json.Unmarshal([]byte(actual), &actualmap)
		if err != nil {
			return nil, false
		}
	case []byte:
		err := json.Unmarshal(actual, &actualmap)
		if err != nil {
			return nil, false
		}
	default:
		return nil, false
	}
	return actualmap, true
}

func (m isJSONObjectMatcher) Matches(arg interface{}) bool {
	_, is := m.internalMatches(arg)
	return is
}

func (m isJSONObjectMatcher) String() string {
	return ""
}
