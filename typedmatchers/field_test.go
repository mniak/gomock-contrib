package typedmatchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

var _ Matcher[bool] = fieldMatcher[bool, string]{}

type StructForFieldMatcherTest struct {
	StringField1 string
	StringField2 string
}

func Test_fieldMatcher(t *testing.T) {
	t.Run("Using typed inline matcher", func(t *testing.T) {
		expectedString := gofakeit.SentenceSimple()

		sut := MatchField[StructForFieldMatcherTest, string](func(x StructForFieldMatcherTest) string {
			return x.StringField1
		}, Inline("value should be exactly the same", func(x string) bool {
			return x == expectedString
		}))

		var fakeStruct StructForFieldMatcherTest
		gofakeit.Struct(&fakeStruct)
		assert.False(t, sut.Matches(fakeStruct), "when struct DOES NOT have the expected string, should NOT match")

		fakeStruct.StringField1 = expectedString
		assert.True(t, sut.Matches(fakeStruct), "when struct HAS the expected string, should match")
	})
}
