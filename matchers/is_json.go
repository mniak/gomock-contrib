package matchers

import (
	"encoding/json"

	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/utils"
)

// type jsonMatcher struct {
// 	expected map[string]any
// }

// func matchJSON(expected map[string]any) jsonMatcher {
// 	return jsonMatcher{
// 		expected: expected,
// 	}
// }

// func JSON(expected map[string]any) jsonMatcher {
// 	return matchJSON(expected)
// }

// func (m jsonMatcher) Matches(arg any) bool {
// 	var actualmap map[string]any
// 	switch actual := arg.(type) {
// 	case string:
// 		err := json.Unmarshal([]byte(actual), &actualmap)
// 		if err != nil {
// 			return false
// 		}
// 	case []byte:
// 		err := json.Unmarshal(actual, &actualmap)
// 		if err != nil {
// 			return false
// 		}
// 	default:
// 		return false
// 	}
// 	return utils.MatchMaps(m.expected, actualmap)
// }

// func (m jsonMatcher) String() string {
// 	pretty, err := json.MarshalIndent(m.expected, "", "  ")
// 	if err != nil {
// 		return fmt.Sprintf("shoud be JSON matching %+v", m.expected)
// 	}
// 	return fmt.Sprintf("shoud be JSON matching %s", pretty)
// }

// func (m jsonMatcher) Got(arg any) string {
// 	defaultMessage := fmt.Sprintf("%+v (%T)", arg, arg)
// 	var actualmap map[string]any
// 	switch actual := arg.(type) {
// 	case string:
// 		err := json.Unmarshal([]byte(actual), &actualmap)
// 		if err != nil {
// 			return defaultMessage
// 		}
// 	case []byte:
// 		err := json.Unmarshal(actual, &actualmap)
// 		if err != nil {
// 			return defaultMessage
// 		}
// 	default:
// 		return defaultMessage
// 	}

// 	pretty, err := json.MarshalIndent(actualmap, "", "  ")
// 	if err != nil {
// 		return defaultMessage
// 	}
// 	return string(pretty)
// }

type isJSONMatcher struct{}

func IsJSON() isJSONMatcher {
	return isJSONMatcher{}
}

func (m isJSONMatcher) Matches(arg any) bool {
	switch actual := arg.(type) {
	case string:
		return json.Valid([]byte(actual))
	case []byte:
		return json.Valid(actual)
	default:
		return false
	}
}

func (m isJSONMatcher) String() string {
	return ""
}

func (m isJSONMatcher) ThatMatches(matcher gomock.Matcher) isJSONThatMatchesMatcher {
	return isJSONThatMatchesMatcher{
		parent:     m,
		submatcher: matcher,
	}
}

type isJSONThatMatchesMatcher struct {
	parent     isJSONMatcher
	submatcher gomock.Matcher
}

func (m isJSONThatMatchesMatcher) Matches(arg any) bool {
	value, is := utils.MatchJSON[any](arg)
	if !is {
		return false
	}
	return m.submatcher.Matches(value)
}

func (m isJSONThatMatchesMatcher) String() string {
	return ""
}
