package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrintMap(t *testing.T) {
	testdata := []struct {
		name     string
		sample   map[string]any
		expected string
	}{
		{
			name:     "Nil map",
			sample:   nil,
			expected: "nil",
		},
		{
			name:     "Empty map",
			sample:   map[string]any{},
			expected: "map[string]any{}",
		},
		{
			name: "String field",
			sample: map[string]any{
				"String": "value",
			},
			expected: `map[string]any{
	"String": "value",
}`,
		},
		{
			name: "String field",
			sample: map[string]any{
				"StringPointer": ToPointer("value"),
			},
			expected: `map[string]any{
	"StringPointer": &"value",
}`,
		},
		{
			name: "Integer field",
			sample: map[string]any{
				"Integer": 123456,
			},
			expected: `map[string]any{
	"Integer": 123456,
}`,
		},
		{
			name: "Boolean field",
			sample: map[string]any{
				"Boolean": true,
			},
			expected: `map[string]any{
	"Boolean": true,
}`,
		},
		{
			name: "String Slice field",
			sample: map[string]any{
				"StringSlice": []string{
					"Item1",
					"Item2",
				},
			},
			expected: `map[string]any{
	"StringSlice": []string{
		"Item1",
		"Item2",
	},
}`,
		},
		{
			name: "Map[string]any field",
			sample: map[string]any{
				"MapStringAny": map[string]any{
					"Integer": 123,
				},
			},
			expected: `map[string]any{
	"MapStringAny": map[string]any{
		"Integer": 123,
	},
}`,
		},
		{
			name: "Struct field",
			sample: map[string]any{
				"Struct": struct {
					String  string
					Integer int
				}{
					String:  "string value",
					Integer: 1234,
				},
			},
			expected: `map[string]any{
	"Struct": struct {
		Integer int
		String string
	}{
		Integer: 1234,
		String: "string value",
	},
}`,
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			result := PrettyPrintMap(td.sample)
			assert.Equal(t, td.expected, result)
		})
	}
}
