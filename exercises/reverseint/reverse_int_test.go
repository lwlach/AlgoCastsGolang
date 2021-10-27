package reverseint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInt(t *testing.T) {
	t.Run("ReverseInt handles 0 as an input", func(t *testing.T) {
		assert.Equal(t, 0, ReverseInt(0))
	})

	t.Run("ReverseInt flips a positive number", func(t *testing.T) {
		assert.Equal(t, 5, ReverseInt(5))
		assert.Equal(t, 51, ReverseInt(15))
		assert.Equal(t, 9, ReverseInt(90))
		assert.Equal(t, 2359, ReverseInt(9532))
	})

	t.Run("ReverseInt flips a negative number", func(t *testing.T) {
		assert.Equal(t, -5, ReverseInt(-5))
		assert.Equal(t, -51, ReverseInt(-15))
		assert.Equal(t, -9, ReverseInt(-90))
		assert.Equal(t, -2359, ReverseInt(-9532))
	})
}
