package utils

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

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
				assert.False(t, MatchMaps(td.expected, td.actual), "should not match")
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
			// float
			{
				typename: "float64",
				convert:  func(x float64) any { return x },
			},
			{
				typename: "float32",
				convert:  func(x float64) any { return float32(x) },
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

					assert.True(t, MatchMaps(map1, map2), "should match")
				})
			}
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

				assert.True(t, MatchMaps(map1, map2), "should match")
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
				assert.True(t, MatchMaps(td.expected, td.actual), "should match")
			})
		}
	})
}
