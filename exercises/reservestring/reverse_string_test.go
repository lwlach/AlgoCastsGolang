package reservestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Run("Reverse reverses a string", func(t *testing.T) {
		assert.Equal(t, "dcba", ReverseString("abcd"))
	})
	t.Run("Reverse reverses a string", func(t *testing.T) {
		assert.Equal(t, "dcba  ", ReverseString("  abcd"))
	})
}
