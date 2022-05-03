package matchers

import (
	"fmt"
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
	return strings.Join(resultList, ";\n and ")
}

func (m allMatcher) Got(arg any) string {
	fallbackMessage := fmt.Sprintf("is %v (%T)", arg, arg)
	if m.submatchers == nil {
		return fallbackMessage
	}
	presenceList := make(map[string]bool)
	resultList := make([]string, 0)
	for _, subm := range m.submatchers {
		gottable, ok := subm.(gomock.GotFormatter)
		if !ok {
			continue
		}
		result := gottable.Got(arg)
		isPresent := presenceList[result]
		if isPresent {
			continue
		}
		presenceList[result] = true
		resultList = append(resultList, result)
	}
	if len(resultList) == 0 {
		return fallbackMessage
	}
	return strings.Join(resultList, ";\n")
}
