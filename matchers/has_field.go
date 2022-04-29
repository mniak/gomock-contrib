package matchers

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/utils"
)

type hasFieldMatcher struct {
	fieldName string
}

func HasField(name string) hasFieldMatcher {
	return hasFieldMatcher{
		fieldName: name,
	}
}

func (m hasFieldMatcher) internalMatches(arg any) (any, bool) {
	value := reflect.ValueOf(arg)

	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Struct:
		structField := value.FieldByName(m.fieldName)
		if structField.Kind() == reflect.Invalid {
			return nil, false
		}
		return structField.Interface(), true

	case reflect.Map:
		mapValue := value.MapIndex(reflect.ValueOf(m.fieldName))
		if mapValue.Kind() == reflect.Invalid {
			return nil, false
		}
		return mapValue.Interface(), true
	}
	return nil, false
}

func (m hasFieldMatcher) Matches(arg any) bool {
	_, found := m.internalMatches(arg)
	return found
}

func (m hasFieldMatcher) String() string {
	return fmt.Sprintf("has field %s", m.fieldName)
}

func (m hasFieldMatcher) ThatMatches(matcher any) hasFieldThatMatchesMatcher {
	return hasFieldThatMatchesMatcher{
		parent:     m,
		submatcher: utils.ArgAsMatcher(matcher),
	}
}

type hasFieldThatMatchesMatcher struct {
	parent     hasFieldMatcher
	submatcher gomock.Matcher
}

func (m hasFieldThatMatchesMatcher) Matches(arg any) bool {
	value, found := m.parent.internalMatches(arg)
	if !found {
		return false
	}
	return m.submatcher.Matches(value)
}

func (m hasFieldThatMatchesMatcher) String() string {
	return fmt.Sprintf("has field %s that %s", m.parent.fieldName, m.submatcher.String())
}

func (m hasFieldThatMatchesMatcher) Got(arg any) string {
	field, found := m.parent.internalMatches(arg)
	if !found {
		return fmt.Sprintf("data without field %s: %v (%T)", m.parent.fieldName, arg, arg)
	}

	if gf, is := m.submatcher.(gomock.GotFormatter); is {
		subgot := gf.Got(field)
		return fmt.Sprintf("field %s %s", m.parent.fieldName, subgot)
	}
	return fmt.Sprintf("field .%s is %v (%T)", m.parent.fieldName, field, field)
}
