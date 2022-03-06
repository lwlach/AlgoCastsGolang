package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("stack can follows first in, last out", func(t *testing.T) {
		s := Stack{}
		s.Push(1)
		s.Push(2)
		s.Push(3)

		assert.Equal(t, 3, s.Pop())
		assert.Equal(t, 2, s.Pop())
		assert.Equal(t, 1, s.Pop())
		assert.Nil(t, s.Pop())
	})

	t.Run("peek returns the first element but doesnt pop it", func(t *testing.T) {
		s := Stack{}
		s.Push(1)
		s.Push(2)
		s.Push(3)

		assert.Equal(t, 3, s.Peek())
		assert.Equal(t, 3, s.Pop())
		assert.Equal(t, 2, s.Peek())
		assert.Equal(t, 2, s.Pop())
		assert.Equal(t, 1, s.Peek())
		assert.Equal(t, 1, s.Pop())
		assert.Nil(t, s.Peek())
	})
}
