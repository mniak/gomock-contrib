package typedmatchers

import (
	"encoding/json"
	"fmt"

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

func JSON(expected map[string]any) jsonMatcher[string] {
	return matchJSON[string](expected)
}

func BinaryJSON(expected map[string]any) jsonMatcher[[]byte] {
	return matchJSON[[]byte](expected)
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
	pretty, err := json.MarshalIndent(m.expected, "", "  ")
	if err != nil {
		return fmt.Sprintf("shoud be JSON matching %+v", m.expected)
	}
	return fmt.Sprintf("shoud be JSON matching %s", pretty)
}

func (m jsonMatcher[T]) Got(actual T) string {
	var actualmap map[string]any
	err := json.Unmarshal([]byte(actual), &actualmap)
	if err != nil {
		return fmt.Sprintf("%+v (%T)", actual, actual)
	}
	pretty, err := json.MarshalIndent(actualmap, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v (%T)", actual, actual)
	}
	return fmt.Sprintf("shoud be JSON matching %s", pretty)
}
