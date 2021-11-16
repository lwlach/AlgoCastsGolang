package capitalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalizeAlternative(t *testing.T) {
	t.Run("capitalizes the first letter of every word in a sentence", func(t *testing.T) {
		actual := CapitalizeAlternative("hi there, how is it going?")
		assert.Equal(t, "Hi There, How Is It Going?", actual)
	})

	t.Run("capitalizes the first letter", func(t *testing.T) {
		actual := CapitalizeAlternative("i love breakfast at bill miller bbq")
		assert.Equal(t, "I Love Breakfast At Bill Miller Bbq", actual)
	})
}
