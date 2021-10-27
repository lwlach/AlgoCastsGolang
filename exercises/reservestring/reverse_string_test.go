package reservestring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Run("Reverse reverses a string", func(t *testing.T) {
		assert.Equal(t, "dcba", ReverseString("abcd"))
	})
	t.Run("Reverse reverses a string", func(t *testing.T) {
		assert.Equal(t, "dcba  ", ReverseString("  abcd"))
	})
}
