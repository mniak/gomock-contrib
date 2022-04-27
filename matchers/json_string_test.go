package matchers

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ gomock.Matcher = JSONString("{}")

type StructForTestsWithJSONString struct {
	FieldString string
	FieldInt    int
	FieldBool   bool
	FieldStruct struct {
		SubFieldString string
		SubFieldInt    int
	}
	ExtraField string
}

func (s StructForTestsWithJSONString) Clone() StructForTestsWithJSONString {
	return s
}

func TestJSON(t *testing.T) {
	t.Run("When type is not string, should fail", func(t *testing.T) {
		m := JSONString("{}")
		assert.False(t, m.Matches(nil), "nil should fail")
		assert.False(t, m.Matches(gofakeit.Bool()), "bool should fail")
		assert.False(t, m.Matches(gofakeit.Int32()), "int32 should fail")
		assert.False(t, m.Matches(StructForTestsWithJSONString{}), "struct should fail")
	})

	t.Run("Basic tests", func(t *testing.T) {
		m := JSONString("{}")
		assert.True(t, m.Matches(`{}`), "empty JSON object should match")
		assert.False(t, m.Matches(`{`), "Invalid JSON should cause failure")
		assert.False(t, m.Matches(`abc123`), "String should cause failure")
		assert.False(t, m.Matches(`"abc123"`), "Quoted String should cause failure")
		assert.False(t, m.Matches(`123`), "int should cause failure")
		assert.False(t, m.Matches(``), "empty string should cause failure")
	})

	t.Run("With field missing on expression, should match", func(t *testing.T) {
		var sample StructForTestsWithJSONString
		gofakeit.Struct(&sample)
		m := JSONString(fmt.Sprintf(`{
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
		var sample StructForTestsWithJSONString
		gofakeit.Struct(&sample)
		m := JSONString(fmt.Sprintf(`{
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
		var sample StructForTestsWithJSONString
		gofakeit.Struct(&sample)
		m := JSONString(fmt.Sprintf(`{
			"FieldString": "%s",
			"FieldInt": %d,
			"FieldStruct": {
				"InexistentSubfield": "this subfield does not exist"
			}
		}`, sample.FieldString, sample.FieldInt))
		jsonBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(jsonBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(jsonBytes)), "As string should fail")
	})

	t.Run("When field has different value, should NOT match", func(t *testing.T) {
		var sample StructForTestsWithJSONString
		gofakeit.Struct(&sample)
		changedSample := sample.Clone()
		changedSample.FieldString = gofakeit.SentenceSimple()

		changedSampleBytes, err := json.Marshal(changedSample)
		require.NoError(t, err)

		m := JSONString(string(changedSampleBytes))

		sampleBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(sampleBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(sampleBytes)), "As string should fail")
	})

	t.Run("When subfield has different value, should NOT match", func(t *testing.T) {
		var sample StructForTestsWithJSONString
		gofakeit.Struct(&sample)
		changedSample := sample.Clone()
		changedSample.FieldStruct.SubFieldString = gofakeit.SentenceSimple()

		changedSampleBytes, err := json.Marshal(changedSample)
		require.NoError(t, err)

		m := JSONString(string(changedSampleBytes))

		sampleBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.False(t, m.Matches(sampleBytes), "As bytes should fail")
		assert.False(t, m.Matches(string(sampleBytes)), "As string should fail")
	})
}
