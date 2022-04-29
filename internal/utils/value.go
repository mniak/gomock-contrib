package utils

import "reflect"

func TryGetValue[T any](value reflect.Value) (T, bool) {
	value = UnwrapValue(value)
	if value.CanInterface() {
		intf := value.Interface()
		val, ok := intf.(T)
		return val, ok
	}
	var def T
	return def, false
}
