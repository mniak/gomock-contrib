package matchers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ gomock.Matcher = JSON("")

type StructForTestsWithJSON struct {
	FieldString string
	FieldInt    int
	FieldBool   bool
	FieldStruct struct {
		SubFieldString string
		SubFieldInt    int
	}
	ExtraField string
}

func (s StructForTestsWithJSON) Clone() StructForTestsWithJSON {
	return s
}

func TestJSON(t *testing.T) {
	t.Run("When type is not string, should fail", func(t *testing.T) {
		m := JSON("{}")
		assert.False(t, m.Matches(nil), "nil should fail")
		assert.False(t, m.Matches(gofakeit.Bool()), "bool should fail")
		assert.False(t, m.Matches(gofakeit.Int32()), "int32 should fail")
		assert.False(t, m.Matches(StructForTestsWithJSON{}), "struct should fail")
	})

	t.Run("Basic tests", func(t *testing.T) {
		m := JSON("{}")
		assert.True(t, m.Matches(`{}`), "empty JSON object should match")
		assert.False(t, m.Matches(`{`), "Invalid JSON should cause failure")
		assert.False(t, m.Matches(`abc123`), "String should cause failure")
		assert.False(t, m.Matches(`"abc123"`), "Quoted String should cause failure")
		assert.False(t, m.Matches(`123`), "int should cause failure")
		assert.False(t, m.Matches(``), "empty string should cause failure")
	})

	t.Run("With field missing on expression, should match", func(t *testing.T) {
		var sample StructForTestsWithJSON
		gofakeit.Struct(&sample)
		m := JSON(fmt.Sprintf(`{
			"FieldString": "%s",
			"FieldInt": %d,
			"FieldBool": %t,
			"FieldStruct": {
				"SubFieldString": "%s",
				"SubFieldInt": %d
			}
		}`,
			sample.FieldString,
			sample.FieldInt,
			sample.FieldBool,
			sample.FieldStruct.SubFieldString,
			sample.FieldStruct.SubFieldInt,
		))

		jsonBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.True(t, m.Matches(jsonBytes), "As bytes should succeed")
		assert.True(t, m.Matches(string(jsonBytes)), "As string should succeed")
	})

	t.Run("With field on expression is not on JSON, should NOT match", func(t *testing.T) {
		var sample StructForTestsWithJSON
		gofakeit.Struct(&sample)
		m := JSON(fmt.Sprintf(`{
			"FieldString": "%s",
			"FieldInt": %d,
			"InexistentField": "this field does not exist"
		}`, sample.FieldString, sample.FieldInt))

		jsonBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(jsonBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(jsonBytes)), "As string should fail")
	})

	t.Run("With subfield on expression is not on JSON, should NOT match", func(t *testing.T) {
		var sample StructForTestsWithJSON
		gofakeit.Struct(&sample)
		m := JSON(fmt.Sprintf(`{
			"FieldString": "%s",
			"FieldInt": %d,
			"FieldStruct": {
				"InexistentSubfield": "this subfield does not exist"
			},
		}`, sample.FieldString, sample.FieldInt))
		jsonBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(jsonBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(jsonBytes)), "As string should fail")
	})

	t.Run("When field has different value, should NOT match", func(t *testing.T) {
		var sample StructForTestsWithJSON
		gofakeit.Struct(&sample)
		changedSample := sample.Clone()
		changedSample.FieldString = gofakeit.SentenceSimple()

		changedSampleBytes, err := json.Marshal(changedSample)
		require.NoError(t, err)

		m := JSON(string(changedSampleBytes))

		sampleBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(sampleBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(sampleBytes)), "As string should fail")
	})

	t.Run("When subfield has different value, should NOT match", func(t *testing.T) {
		var sample StructForTestsWithJSON
		gofakeit.Struct(&sample)
		changedSample := sample.Clone()
		changedSample.FieldStruct.SubFieldString = gofakeit.SentenceSimple()

		changedSampleBytes, err := json.Marshal(changedSample)
		require.NoError(t, err)

		m := JSON(string(changedSampleBytes))

		sampleBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(sampleBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(sampleBytes)), "As string should fail")
	})
}

