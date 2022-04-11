package matchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	_ gomock.Matcher = &typedMatcher[string]{}
	_ gomock.Matcher = &typedMatcher[int32]{}
)

func TestTyped(t *testing.T) {
	t.Run("simply test type", func(t *testing.T) {
		t.Run("type=String", func(t *testing.T) {
			sut := Typed[string]()

			assert.False(t, sut.Matches(gofakeit.Int32()), "should not match int")
			assert.False(t, sut.Matches(gofakeit.Bool()), "should not match bool")
			assert.True(t, sut.Matches(gofakeit.SentenceSimple()), "should match string")
		})

		t.Run("type=Int", func(t *testing.T) {
			sut := Typed[int]()

			assert.False(t, sut.Matches(gofakeit.SentenceSimple()), "should match string")
			assert.False(t, sut.Matches(gofakeit.Bool()), "should not match bool")
			assert.False(t, sut.Matches(gofakeit.Int32()), "should not match int32")
			assert.True(t, sut.Matches(int(gofakeit.Int32())), "should match int")
		})
	})
}
