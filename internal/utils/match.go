package utils

import (
	"encoding/json"
	"reflect"
)

func MatchBytes(expected, data []byte) bool {
	var expectedmap map[string]any
	err := json.Unmarshal(expected, &expectedmap)
	if err != nil {
		return false
	}

	var actualmap map[string]any
	err = json.Unmarshal(data, &actualmap)
	if err != nil {
		return false
	}
	return MatchMaps(expectedmap, actualmap)
}

func MatchMaps(expectedmap, actualMap map[string]any) bool {
	for key, expectedValue := range expectedmap {
		actualValue, found := actualMap[key]
		if !found {
			return false
		}
		expectedReflectionValue := reflect.ValueOf(expectedValue)
		actualReflectionValue := reflect.ValueOf(actualValue)
		if !MatchValues(expectedReflectionValue, actualReflectionValue) {
			return false
		}
	}
	return true
}

func asFloat(value reflect.Value) (float64, bool) {
	if value.CanFloat() {
		return value.Float(), true
	} else if value.CanInt() {
		return float64(value.Int()), true
	} else if value.CanUint() {
		return float64(value.Uint()), true
	}

	return 0, false
}

func MatchValues(expected, actual reflect.Value) bool {
	expected = UnwrapValue(expected)
	actual = UnwrapValue(actual)

	actualAsFloat, actualIsFloat := asFloat(actual)
	expectedAsFloat, expectedIsFloat := asFloat(expected)
	if actualIsFloat && expectedIsFloat {
		return actualAsFloat == expectedAsFloat
	}

	if expected.Kind() != actual.Kind() {
		return false
	}
	switch actual.Kind() {
	case reflect.Map:
		for _, idx := range expected.MapKeys() {
			expectedValue := expected.MapIndex(idx)
			actualValue := actual.MapIndex(idx)
			if !MatchValues(expectedValue, actualValue) {
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
			if !MatchValues(expectedValue, actualValue) {
				return false
			}
		}
		return true
	}
	return actual.Interface() == expected.Interface()
}

func MatchJSON[T any](arg any) (T, bool) {
	var result T
	switch actual := arg.(type) {
	case string:
		err := json.Unmarshal([]byte(actual), &result)
		if err != nil {
			return result, false
		}
	case []byte:
		err := json.Unmarshal(actual, &result)
		if err != nil {
			return result, false
		}
	default:
		return result, false
	}
	return result, true
}
