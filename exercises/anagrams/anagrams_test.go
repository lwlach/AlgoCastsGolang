package anagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnagrams(t *testing.T) {
	t.Run("\"hello\" is an anagram of \"llohe\"", func(t *testing.T) {
		assert.True(t, Anagram("hello", "llohe"))
	})
	t.Run("\"Whoa! Hi!\" is an anagram of \"Hi! Whoa!\"", func(t *testing.T) {
		assert.True(t, Anagram("Whoa! Hi!", "Hi! Whoa!"))
	})
	t.Run("One One\" is not an anagram of \"Two two two\"", func(t *testing.T) {
		assert.False(t, Anagram("One One", "Two two two"))
	})
	t.Run("\"One one\" is not an anagram of \"One one c\"", func(t *testing.T) {
		assert.False(t, Anagram("One one", "One one c"))
	})
	t.Run("\"A tree, a life, a bench\" is not an anagram of \"A tree, a fence, a yard\"", func(t *testing.T) {
		assert.False(t, Anagram("A tree, a life, a bench", "A tree, a fence, a yard"))
	})
}
