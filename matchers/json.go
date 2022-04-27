package matchers

import (
	"encoding/json"
	"reflect"
)

type jsonMatcher struct {
	expected []byte
}

func JSON(j string) jsonMatcher {
	return jsonMatcher{
		expected: []byte(j),
	}
}

func (m jsonMatcher) Matches(x interface{}) bool {
	switch val := x.(type) {
	case string:
		return matchBytes(m.expected, []byte(val))
	case []byte:
		return matchBytes(m.expected, val)
	default:
		return false
	}
}

func matchBytes(expected, data []byte) bool {
	var expectedmap map[string]any
	err := json.Unmarshal(expected, &expectedmap)
	if err != nil {
		return false
	}

	var datamap map[string]any
	err = json.Unmarshal(data, &datamap)
	if err != nil {
		return false
	}
	return matchMaps(expectedmap, datamap)
}

func matchMaps(expectedmap, actualMap map[string]any) bool {
	for key, expectedValue := range expectedmap {
		actualValue, found := actualMap[key]
		if !found {
			return false
		}
		expectedReflectionValue := reflect.ValueOf(expectedValue)
		actualReflectionValue := reflect.ValueOf(actualValue)
		if expectedReflectionValue.Type() != actualReflectionValue.Type() {
			return false
		}
		switch actualReflectionValue.Kind() {
		case reflect.Map:
			expectedValueAsMap, expectedIsMap := expectedValue.(map[string]any)
			actualValueAsMap, actualIsMap := actualValue.(map[string]any)
			if expectedIsMap && actualIsMap {
				return matchMaps(expectedValueAsMap, actualValueAsMap)
			}

			return actualReflectionValue.Elem() == expectedReflectionValue.Elem()
		}
		if actualValue != expectedValue {
			return false
		}
	}
	return true
}

func (m jsonMatcher) String() string {
	return "shoud be a valid JSON string"
}