func Test_matchMaps(t *testing.T) {
	t.Run("When maps are different, should not match", func(t *testing.T) {
		fakenumber := int(gofakeit.Int32())
		faketext := gofakeit.SentenceSimple()

		testdata := []struct {
			name     string
			expected map[string]any
			actual   map[string]any
		}{
			{
				name: "fields have different types: Int/String",
				expected: map[string]any{
					"field1": fakenumber,
				},
				actual: map[string]any{
					"field1": strconv.Itoa(fakenumber),
				},
			},
			{
				name: "fields have different types: String/Int",
				expected: map[string]any{
					"field1": strconv.Itoa(fakenumber),
				},
				actual: map[string]any{
					"field1": fakenumber,
				},
			},
			{
				name: "fields have same type but different value",
				expected: map[string]any{
					"field1": faketext,
				},
				actual: map[string]any{
					"field1": faketext + "-suffix",
				},
			},
			// With sub struct
			{
				name: "fields have different types: Int/String",
				expected: map[string]any{
					"root": map[string]any{
						"field1": fakenumber,
					},
				},
				actual: map[string]any{
					"root": map[string]any{
						"field1": strconv.Itoa(fakenumber),
					},
				},
			},
			{
				name: "fields have different types: String/Int",
				expected: map[string]any{
					"root": map[string]any{
						"field1": strconv.Itoa(fakenumber),
					},
				},
				actual: map[string]any{
					"root": map[string]any{
						"field1": fakenumber,
					},
				},
			},
			{
				name: "fields have same type but different value",
				expected: map[string]any{
					"root": map[string]any{
						"field1": faketext,
					},
				},
				actual: map[string]any{
					"root": map[string]any{
						"field1": faketext + "-suffix",
					},
				},
			},
			// equal, but expecting more fields
			{
				name: "actual empty, expecting more fields",
				expected: map[string]any{
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
				},
				actual: map[string]any{},
			},
			{
				name: "expecting more fields",
				expected: map[string]any{
					"string":        faketext,
					"int":           fakenumber,
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
				},
				actual: map[string]any{
					"string": faketext,
					"int":    fakenumber,
				},
			},
			{
				name: "array in wrong order",
				expected: map[string]any{
					"array": []string{
						faketext + "0",
						faketext + "1",
					},
				},
				actual: map[string]any{
					"array": []string{
						faketext + "1",
						faketext + "0",
					},
				},
			},
		}
		for _, td := range testdata {
			t.Run(td.name, func(t *testing.T) {
				assert.False(t, matchMaps(td.expected, td.actual), "should not match")
			})
		}
	})

	t.Run("When map is exactly the same, should match", func(t *testing.T) {
		testdata := []struct {
			name  string
			value any
		}{
			{
				name:  "string",
				value: gofakeit.SentenceSimple(),
			},
			// int
			{
				name:  "int",
				value: int(gofakeit.Int32()),
			},
			{
				name:  "int8",
				value: gofakeit.Int8(),
			},
			{
				name:  "int16",
				value: gofakeit.Int16(),
			},
			{
				name:  "int32",
				value: gofakeit.Int32(),
			},
			{
				name:  "int64",
				value: gofakeit.Int64(),
			},
			// uint
			{
				name:  "uint",
				value: uint(gofakeit.Uint32()),
			},
			{
				name:  "uint8",
				value: gofakeit.Uint8(),
			},
			{
				name:  "uint16",
				value: gofakeit.Uint16(),
			},
			{
				name:  "uint32",
				value: gofakeit.Uint32(),
			},
			{
				name:  "uint64",
				value: gofakeit.Uint64(),
			},
			// maps, slice, array
			{
				name: "map",
				value: map[any]any{
					1: gofakeit.FarmAnimal(),
				},
			},
			{
				name: "slice",
				value: []any{
					gofakeit.FarmAnimal(),
				},
			},
			{
				name: "array",
				value: [1]any{
					gofakeit.FarmAnimal(),
				},
			},
		}
		for _, td := range testdata {
			t.Run(td.name, func(t *testing.T) {
				map1 := map[string]any{
					"field1": td.value,
				}

				map2 := map[string]any{
					"field1": td.value,
				}

				assert.True(t, matchMaps(map1, map2), "should match")
			})
		}
	})
	t.Run("Test cases that should match", func(t *testing.T) {
		fakenumber := int(gofakeit.Int32())
		faketext := gofakeit.SentenceSimple()

		testdata := []struct {
			name     string
			expected map[string]any
			actual   map[string]any
		}{
			{
				name: "exactly the same: String",
				expected: map[string]any{
					"field1": faketext,
				},
				actual: map[string]any{
					"field1": faketext,
				},
			},
			{
				name: "exactly the same: Int",
				expected: map[string]any{
					"field1": fakenumber,
				},
				actual: map[string]any{
					"field1": fakenumber,
				},
			},
			{
				name:     "exactly the same: empty",
				expected: map[string]any{},
				actual:   map[string]any{},
			},
			{
				name:     "expected has less fields than actual (has none)",
				expected: map[string]any{},
				actual: map[string]any{
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
				},
			},
			{
				name: "expected has less fields than actual (matches some)",
				expected: map[string]any{
					"string": faketext,
					"int":    fakenumber,
				},
				actual: map[string]any{
					"string":        faketext,
					"int":           fakenumber,
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
					gofakeit.UUID(): gofakeit.SentenceSimple(),
				},
			},
			{
				name: "slice",
				expected: map[string]any{
					"slice": []string{
						faketext,
					},
				},
				actual: map[string]any{
					"slice": []string{
						faketext,
					},
				},
			},
			{
				name: "slice in right order",
				expected: map[string]any{
					"slice": []string{
						faketext + "0",
						faketext + "1",
					},
				},
				actual: map[string]any{
					"slice": []string{
						faketext + "0",
						faketext + "1",
					},
				},
			},
		}
		for _, td := range testdata {
			t.Run(td.name, func(t *testing.T) {
				assert.True(t, matchMaps(td.expected, td.actual), "should match")
			})
		}
	})
}
