package matchers

import (
	"encoding/json"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/testing/mocks"
	"github.com/mniak/gomock-contrib/internal/utils"
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

				assert.Equal(t, "is valid JSON", sut.String())
				assert.False(t, sut.Matches(td.invalidData), "should not match string")
				assert.Equal(t, "is "+td.invalidData+" (string)", sut.Got(td.invalidData))
				assert.False(t, sut.Matches([]byte(td.invalidData)), "should not match bytes")
				assert.Equal(t, "is "+td.invalidData+" ([]byte)", sut.Got([]byte(td.invalidData)))
			})
		}
	})

	t.Run("When not string or bytes, should fail", func(t *testing.T) {
		testdata := []struct {
			desc        string
			invalidData any
			gotMessage  string
		}{
			{
				desc:        "int",
				invalidData: 123,
				gotMessage:  "data with invalid type: 123 (int)",
			},
			{
				desc:        "boolean",
				invalidData: false,
				gotMessage:  "data with invalid type: false (bool)",
			},
		}
		for _, td := range testdata {
			t.Run(td.desc, func(t *testing.T) {
				sut := IsJSON()

				assert.Equal(t, "is valid JSON", sut.String())
				assert.False(t, sut.Matches(td.invalidData), "should not match")
				assert.Equal(t, td.gotMessage, sut.Got(td.invalidData))
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
				desc:      "pointer to string",
				validData: utils.ToPointer(gofakeit.SentenceSimple()),
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

func TestIsJSON_ThatMatches_Messages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testdata := []struct {
		name         string
		sut          isJSONThatMatchesMatcher
		sampleValue  any
		expectedGot  string
		expectedWant string
	}{
		{
			name: "Invalid type (int)",
			sut: func() isJSONThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<wrong_int>")
				return IsJSON().ThatMatches(mock)
			}(),
			sampleValue:  123,
			expectedGot:  "data with invalid type: 123 (int)",
			expectedWant: "is a valid JSON that <wrong_int>",
		},
		{
			name: "Invalid type (boolean)",
			sut: func() isJSONThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<wrong_bool>")
				return IsJSON().ThatMatches(mock)
			}(),
			sampleValue:  false,
			expectedGot:  "data with invalid type: false (bool)",
			expectedWant: "is a valid JSON that <wrong_bool>",
		},
		{
			name: "Ill-formatted json",
			sut: func() isJSONThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<ill_test>")
				return IsJSON().ThatMatches(mock)
			}(),
			sampleValue:  `{ "key": "value"`,
			expectedGot:  `is { "key": "value" (string)`,
			expectedWant: "is a valid JSON that <ill_test>",
		},
		{
			name: "JSON not matching",
			sut: func() isJSONThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<want_json_not_matching>")
				mock.EXPECT().Got(map[string]any{
					"key": "value",
				}).Return("<got_json_not_matching>")
				return IsJSON().ThatMatches(mock)
			}(),
			sampleValue:  `{"key": "value"}`,
			expectedGot:  "<got_json_not_matching>",
			expectedWant: "is a valid JSON that <want_json_not_matching>",
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			wantMessage := td.sut.String()
			gotMessage := td.sut.Got(td.sampleValue)

			assert.Equal(t, td.expectedWant, wantMessage)
			assert.Equal(t, td.expectedGot, gotMessage)
		})
	}
}
