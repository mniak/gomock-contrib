package matchers

import (
	"encoding/json"

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
