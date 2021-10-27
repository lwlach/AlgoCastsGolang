package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	t.Run("'aba' is a palindrome", func(t *testing.T) {
		assert.True(t, Palindrome("aba"))
	})

	t.Run("' aba' is not a palindrome", func(t *testing.T) {
		assert.False(t, Palindrome(" aba"))
	})

	t.Run("'aba ' is not a palindrome", func(t *testing.T) {
		assert.False(t, Palindrome("aba "))
	})

	t.Run("'greetings' is not a palindrome", func(t *testing.T) {
		assert.False(t, Palindrome("greetings"))
	})

	t.Run("'1000000001' is a palindrome", func(t *testing.T) {
		assert.True(t, Palindrome("1000000001"))
	})

	t.Run("'Fish hsif' is not a palindrome", func(t *testing.T) {
		assert.False(t, Palindrome("Fish hsif"))
	})

	t.Run("'pennep' is a palindrome", func(t *testing.T) {
		assert.True(t, Palindrome("pennep"))
	})
}
