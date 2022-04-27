package utils

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestTryGetValue(t *testing.T) {
	testdata := []struct {
		name     string
		expected any
		fn       func(reflect.Value) (any, bool)
	}{
		{
			name:     "string",
			expected: gofakeit.SentenceSimple(),
			fn: func(v reflect.Value) (any, bool) {
				return TryGetValue[string](v)
			},
		},
		{
			name:     "int",
			expected: int(gofakeit.Int64()),
			fn: func(v reflect.Value) (any, bool) {
				return TryGetValue[int](v)
			},
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			reflectedValue := reflect.ValueOf(td.expected)
			result, ok := td.fn(reflectedValue)
			assert.True(t, ok, "should work")
			assert.Equal(t, td.expected, result)
		})
	}
}
