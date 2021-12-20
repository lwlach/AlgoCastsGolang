package vowels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVowels(t *testing.T) {
	t.Run("returns the number of vowels used", func(t *testing.T) {
		assert.Equal(t, Vowels("aeiou"), 5)
	})
	t.Run("returns the number of vowels used when they are capitalized", func(t *testing.T) {
		assert.Equal(t, Vowels("AEIOU"), 5)
	})
	t.Run("returns the number of vowels used", func(t *testing.T) {
		assert.Equal(t, Vowels("abcdefghijklmnopqrstuvwxyz"), 5)
	})
	t.Run("returns the number of vowels used", func(t *testing.T) {
		assert.Zero(t, Vowels("bcdfghjkl"))
	})
}
