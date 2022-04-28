package matchers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mniak/gomock-contrib/internal/utils"
)

type likeMapMatcher struct {
	expectedMap map[string]any
}

func LikeMap(expected map[string]any) likeMapMatcher {
	return likeMapMatcher{
		expectedMap: expected,
	}
}

func (m likeMapMatcher) Matches(arg any) bool {
	expectedValue := reflect.ValueOf(m.expectedMap)
	actualValue := reflect.ValueOf(arg)
	return utils.MatchValues(expectedValue, actualValue)
}

func (m likeMapMatcher) String() string {
	pretty := utils.PrettyPrintMap(m.expectedMap)
	return fmt.Sprintf("shoud match %s", pretty)
}

func (m likeMapMatcher) Got(arg any) string {
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
