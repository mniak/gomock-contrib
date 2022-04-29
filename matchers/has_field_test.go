package matchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/testing/mocks"
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

func TestHasField_ThatMatches_Messages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testdata := []struct {
		name         string
		sut          hasFieldThatMatchesMatcher
		sampleValue  any
		expectedGot  string
		expectedWant string
	}{
		// Without field
		{
			name:         "Can't find field, without submatcher (string)",
			sut:          HasField("MyField").ThatMatches("field_value"),
			sampleValue:  "wrong_value",
			expectedGot:  "data without field MyField: wrong_value (string)",
			expectedWant: "has field MyField that is equal to field_value (string)",
		},
		{
			name:         "Can't find field, with submatcher (string)",
			sut:          HasField("MyField").ThatMatches(gomock.Eq("field_value")),
			sampleValue:  "wrong_value",
			expectedGot:  "data without field MyField: wrong_value (string)",
			expectedWant: "has field MyField that is equal to field_value (string)",
		},
		// Int
		{
			name:         "Can't find field, without submatcher (int)",
			sut:          HasField("MyField").ThatMatches("field_value"),
			sampleValue:  123,
			expectedGot:  "data without field MyField: 123 (int)",
			expectedWant: "has field MyField that is equal to field_value (string)",
		},
		{
			name:         "Can't find field, with submatcher (int)",
			sut:          HasField("MyField").ThatMatches(gomock.Eq("field_value")),
			sampleValue:  123,
			expectedGot:  "data without field MyField: 123 (int)",
			expectedWant: "has field MyField that is equal to field_value (string)",
		},
		// Mocked submatcher
		{
			name: "Can't find field, with mocked submatcher",
			sut: func() hasFieldThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<submatcher.String()>")
				return HasField("MyField").ThatMatches(mock)
			}(),
			sampleValue:  "wrong_value",
			expectedGot:  "data without field MyField: wrong_value (string)",
			expectedWant: "has field MyField that <submatcher.String()>",
		},
		{
			name: "Struct",
			sut: func() hasFieldThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<want_struct_field>")
				mock.EXPECT().Got("struct_myfield_value").Return("<got_struct_field>")
				return HasField("MyField").ThatMatches(mock)
			}(),
			sampleValue:  struct{ MyField string }{MyField: "struct_myfield_value"},
			expectedGot:  "field MyField <got_struct_field>",
			expectedWant: "has field MyField that <want_struct_field>",
		},
		// {
		// 	name: "Using mocked submatcher (int)",
		// 	sut: func() hasFieldThatMatchesMatcher {
		// 		mock := mocks.NewMockMatcherGotFormatter(ctrl)
		// 		mock.EXPECT().String().Return("<submatcher.String()>")
		// 		return HasField("MyField").ThatMatches(mock)
		// 	}(),
		// 	sampleValue:  123,
		// 	expectedGot:  ".MyField is 123 (int)",
		// 	expectedWant: ".MyField <submatcher.String()>",
		// },
		// // Mocked submatcher that implements GotMatcher
		// {
		// 	name: "Using mocked submatcher that is GotFormatter",
		// 	sut: func() hasFieldThatMatchesMatcher {
		// 		mock := mocks.NewMockMatcherGotFormatter(ctrl)
		// 		mock.EXPECT().String().Return("<submatcher.String()>")
		// 		mock.EXPECT().Got(gomock.Any()).Return("<submatcher.Got(...)>")
		// 		return HasField("MyField").ThatMatches(mock)
		// 	}(),
		// 	sampleValue:  gofakeit.SentenceSimple(),
		// 	expectedGot:  ".MyField <submatcher.Got(...)>",
		// 	expectedWant: ".MyField <submatcher.String()>",
		// },
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
