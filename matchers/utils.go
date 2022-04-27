package matchers

import (
	"encoding/json"
	"reflect"
)

func matchBytes(expected, data []byte) bool {
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
	return matchMaps(expectedmap, actualmap)
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

func matchValues(expected, actual reflect.Value) bool {
	for expected.Kind() == reflect.Pointer || expected.Kind() == reflect.Interface {
		expected = expected.Elem()
	}
	for actual.Kind() == reflect.Pointer || actual.Kind() == reflect.Interface {
		actual = actual.Elem()
	}

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
