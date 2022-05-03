package matchers

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/gomock-contrib/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	_ gomock.Matcher      = All()
	_ gomock.GotFormatter = All()
)

func TestAll_Matches(t *testing.T) {
	t.Run("When no submatchers, should match", func(t *testing.T) {
		sample := gofakeit.Address()
		sut := All()
		result := sut.Matches(sample)
		assert.True(t, result)
	})
	t.Run("Submatchers", func(t *testing.T) {
		testdata := []struct {
			name     string
			a        bool
			b        bool
			expected bool
		}{
			{
				name:     "false && false",
				a:        false,
				b:        false,
				expected: false,
			},
			{
				name:     "false && true",
				a:        false,
				b:        true,
				expected: false,
			},
			{
				name:     "true && false",
				a:        true,
				b:        false,
				expected: false,
			},
			{
				name:     "true && true",
				a:        true,
				b:        true,
				expected: true,
			},
		}
		for _, td := range testdata {
			t.Run(td.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				sample := gofakeit.Address()

				mock1 := mocks.NewMockMatcherGotFormatter(ctrl)
				mock1.EXPECT().Matches(sample).Return(td.a).MaxTimes(1)

				mock2 := mocks.NewMockMatcherGotFormatter(ctrl)
				mock2.EXPECT().Matches(sample).Return(td.b).MaxTimes(1)

				sut := All(mock1, mock2)

				result := sut.Matches(sample)
				assert.Equal(t, td.expected, result)
			})
		}
	})
}

func TestAll_WantMessage(t *testing.T) {
	t.Run("When no submatchers, return 'anything'", func(t *testing.T) {
		sut := All()
		assert.Equal(t, "anything", sut.String())
	})
	t.Run("Happy scenario", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock1 := mocks.NewMockMatcherGotFormatter(ctrl)
		mock1Want := gofakeit.SentenceSimple()
		mock1.EXPECT().String().Return(mock1Want)

		mock2 := mocks.NewMockMatcherGotFormatter(ctrl)
		mock2Want := gofakeit.SentenceSimple()
		mock2.EXPECT().String().Return(mock2Want)

		sut := All(mock1, mock2)

		result := sut.String()
		expected := fmt.Sprintf("%s;\n and %s", mock1Want, mock2Want)
		assert.Equal(t, expected, result)
	})
}

func TestAll_GotMessage(t *testing.T) {
	t.Run("When no submatchers, fallback to pretty print", func(t *testing.T) {
		sample := map[string]any{
			"key1": 1234,
		}
		sut := All()
		got := sut.Got(sample)
		assert.Equal(t, `map[string]any{
	"key1": 1234,
}`, got)
	})
	t.Run("When messages different, should list both", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sample := gofakeit.Address()

		mock1 := mocks.NewMockMatcherGotFormatter(ctrl)
		mock1Got := gofakeit.SentenceSimple()
		mock1.EXPECT().Got(sample).Return(mock1Got)

		mock2 := mocks.NewMockMatcherGotFormatter(ctrl)
		mock2Got := gofakeit.SentenceSimple()
		mock2.EXPECT().Got(sample).Return(mock2Got)

		sut := All(mock1, mock2)

		result := sut.Got(sample)
		expected := fmt.Sprintf("%s;\n%s", mock1Got, mock2Got)
		assert.Equal(t, expected, result)
	})

	t.Run("When messages equal, should deduplicate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sample := gofakeit.Address()
		got := gofakeit.SentenceSimple()

		mock1 := mocks.NewMockMatcherGotFormatter(ctrl)
		mock1.EXPECT().Got(sample).Return(got)

		mock2 := mocks.NewMockMatcherGotFormatter(ctrl)
		mock2.EXPECT().Got(sample).Return(got)

		sut := All(mock1, mock2)

		result := sut.Got(sample)
		expected := fmt.Sprintf("%s", got)
		assert.Equal(t, expected, result)
	})
}
