package matchers

import (
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
