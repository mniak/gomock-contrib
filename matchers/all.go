package matchers

import (
	"strings"

	"github.com/golang/mock/gomock"
)

type allMatcher struct {
	submatchers []gomock.Matcher
}

func All(submatchers ...gomock.Matcher) allMatcher {
	return allMatcher{
		submatchers: submatchers,
	}
}

func (m allMatcher) Matches(arg any) bool {
	if m.submatchers == nil {
		return true
	}
	for _, subm := range m.submatchers {
		res := subm.Matches(arg)
		if !res {
			return false
		}
	}
	return true
}

func (m allMatcher) String() string {
	if m.submatchers == nil {
		return "anything"
	}
	resultList := make([]string, len(m.submatchers))
	for idx, subm := range m.submatchers {
		resultList[idx] = subm.String()
	}
	return strings.Join(resultList, ";\n")
}

func (m allMatcher) Got(got any) string {
	return ""
}
