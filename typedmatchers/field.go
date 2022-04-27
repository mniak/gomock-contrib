package typedmatchers

import (
	"fmt"
	"reflect"

	"github.com/mniak/gomock-contrib/internal/utils"
)

type fieldMatcher[T any, F any] struct {
	fieldName string
	matcher   Matcher[F]
}

func Field[T any, F any](fieldName string, matcher Matcher[F]) fieldMatcher[T, F] {
	return fieldMatcher[T, F]{
		fieldName: fieldName,
		matcher:   matcher,
	}
}

func FieldGeneric[T any](fieldName string, matcher Matcher[any]) fieldMatcher[T, any] {
	return fieldMatcher[T, any]{
		fieldName: fieldName,
		matcher:   matcher,
	}
}

func (m fieldMatcher[T, F]) applySelector(x T) (F, bool) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Struct {
		field := value.FieldByName(m.fieldName)
		return utils.TryGetValue[F](field)
	}
	var f F
	return f, false
}

func (m fieldMatcher[T, F]) Matches(x T) bool {
	fieldValue, ok := m.applySelector(x)
	if !ok {
		return false
	}
	return m.matcher.Matches(fieldValue)
}

func (m fieldMatcher[T, F]) String() string {
	return fmt.Sprintf("field %s %s", m.fieldName, m.matcher.String())
}

func (m fieldMatcher[T, F]) Got(actual T) string {
	fieldValue, ok := m.applySelector(actual)
	if !ok {
		return fmt.Sprintf("field %s could not be found", m.fieldName)
	}
	return fmt.Sprintf("field %s: %s", m.fieldName, formatGottenArg(m.matcher, fieldValue))
}
