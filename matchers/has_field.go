package matchers

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

type hasFieldMatcher struct {
	fieldName string
}

func HasField(name string) hasFieldMatcher {
	return hasFieldMatcher{
		fieldName: name,
	}
}

func (m hasFieldMatcher) internalMatches(arg any) (reflect.Value, bool) {
	value := reflect.ValueOf(arg)

	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Struct:
		structField := value.FieldByName(m.fieldName)
		skind := structField.Kind()
		_ = skind
		return structField, structField.Kind() != reflect.Invalid

	case reflect.Map:
		mapValue := value.MapIndex(reflect.ValueOf(m.fieldName))
		return mapValue, mapValue.Kind() != reflect.Invalid
	}
	return reflect.Value{}, false
}

func (m hasFieldMatcher) Matches(arg any) bool {
	_, found := m.internalMatches(arg)
	return found
}

func (m hasFieldMatcher) String() string {
	return fmt.Sprintf("has field %s", m.fieldName)
}

func (m hasFieldMatcher) ThatMatches(arg any) hasFieldThatMatchMatcher {
	var submatcher gomock.Matcher
	if sub, is := arg.(gomock.Matcher); is {
		submatcher = sub
	} else {
		submatcher = gomock.Eq(arg)
	}
	return hasFieldThatMatchMatcher{
		parent:     m,
		submatcher: submatcher,
	}
}

type hasFieldThatMatchMatcher struct {
	parent     hasFieldMatcher
	submatcher gomock.Matcher
}

func (m hasFieldThatMatchMatcher) Matches(arg any) bool {
	value, found := m.parent.internalMatches(arg)
	if !found || !value.CanInterface() {
		return false
	}
	val := value.Interface()
	return m.submatcher.Matches(val)
}

func (m hasFieldThatMatchMatcher) String() string {
	return ""
}
