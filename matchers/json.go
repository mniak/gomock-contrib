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
		if !matchValues(expectedReflectionValue, actualReflectionValue) {
			return false
		}
	}
	return true
}

func matchValues(expected, actual reflect.Value) bool {
	if expected.Type() != actual.Type() {
		return false
	}
	switch actual.Kind() {
	case reflect.Map:
		for _, idx := range expected.MapKeys() {
			expectedValue := expected.MapIndex(idx)
			actualValue := actual.MapIndex(idx)
			if !matchValues(expectedValue, actualValue) {
				return false
			}
		}
		return true
	case reflect.Slice:
		if actual.Len() != expected.Len() {
			return false
		}

		for idx := 0; idx < actual.Len(); idx++ {
			expectedValue := expected.Index(idx)
			actualValue := actual.Index(idx)
			if !matchValues(expectedValue, actualValue) {
				return false
			}
		}
		return true
	}
	return actual.Interface() == expected.Interface()
}

func (m jsonMatcher) String() string {
	return "shoud be a valid JSON string"
}
