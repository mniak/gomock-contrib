package matchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestIsValidGUID(t *testing.T) {
	sut := IsValidGUID()

	assert.False(t, sut.Matches(gofakeit.SentenceSimple()))
	assert.True(t, sut.Matches(gofakeit.UUID()))
}
