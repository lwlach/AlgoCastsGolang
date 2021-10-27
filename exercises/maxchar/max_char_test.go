package maxchar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxChar(t *testing.T) {
	t.Run("Finds the most frequently used char", func(t *testing.T) {
		assert.Equal(t, 'a', MaxChar("abcdefghijklmnaaaaa"))
	})
	t.Run("Works with numbers in the string", func(t *testing.T) {
		assert.Equal(t, '1', MaxChar("ab1c1d1e1f1g1"))
	})
}
