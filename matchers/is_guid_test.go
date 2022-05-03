package matchers

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestIsGUID(t *testing.T) {
	sut := IsGUID()

	assert.False(t, sut.Matches(gofakeit.SentenceSimple()))
	assert.True(t, sut.Matches(gofakeit.UUID()))
}
