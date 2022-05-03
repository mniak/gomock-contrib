package matchers

import (
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
	pretty := utils.PrettyPrint(m.expectedMap)
	return fmt.Sprintf("matches %s", pretty)
}

func (m likeMapMatcher) Got(arg any) string {
	if arg == nil {
		return "is nil"
	}
	defaultMessage := fmt.Sprintf("is %+v (%T)", arg, arg)

	asMap, ok := arg.(map[string]any)
	if !ok {
		return defaultMessage
	}

	pretty := utils.PrettyPrint(asMap)
	return fmt.Sprintf("is %s", pretty)
}
