package typedmatchers

import (
	"encoding/json"

	"github.com/mniak/gomock-contrib/internal/utils"
)

type jsonMatcher[T binary] struct {
	expected map[string]any
}

func matchJSON[T binary](expected map[string]any) jsonMatcher[T] {
	return jsonMatcher[T]{
		expected: expected,
	}
}

func (m jsonMatcher[T]) Matches(actual T) bool {
	var actualmap map[string]any
	err := json.Unmarshal([]byte(actual), &actualmap)
	if err != nil {
		return false
	}
	return utils.MatchMaps(m.expected, actualmap)
}

func (m jsonMatcher[T]) String() string {
	return "shoud be a valid JSON matching the specified map"
}

func JSON(expected map[string]any) jsonMatcher[string] {
	return matchJSON[string](expected)
}

func BinaryJSON(expected map[string]any) jsonMatcher[[]byte] {
	return matchJSON[[]byte](expected)
}
