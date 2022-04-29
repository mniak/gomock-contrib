package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const prettyPrintIndentation = "\t"

func PrettyPrintMap(m map[string]any) string {
	value := reflect.ValueOf(m)
	return internalPrettyPrint(value, "")
}

func internalPrettyPrintType(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Pointer {
		return "*" + internalPrettyPrintType(rtype.Elem())
	}

	k := rtype.Kind()
	switch k {
	case reflect.Map:
		keyStr := internalPrettyPrintType(rtype.Key())
		valueStr := internalPrettyPrintType(rtype.Elem())
		return fmt.Sprintf("map[%s]%s", keyStr, valueStr)

	case reflect.Interface:
		name := rtype.Name()
		if name != "" {
			return name
		} else {
			return "any"
		}
	case reflect.String:
		return rtype.Name()
	default:
		return fmt.Sprintf("(could not format type %+v)", rtype)
	}
}

func valueIsNillable(value reflect.Value) bool {
	k := value.Kind()
	switch k {
	case reflect.Chan:
	case reflect.Func:
	case reflect.Map:
	case reflect.Pointer, reflect.UnsafePointer:
	case reflect.Interface:
	case reflect.Slice:
	default:
		return false
	}
	return true
}

func internalPrettyPrint(value reflect.Value, newlinePrefix string) string {
	if valueIsNillable(value) && value.IsNil() {
		return "nil"
	}
	kind := value.Kind()
	switch kind {
	case reflect.Map:
		mapTypeStr := internalPrettyPrintType(value.Type())
		mapKeys := value.MapKeys()

		if len(mapKeys) == 0 {
			return fmt.Sprintf("%s{}", mapTypeStr)
		}

		lines := make([]string, 0)
		for _, mapKey := range mapKeys {
			mapValue := value.MapIndex(mapKey)

			keyStr := internalPrettyPrint(mapKey, newlinePrefix+prettyPrintIndentation)
			valueStr := internalPrettyPrint(mapValue, newlinePrefix+prettyPrintIndentation)

			lines = append(lines, fmt.Sprintf("%s%s: %s,\n%s",
				prettyPrintIndentation,
				keyStr,
				valueStr,
				newlinePrefix,
			))
		}
		slices.Sort(lines)
		return fmt.Sprintf("%s{\n%s%s}", mapTypeStr, newlinePrefix, strings.Join(lines, ""))

	case reflect.Slice:
		sliceTypeStr := internalPrettyPrintType(value.Type().Elem())

		length := value.Len()
		if length == 0 {
			return fmt.Sprintf("[]%s{}", sliceTypeStr)
		}

		lines := make([]string, length)
		for idx := 0; idx < length; idx++ {
			value := value.Index(idx)
			valueStr := internalPrettyPrint(value, newlinePrefix+prettyPrintIndentation)

			lines[idx] = fmt.Sprintf("%s%s,\n%s",
				prettyPrintIndentation,
				valueStr,
				newlinePrefix,
			)
		}
		return fmt.Sprintf("[]%s{\n%s%s}", sliceTypeStr, newlinePrefix, strings.Join(lines, ""))
	case reflect.String:
		return strconv.Quote(value.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.Itoa(int(value.Int()))
	// case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	// 	fallthrough
	// case reflect.Float32, reflect.Float64:
	// 	fallthrough
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	// case reflect.UnsafePointer:
	// case reflect.Struct:
	// case reflect.Complex64:
	// case reflect.Complex128:
	// case reflect.Array:
	// case reflect.Chan:
	// case reflect.Func:
	case reflect.Interface:
		return internalPrettyPrint(value.Elem(), newlinePrefix)
	case reflect.Pointer:
		elemPretty := internalPrettyPrint(value.Elem(), newlinePrefix)
		return fmt.Sprintf("&%s", elemPretty)
	default:
		v := value.Interface()
		return fmt.Sprintf("%v (%T)", v, v)
	}
}
