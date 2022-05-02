package utils

import "reflect"

func UnwrapValue(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}

func tryGetValue[T any](value reflect.Value) (T, bool) {
	if value.IsValid() && value.CanInterface() {
		interf := value.Interface()
		t, ok := interf.(T)
		return t, ok
	}
	var t T
	return t, false
}
