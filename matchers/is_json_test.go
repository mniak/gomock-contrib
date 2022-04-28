package matchers

import (
	"encoding/json"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// var (
// 	_ gomock.Matcher      = JSON(map[string]any{})
// 	_ gomock.GotFormatter = JSON(map[string]any{})
// )

// type StructForTestsWithJSONObject struct {
// 	FieldString string
// 	FieldInt    int
// 	FieldBool   bool
// 	FieldStruct struct {
// 		SubFieldString string
// 		SubFieldInt    int
// 	}
// 	ExtraField string
// }

// func (s StructForTestsWithJSONObject) Clone() StructForTestsWithJSONObject {
// 	return s
// }

// func (s StructForTestsWithJSONObject) ToMap() map[string]any {
// 	bytes, _ := json.Marshal(s)
// 	var result map[string]any
// 	json.Unmarshal(bytes, &result)
// 	return result
// }

// func TestJSONObject(t *testing.T) {
// 	t.Run("Basic tests", func(t *testing.T) {
// 		m := JSON(map[string]any{})
// 		assert.True(t, m.Matches(`{}`), "empty JSON object should match")
// 		assert.False(t, m.Matches(`{`), "Invalid JSON should cause failure")
// 		assert.False(t, m.Matches(`abc123`), "String should cause failure")
// 		assert.False(t, m.Matches(`"abc123"`), "Quoted String should cause failure")
// 		assert.False(t, m.Matches(`123`), "int should cause failure")
// 		assert.False(t, m.Matches(``), "empty string should cause failure")
// 	})

// 	t.Run("With field missing on expression, should match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		expectedMap := map[string]any{
// 			"FieldString": sample.FieldString,
// 			"FieldInt":    sample.FieldInt,
// 			"FieldBool":   sample.FieldBool,
// 			"FieldStruct": map[string]any{
// 				"SubFieldString": sample.FieldStruct.SubFieldString,
// 				"SubFieldInt":    sample.FieldStruct.SubFieldInt,
// 			},
// 		}
// 		sut := JSON(expectedMap)

// 		jsonBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.True(t, sut.Matches(string(jsonBytes)), "As string should succeed")
// 		assert.True(t, sut.Matches(jsonBytes), "As bytes should succeed")
// 	})

// 	t.Run("With field on expression is not on JSON, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		expectedMap := map[string]any{
// 			"FieldString":     sample.FieldString,
// 			"FieldInt":        sample.FieldInt,
// 			"InexistentField": "this field does not exist",
// 		}

// 		sut := JSON(expectedMap)

// 		jsonBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(jsonBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(jsonBytes), "As bytes should fail")
// 	})

// 	t.Run("With subfield on expression is not on JSON, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		expectedMap := map[string]any{
// 			"FieldString": sample.FieldString,
// 			"FieldInt":    sample.FieldInt,
// 			"FieldStruct": map[string]any{
// 				"InexistentSubfield": "this subfield does not exist",
// 			},
// 		}

// 		sut := JSON(expectedMap)

// 		jsonBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(jsonBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(jsonBytes), "As bytes should fail")
// 	})

// 	t.Run("When field has different value, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		changedSample := sample.Clone()
// 		changedSample.FieldString = gofakeit.SentenceSimple()

// 		sut := JSON(changedSample.ToMap())

// 		sampleBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(sampleBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(sampleBytes), "As bytes should fail")
// 	})

// 	t.Run("When subfield has different value, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		changedSample := sample.Clone()
// 		changedSample.FieldStruct.SubFieldString = gofakeit.SentenceSimple()

// 		sut := JSON(changedSample.ToMap())

// 		sampleBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(sampleBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(sampleBytes), "As bytes should fail")
// 	})
// }

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

// func TestIsJSON_ThatMatchesMap(t *testing.T) {
// 	t.Run("Basic tests", func(t *testing.T) {
// 		m := JSON(map[string]any{})
// 		assert.True(t, m.Matches(`{}`), "empty JSON object should match")
// 		assert.False(t, m.Matches(`{`), "Invalid JSON should cause failure")
// 		assert.False(t, m.Matches(`abc123`), "String should cause failure")
// 		assert.False(t, m.Matches(`"abc123"`), "Quoted String should cause failure")
// 		assert.False(t, m.Matches(`123`), "int should cause failure")
// 		assert.False(t, m.Matches(``), "empty string should cause failure")
// 	})

// 	t.Run("With field missing on expression, should match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		expectedMap := map[string]any{
// 			"FieldString": sample.FieldString,
// 			"FieldInt":    sample.FieldInt,
// 			"FieldBool":   sample.FieldBool,
// 			"FieldStruct": map[string]any{
// 				"SubFieldString": sample.FieldStruct.SubFieldString,
// 				"SubFieldInt":    sample.FieldStruct.SubFieldInt,
// 			},
// 		}
// 		sut := IsJSON().ThatMatchesMap(expectedMap)

// 		jsonBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.True(t, sut.Matches(string(jsonBytes)), "As string should succeed")
// 		assert.True(t, sut.Matches(jsonBytes), "As bytes should succeed")
// 	})

// 	t.Run("With field on expression is not on JSON, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		expectedMap := map[string]any{
// 			"FieldString":     sample.FieldString,
// 			"FieldInt":        sample.FieldInt,
// 			"InexistentField": "this field does not exist",
// 		}

// 		sut := IsJSON().ThatMatchesMap(expectedMap)

// 		jsonBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(jsonBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(jsonBytes), "As bytes should fail")
// 	})

// 	t.Run("With subfield on expression is not on JSON, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		expectedMap := map[string]any{
// 			"FieldString": sample.FieldString,
// 			"FieldInt":    sample.FieldInt,
// 			"FieldStruct": map[string]any{
// 				"InexistentSubfield": "this subfield does not exist",
// 			},
// 		}

// 		sut := IsJSON().ThatMatchesMap(expectedMap)

// 		jsonBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(jsonBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(jsonBytes), "As bytes should fail")
// 	})

// 	t.Run("When field has different value, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		changedSample := sample.Clone()
// 		changedSample.FieldString = gofakeit.SentenceSimple()

// 		sut := IsJSON().ThatMatchesMap(changedSample.ToMap())

// 		sampleBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(sampleBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(sampleBytes), "As bytes should fail")
// 	})

// 	t.Run("When subfield has different value, should NOT match", func(t *testing.T) {
// 		var sample StructForTestsWithJSONObject
// 		gofakeit.Struct(&sample)
// 		changedSample := sample.Clone()
// 		changedSample.FieldStruct.SubFieldString = gofakeit.SentenceSimple()

// 		sut := IsJSON().ThatMatchesMap(changedSample.ToMap())

// 		sampleBytes, err := json.Marshal(sample)
// 		require.NoError(t, err)

// 		assert.False(t, sut.Matches(string(sampleBytes)), "As string should fail")
// 		assert.False(t, sut.Matches(sampleBytes), "As bytes should fail")
// 	})
// }
