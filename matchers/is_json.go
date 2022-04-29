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

func (m isJSONMatcher) String() string {
	return "is valid JSON"
}

func (m isJSONMatcher) ThatMatches(matcher any) isJSONThatMatchesMatcher {
	return isJSONThatMatchesMatcher{
		parent:     m,
		submatcher: utils.ArgAsMatcher(matcher),
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
	return fmt.Sprintf("should be a valid JSON that %s", m.submatcher.String())
}

func (m isJSONThatMatchesMatcher) Got(arg any) string {
	return m.submatcher.String()
}
