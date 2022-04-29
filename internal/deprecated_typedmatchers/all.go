package typedmatchers

import "strings"

type recordedFailure struct {
	want string
	got  string
}

type allMatcher[T any] struct {
	matchers         []Matcher[T]
	recordedFailures []recordedFailure
}

func All[T any](matchers ...Matcher[T]) *allMatcher[T] {
	return &allMatcher[T]{
		matchers: matchers,
	}
}

func (m *allMatcher[T]) Matches(x T) bool {
	result := true
	m.recordedFailures = make([]recordedFailure, 0)
	for _, matcher := range m.matchers {
		if !matcher.Matches(x) {
			m.recordedFailures = append(m.recordedFailures, recordedFailure{
				want: m.String(),
				got:  formatGottenArg(matcher, x),
			})
			result = false
		}
	}
	return result
}

func (m *allMatcher[T]) String() string {
	if m.recordedFailures == nil || len(m.recordedFailures) == 0 {

		ss := make([]string, 0, len(m.matchers))
		for _, matcher := range m.matchers {
			ss = append(ss, matcher.String())
		}
		return strings.Join(ss, "; ")

	} else {

		failedWants := make([]string, 0, len(m.recordedFailures))
		for _, failure := range m.recordedFailures {
			failedWants = append(failedWants, failure.want)
		}
		return strings.Join(failedWants, "\n")

	}
}

func (m *allMatcher[T]) Got(got T) string {
	if m.recordedFailures == nil || len(m.recordedFailures) == 0 {

		ss := make([]string, 0, len(m.matchers))
		for _, matcher := range m.matchers {
			ss = append(ss, formatGottenArg(matcher, got))
		}
		return strings.Join(ss, "\n")

	} else {

		failedGots := make([]string, 0, len(m.recordedFailures))
		for _, failure := range m.recordedFailures {
			failedGots = append(failedGots, failure.got)
		}
		return strings.Join(failedGots, "\n")

	}
}
