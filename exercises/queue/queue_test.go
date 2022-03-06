package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("order of elements ir maintained", func(t *testing.T) {
		q := Queue{}
		q.Add(1)
		q.Add(2)
		q.Add(3)

		assert.Equal(t, 1, q.Remove())
		assert.Equal(t, 2, q.Remove())
		assert.Equal(t, 3, q.Remove())
		assert.Nil(t, q.Remove())
	})
}
