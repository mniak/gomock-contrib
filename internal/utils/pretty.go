package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/mock/gomock"
	"golang.org/x/exp/slices"
)

const prettyPrintIndentation = "\t"

func PrettyPrintMap(m map[string]any) string {
	value := reflect.ValueOf(m)
	return internalPrettyPrint(value, "")
}

var gomockMatcherType = reflect.TypeOf((*gomock.Matcher)(nil)).Elem()

func internalPrettyPrintType(rtype reflect.Type, newlinePrefix string) string {
	typename := fmt.Sprint(rtype)
	_ = typename
	if rtype.Implements(gomockMatcherType) {
		fmt.Println("")
	}
	if rtype.AssignableTo(gomockMatcherType) {
		fmt.Println("")
	}
	if rtype.ConvertibleTo(gomockMatcherType) {
		fmt.Println("")
	}
	if rtype.Kind() == reflect.Pointer {
		return "*" + internalPrettyPrintType(rtype.Elem(), newlinePrefix)
	}

	switch rtype.Kind() {
	case reflect.Interface:
		name := rtype.Name()
		if name != "" {
			return name
		} else {
			return "any"
		}

	case reflect.Map:
		keyStr := internalPrettyPrintType(rtype.Key(), newlinePrefix)
		valueStr := internalPrettyPrintType(rtype.Elem(), newlinePrefix)
		return fmt.Sprintf("map[%s]%s", keyStr, valueStr)

	case reflect.Struct:
		numFields := rtype.NumField()
		if numFields == 0 {
			return "struct {}"
		}

		lines := make([]string, numFields)
		for idx := 0; idx < numFields; idx++ {
			field := rtype.FieldByIndex([]int{idx})
			fieldNameStr := field.Name
			fieldTypeStr := internalPrettyPrintType(field.Type, newlinePrefix+prettyPrintIndentation)

			lines[idx] = fmt.Sprintf("%s%s %s\n%s",
				prettyPrintIndentation,
				fieldNameStr,
				fieldTypeStr,
				newlinePrefix,
			)
		}
		slices.Sort(lines)
		return fmt.Sprintf("struct {\n%s%s}", newlinePrefix, strings.Join(lines, ""))

	default:
		return rtype.Name()
	}
}

func valueIsNillable(value reflect.Value) bool {
	switch value.Kind() {
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
	vk := value.Kind()
	switch vk {
	case reflect.Map:
		mapTypeStr := internalPrettyPrintType(value.Type(), newlinePrefix)
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
		sliceTypeStr := internalPrettyPrintType(value.Type().Elem(), newlinePrefix)

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

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.Itoa(int(value.Uint()))

	case reflect.Float32, reflect.Float64:
		return strconv.Itoa(int(value.Float()))

	case reflect.Bool:
		return strconv.FormatBool(value.Bool())

	case reflect.Interface:
		if m, ok := tryGetValue[gomock.Matcher](value); ok {
			return strconv.Quote(m.String())
		}
		return internalPrettyPrint(value.Elem(), newlinePrefix)

	case reflect.Pointer:
		elemPretty := internalPrettyPrint(value.Elem(), newlinePrefix)
		return fmt.Sprintf("&%s", elemPretty)

	case reflect.Struct:
		structType := value.Type()
		structTypeStr := internalPrettyPrintType(structType, newlinePrefix)
		numFields := value.NumField()
		if numFields == 0 {
			return fmt.Sprintf("%s{}", structTypeStr)
		}

		lines := make([]string, numFields)

		for idx := 0; idx < numFields; idx++ {
			value := value.FieldByIndex([]int{idx})
			keyStr := structType.FieldByIndex([]int{idx}).Name
			valueStr := internalPrettyPrint(value, newlinePrefix+prettyPrintIndentation)

			lines[idx] = fmt.Sprintf("%s%s: %s,\n%s",
				prettyPrintIndentation,
				keyStr,
				valueStr,
				newlinePrefix,
			)
		}
		slices.Sort(lines)
		return fmt.Sprintf("%s{\n%s%s}", structTypeStr, newlinePrefix, strings.Join(lines, ""))

	default:
		v := value.Interface()
		return fmt.Sprintf("%v (%T)", v, v)
	}
}
