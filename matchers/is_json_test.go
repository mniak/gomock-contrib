package matchers

import (
	"encoding/json"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	_ gomock.Matcher = IsJSON()
	_ gomock.Matcher = IsJSON().ThatMatchesMap(map[string]any{})
)

func TestIsJSON(t *testing.T) {
	t.Run("Invalid JSON object should not match", func(t *testing.T) {
		testdata := []struct {
			desc        string
			invalidData string
		}{
			// strings
			{
				desc:        "missing closing bracket",
				invalidData: `{"key": "value"`,
			},
			{
				desc:        "missing opening bracket",
				invalidData: `"key": "value"}`,
			},
			{
				desc:        "double closing bracket",
				invalidData: `{"key": "value"}}`,
			},
			{
				desc:        "extra comma",
				invalidData: `{"key": "value",}`,
			},
			{
				desc:        "unquoted string",
				invalidData: `value`,
			},
		}
		for _, td := range testdata {
			t.Run(td.desc, func(t *testing.T) {
				sut := IsJSON()

				assert.False(t, sut.Matches(td.invalidData), "should not match string")
				assert.False(t, sut.Matches([]byte(td.invalidData)), "should not match bytes")
			})
		}
	})

	t.Run("Valid JSON object should match", func(t *testing.T) {
		testdata := []struct {
			desc      string
			validData any
		}{
			{
				desc:      "empty map",
				validData: map[string]any{},
			},
			{
				desc: "with fields",
				validData: map[string]any{
					"string_field": gofakeit.Word(),
					"int_field":    gofakeit.Int32(),
					"bool_field":   gofakeit.Bool(),
					"object_field": map[string]any{
						"string_field": gofakeit.Word(),
						"int_field":    gofakeit.Int32(),
						"bool_field":   gofakeit.Bool(),
					},
					"array_field": []map[string]any{
						{
							"string_field": gofakeit.Word(),
							"int_field":    gofakeit.Int32(),
							"bool_field":   gofakeit.Bool(),
						},
						{
							"string_field": gofakeit.Word(),
							"int_field":    gofakeit.Int32(),
							"bool_field":   gofakeit.Bool(),
						},
					},
				},
			},
			{
				desc:      "string",
				validData: gofakeit.SentenceSimple(),
			},
			{
				desc:      "float64",
				validData: gofakeit.Float64(),
			},
			{
				desc:      "int32",
				validData: gofakeit.Int32(),
			},
			{
				desc:      "boolean",
				validData: gofakeit.Bool(),
			},
			{
				desc: "boolean",
				validData: func() any {
					result := make([]map[string]any, 10)
					gofakeit.Slice(&result)
					return result
				}(),
			},
		}
		for _, td := range testdata {
			t.Run(td.desc, func(t *testing.T) {
				dataBytes, err := json.Marshal(td.validData)
				require.NoError(t, err)

				sut := IsJSON()

				assert.True(t, sut.Matches(dataBytes), "should match bytes")
				assert.True(t, sut.Matches(string(dataBytes)), "should match string")
			})
		}
	})
}
