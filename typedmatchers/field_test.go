package typedmatchers

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	_ Matcher[bool]      = fieldMatcher[bool, string]{}
	_ GotFormatter[bool] = fieldMatcher[bool, string]{}
)

type StructForFieldMatcherTest struct {
	StringField1 string
	StringField2 string
}

func TestFieldMatcher(t *testing.T) {
	expectedString := gofakeit.SentenceSimple()
	wantMatcherString := "value should be exactly the same"
	expectedMatcherString := fmt.Sprintf("field StringField1 %s", wantMatcherString)

	sut := Field[StructForFieldMatcherTest, string](
		"StringField1",
		Inline(wantMatcherString, func(x string) bool {
			return x == expectedString
		}),
	)

	var fakeStruct StructForFieldMatcherTest
	gofakeit.Struct(&fakeStruct)
	assert.False(t, sut.Matches(fakeStruct), "when struct DOES NOT have the expected string, should NOT match")

	fakeStruct.StringField1 = expectedString
	assert.True(t, sut.Matches(fakeStruct), "when struct HAS the expected string, should match")

	assert.Equal(t, expectedMatcherString, sut.String())
}

func TestFieldMatcherInterface(t *testing.T) {
	expectedString := gofakeit.SentenceSimple()

	sut := FieldGeneric[StructForFieldMatcherTest]("StringField1", gomock.Eq(expectedString))

	var fakeStruct StructForFieldMatcherTest
	gofakeit.Struct(&fakeStruct)
	assert.False(t, sut.Matches(fakeStruct), "when struct DOES NOT have the expected string, should NOT match")

	fakeStruct.StringField1 = expectedString
	assert.True(t, sut.Matches(fakeStruct), "when struct HAS the expected string, should match")
}

func TestFieldEqual(t *testing.T) {
	expectedString := gofakeit.SentenceSimple()

	sut := FieldEqual[StructForFieldMatcherTest]("StringField1", expectedString)

	var fakeStruct StructForFieldMatcherTest
	gofakeit.Struct(&fakeStruct)
	assert.False(t, sut.Matches(fakeStruct), "when struct DOES NOT have the expected string, should NOT match")

	fakeStruct.StringField1 = expectedString
	assert.True(t, sut.Matches(fakeStruct), "when struct HAS the expected string, should match")
}
