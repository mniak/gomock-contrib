package matchers

func (m isJSONMatcher) ThatMatchesMap(expectedMap map[string]any) isJSONThatMatchesMatcher {
	return m.ThatMatches(LikeMap(expectedMap))
}
