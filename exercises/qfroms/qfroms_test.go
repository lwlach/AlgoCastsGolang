package qfroms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueFromStacks(t *testing.T) {
	t.Run("order of elements is maintained", func(t *testing.T) {
		var q Queue
		q.Add(1)
		q.Add(2)
		q.Add(3)
		assert.Equal(t, 1, q.Remove())
		assert.Equal(t, 2, q.Remove())
		assert.Equal(t, 3, q.Remove())
		assert.Nil(t, q.Remove())
	})

	t.Run("peek returns, but does not remove, the first value", func(t *testing.T) {
		var q Queue
		q.Add(1)
		q.Add(2)
		assert.Equal(t, 1, q.Peek())
		assert.Equal(t, 1, q.Peek())
		assert.Equal(t, 1, q.Remove())
		assert.Equal(t, 2, q.Remove())
	})
}
