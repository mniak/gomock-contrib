package matchers

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	_ gomock.Matcher      = LikeMap(map[string]any{})
	_ gomock.GotFormatter = LikeMap(map[string]any{})
)

type Stringable struct {
	Text string
}

func (s Stringable) String() string {
	return s.Text
}

func TestLikeMapMatcher(t *testing.T) {
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
				sut := LikeMap(td.expected)
				assert.False(t, sut.Matches(td.actual), "should not match")
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

				sut := LikeMap(map1)
				assert.True(t, sut.Matches(map2), "should match")
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
			{
				name: "stringable struct on expectation",
				expected: map[string]any{
					"TextField": Stringable{
						Text: faketext,
					},
				},
				actual: map[string]any{
					"TextField": faketext,
				},
			},
			{
				name: "stringable struct on actual value",
				expected: map[string]any{
					"TextField": faketext,
				},
				actual: map[string]any{
					"TextField": Stringable{
						Text: faketext,
					},
				},
			},
		}
		for _, td := range testdata {
			t.Run(td.name, func(t *testing.T) {
				sut := LikeMap(td.expected)
				assert.True(t, sut.Matches(td.actual), "should match")
			})
		}
	})

	t.Run("All number types should match each other", func(t *testing.T) {
		types := []struct {
			typename string
			convert  func(x float64) any
		}{
			// int
			{
				typename: "int",
				convert:  func(x float64) any { return int(x) },
			},
			{
				typename: "int8",
				convert:  func(x float64) any { return int8(x) },
			},
			{
				typename: "int16",
				convert:  func(x float64) any { return int16(x) },
			},
			{
				typename: "int32",
				convert:  func(x float64) any { return int32(x) },
			},
			{
				typename: "int64",
				convert:  func(x float64) any { return int64(x) },
			},
			// uint
			{
				typename: "uint",
				convert:  func(x float64) any { return uint(x) },
			},
			{
				typename: "uint8",
				convert:  func(x float64) any { return uint8(x) },
			},
			{
				typename: "uint16",
				convert:  func(x float64) any { return uint16(x) },
			},
			{
				typename: "uint32",
				convert:  func(x float64) any { return uint32(x) },
			},
			{
				typename: "uint64",
				convert:  func(x float64) any { return uint64(x) },
			},
		}

		for _, type1 := range types {
			for _, type2 := range types {
				t.Run(fmt.Sprintf("%s==%s", type1.typename, type2.typename), func(t *testing.T) {
					number := float64(gofakeit.IntRange(50, 100))

					map1 := map[string]any{
						"key": type1.convert(number),
					}

					map2 := map[string]any{
						"key": type2.convert(number),
					}

					sut := LikeMap(map1)
					assert.True(t, sut.Matches(map2), "should match")
				})
			}
		}
	})
}

func TestLikeMapMatcher_WantString(t *testing.T) {
	testdata := []struct {
		name            string
		expectedMap     map[string]any
		expectedMessage string
	}{
		{
			name:            "empty map",
			expectedMap:     map[string]any{},
			expectedMessage: "matches map[string]any{}",
		},
		{
			name: "basic map with 2 fields",
			expectedMap: map[string]any{
				"field1": "value1",
				"field2": 2,
			},
			expectedMessage: `matches map[string]any{
	"field1": "value1",
	"field2": 2,
}`,
		},
		{
			name: "map with submap",
			expectedMap: map[string]any{
				"submap": map[string]any{
					"field1": "value1",
					"field2": 2,
				},
			},
			expectedMessage: `matches map[string]any{
	"submap": map[string]any{
		"field1": "value1",
		"field2": 2,
	},
}`,
		},
		{
			name: "map with slice",
			expectedMap: map[string]any{
				"slice": []string{
					"slice item 1",
					"slice item 2",
				},
			},
			expectedMessage: `matches map[string]any{
	"slice": []string{
		"slice item 1",
		"slice item 2",
	},
}`,
		},
		{
			name: "stringable struct on expecation side",
			expectedMap: map[string]any{
				"stringable": Stringable{
					Text: "inner message",
				},
			},
			expectedMessage: `matches map[string]any{
	"stringable": "inner message",
}`,
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			sut := LikeMap(td.expectedMap)
			assert.Equal(t, td.expectedMessage, sut.String())
		})
	}
}

func TestLikeMapMatcher_GotString(t *testing.T) {
	testdata := []struct {
		name            string
		data            any
		expectedMessage string
	}{
		{
			name:            "nil",
			data:            nil,
			expectedMessage: "is nil",
		},
		{
			name:            "empty map",
			data:            map[string]any{},
			expectedMessage: "is map[string]any{}",
		},
		{
			name: "empty map",
			data: map[string]any{
				"stringable": Stringable{
					Text: "hello world",
				},
			},
			expectedMessage: `is map[string]any{
	"stringable": "hello world",
}`,
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			expectedMap := map[string]any{gofakeit.UUID(): gofakeit.SentenceSimple()}
			sut := LikeMap(expectedMap)
			assert.Equal(t, td.expectedMessage, sut.Got(td.data))
		})
	}
}

func TestLikeMap_AcceptMatchersInFields(t *testing.T) {
	testdata := []bool{true, false}
	for _, b := range testdata {
		t.Run(fmt.Sprint(b), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fakeValue := gofakeit.SentenceSimple()
			fakeSubMatcherWant := gofakeit.SentenceSimple()

			sample := map[string]any{
				"Field": fakeValue,
			}

			mock := mocks.NewMockMatcherGotFormatter(ctrl)
			mock.EXPECT().Matches(fakeValue).Return(b)
			mock.EXPECT().String().Return(fakeSubMatcherWant)

			sut := LikeMap(map[string]any{
				"Field": mock,
			})
			assert.Equal(t, b, sut.Matches(sample))

			expectedWant := fmt.Sprintf(`matches map[string]any{
	"Field": "%s",
}`, fakeSubMatcherWant)
			assert.Equal(t, expectedWant, sut.String())

			expectedGot := fmt.Sprintf(`map[string]any{
	"Field": "%s",
}`, fakeValue)

			assert.Equal(t, expectedGot, sut.Got(sample))
		})
	}
}
