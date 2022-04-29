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
			name: "struct",
			sut: func() hasFieldThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<msg_want_struct>")
				mock.EXPECT().Got("the_value").Return("<msg_got_struct>")
				return HasField("ArrayField").ThatMatches(mock)
			}(),
			sampleValue:  struct{ ArrayField string }{ArrayField: "the_value"},
			expectedGot:  "field ArrayField <msg_got_struct>",
			expectedWant: "has field ArrayField that <msg_want_struct>",
		},
		{
			name: "map[string]string",
			sut: func() hasFieldThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<msg_want_map[string]string>")
				mock.EXPECT().Got("the_value_of_the_field").Return("<msg_got_map[string]string>")
				return HasField("KeyOfTheMap").ThatMatches(mock)
			}(),
			sampleValue:  map[string]string{"KeyOfTheMap": "the_value_of_the_field"},
			expectedGot:  "field KeyOfTheMap <msg_got_map[string]string>",
			expectedWant: "has field KeyOfTheMap that <msg_want_map[string]string>",
		},
		{
			name: "map[string]int",
			sut: func() hasFieldThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<msg_want_map[string]int>")
				mock.EXPECT().Got(999123).Return("<msg_got_map[string]int>")
				return HasField("KeyOfTheMap").ThatMatches(mock)
			}(),
			sampleValue:  map[string]int{"KeyOfTheMap": 999123},
			expectedGot:  "field KeyOfTheMap <msg_got_map[string]int>",
			expectedWant: "has field KeyOfTheMap that <msg_want_map[string]int>",
		},
		{
			name: "map[string]any",
			sut: func() hasFieldThatMatchesMatcher {
				mock := mocks.NewMockMatcherGotFormatter(ctrl)
				mock.EXPECT().String().Return("<msg_want_map[string]any>")
				mock.EXPECT().Got(123456789).Return("<msg_got_map[string]any>")
				return HasField("KeyOfTheMap").ThatMatches(mock)
			}(),
			sampleValue:  map[string]any{"KeyOfTheMap": 123456789},
			expectedGot:  "field KeyOfTheMap <msg_got_map[string]any>",
			expectedWant: "has field KeyOfTheMap that <msg_want_map[string]any>",
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
