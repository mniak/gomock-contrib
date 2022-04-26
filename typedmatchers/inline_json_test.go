package typedmatchers

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ Matcher[string] = InlineJSON[StructForTestsWithTypedInlineJSON](nil)

type StructForTestsWithTypedInlineJSON struct {
	Field1 string
	Field2 int
}

func TestInlineJSON(t *testing.T) {
	t.Run("Nil function should behave as a return true fn", func(t *testing.T) {
		m := InlineJSON[StructForTestsWithTypedInlineJSON](nil)
		assert.True(t, m.Matches(`{}`), "empty JSON object should match")
		assert.False(t, m.Matches(`{`), "Invalid JSON should cause failure")
		assert.False(t, m.Matches(`abc123`), "String should cause failure")
		assert.False(t, m.Matches(`"abc123"`), "Quoted String should cause failure")
		assert.False(t, m.Matches(`123`), "int should cause failure")
		assert.False(t, m.Matches(``), "empty string should cause failure")
	})

	t.Run("Basic tests", func(t *testing.T) {
		m := InlineJSON(func(x StructForTestsWithTypedInlineJSON) bool {
			return true
		})
		assert.True(t, m.Matches(`{}`), "empty JSON object should match")
		assert.False(t, m.Matches(`{`), "Invalid JSON should cause failure")
		assert.False(t, m.Matches(`abc123`), "String should cause failure")
		assert.False(t, m.Matches(`123`), "int should cause failure")
		assert.False(t, m.Matches(``), "empty string should cause failure")
	})

	t.Run("Simple inline boolean", func(t *testing.T) {
		for _, b := range []bool{true, false} {
			t.Run(fmt.Sprintf("when inline fn returns %t", b), func(t *testing.T) {
				m := InlineJSON(func(x StructForTestsWithTypedInlineJSON) bool {
					return b
				})
				assert.Equal(t, b, m.Matches(`{}`))
			})
		}
	})

	t.Run("Happy scenario", func(t *testing.T) {
		var sample StructForTestsWithTypedInlineJSON
		gofakeit.Struct(&sample)

		m := InlineJSON(func(x StructForTestsWithTypedInlineJSON) bool {
			return x.Field1 == sample.Field1 &&
				x.Field2 == sample.Field2
		})

		jsonBytes, err := json.Marshal(sample)
		require.NoError(t, err)

		assert.True(t, m.Matches(string(jsonBytes)))
	})
}
