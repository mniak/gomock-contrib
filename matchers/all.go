package matchers

import "github.com/golang/mock/gomock"

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
	return ""
}

func (m allMatcher) Got(got any) string {
	return ""
}
