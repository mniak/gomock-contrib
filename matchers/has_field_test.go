package matchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	_ gomock.Matcher = HasField("")
	_ gomock.Matcher = HasField("").ThatMatches("")
	_ gomock.Matcher = HasField("").ThatMatches(gomock.Any())
)

func TestHasField(t *testing.T) {
	sut := HasField("MyField")

	t.Run("struct", func(t *testing.T) {
		goodStruct := struct{ MyField string }{MyField: gofakeit.SentenceSimple()}
		withWrongKey := struct{ myfield string }{myfield: gofakeit.SentenceSimple()}

		assert.True(t, sut.Matches(goodStruct), "right value should match")
		assert.False(t, sut.Matches(withWrongKey), "wrong key should not match")
	})

	t.Run("map[string]any", func(t *testing.T) {
		goodMap := map[string]any{"MyField": gofakeit.SentenceSimple()}
		withWrongKey := map[string]any{"myfield": gofakeit.SentenceSimple()}

		assert.True(t, sut.Matches(goodMap), "right value should match")
		assert.False(t, sut.Matches(withWrongKey), "wrong key should not match")
	})

	t.Run("map[string]string", func(t *testing.T) {
		goodMap := map[string]string{"MyField": gofakeit.SentenceSimple()}
		withWrongKey := map[string]string{"myfield": gofakeit.SentenceSimple()}

		assert.True(t, sut.Matches(goodMap), "right value should match")
		assert.False(t, sut.Matches(withWrongKey), "wrong key should not match")
	})
}

func TestHasField_ThatMatches(t *testing.T) {
	fakeValue := gofakeit.SentenceSimple()

	testdata := []struct {
		name string
		sut  gomock.Matcher
	}{
		{
			name: "Using value directly as matcher",
			sut:  HasField("MyField").ThatMatches(fakeValue),
		},
		{
			name: "Using submatcher",
			sut:  HasField("MyField").ThatMatches(gomock.Eq(fakeValue)),
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			wrongFakeValue := gofakeit.SentenceSimple()

			t.Run("struct", func(t *testing.T) {
				goodStruct := struct{ MyField string }{MyField: fakeValue}
				withWrongKey := struct{ myfield string }{myfield: fakeValue}
				withWrongValue := struct{ MyField string }{MyField: wrongFakeValue}

				assert.True(t, td.sut.Matches(goodStruct), "right value should match")
				assert.False(t, td.sut.Matches(withWrongKey), "wrong key should not match")
				assert.False(t, td.sut.Matches(withWrongValue), "wrong value should not match")
			})

			t.Run("map[string]any", func(t *testing.T) {
				goodMap := map[string]any{"MyField": fakeValue}
				withWrongKey := map[string]any{"myfield": fakeValue}
				withWrongValue := map[string]any{"MyField": wrongFakeValue}

				assert.True(t, td.sut.Matches(goodMap), "right value should match")
				assert.False(t, td.sut.Matches(withWrongKey), "wrong key should not match")
				assert.False(t, td.sut.Matches(withWrongValue), "wrong value should not match")
			})

			t.Run("map[string]string", func(t *testing.T) {
				goodMap := map[string]string{"MyField": fakeValue}
				withWrongKey := map[string]string{"myfield": fakeValue}
				withWrongValue := map[string]string{"MyField": wrongFakeValue}

				assert.True(t, td.sut.Matches(goodMap), "right value should match")
				assert.False(t, td.sut.Matches(withWrongKey), "wrong key should not match")
				assert.False(t, td.sut.Matches(withWrongValue), "wrong value should not match")
			})
		})
	}
}
