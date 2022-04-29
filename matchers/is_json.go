package matchers

import (
	"encoding/json"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/utils"
)

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

func (m isJSONMatcher) Got(arg any) string {
	switch actual := arg.(type) {
	case string:
		return fmt.Sprintf("is %v (%T)", actual, actual)
	case []byte:
		return fmt.Sprintf("is %v (%T)", string(actual), actual)
	default:
		return fmt.Sprintf("data with invalid type: %v (%T)", arg, arg)
	}
}

func (m isJSONMatcher) String() string {
	return "is valid JSON"
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
	return fmt.Sprintf("is a valid JSON that %s", m.submatcher.String())
}

func (m isJSONThatMatchesMatcher) Got(arg any) string {
	if !m.parent.Matches(arg) {
		return m.parent.Got(arg)
	}
	if gf, is := m.submatcher.(gomock.GotFormatter); is {
		return gf.Got(arg)
	}
	return fmt.Sprintf("is %v (%T)", arg, arg)
}
