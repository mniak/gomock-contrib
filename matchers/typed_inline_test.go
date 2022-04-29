package matchers

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/brianvoe/gofakeit/v6"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	_ gomock.Matcher = typedMatcher[string]{}
// 	_ gomock.Matcher = typedMatcher[int32]{}
// 	_ gomock.Matcher = TypedInline[string]("", nil)
// 	_ gomock.Matcher = TypedInline[int32]("", nil)
// )

// func TestInlineTypedInline(t *testing.T) {
// 	for _, b := range []bool{true, false} {
// 		t.Run(fmt.Sprintf("when inline fn returns %t", b), func(t *testing.T) {
// 			t.Run("simply test type", func(t *testing.T) {
// 				t.Run("type=String", func(t *testing.T) {
// 					expectedValue := gofakeit.SentenceSimple()
// 					sut := TypedInline("message", func(x string) bool {
// 						return b && x == expectedValue
// 					})

// 					assert.False(t, sut.Matches(gofakeit.Int32()), "should not match int")
// 					assert.False(t, sut.Matches(gofakeit.Bool()), "should not match bool")
// 					assert.False(t, sut.Matches(gofakeit.SentenceSimple()), "should not match random string")

// 					assert.Equal(t, b, sut.Matches(expectedValue), "should match exact value")
// 				})

// 				t.Run("type=Int", func(t *testing.T) {
// 					expectedValue := int(gofakeit.Int32())
// 					sut := TypedInline("message", func(x int) bool {
// 						return b && x == expectedValue
// 					})

// 					assert.False(t, sut.Matches(gofakeit.SentenceSimple()), "should match string")
// 					assert.False(t, sut.Matches(gofakeit.Bool()), "should not match bool")
// 					assert.False(t, sut.Matches(gofakeit.Int32()), "should not match int32")
// 					assert.False(t, sut.Matches(int(gofakeit.Int32())), "should not match random int")

// 					assert.Equal(t, b, sut.Matches(expectedValue), "should match exact value")
// 				})
// 			})
// 		})
// 	}
// }
