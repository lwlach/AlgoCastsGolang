package weave

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeave(t *testing.T) {
	t.Run("peek returns, but does not remove, the first value", func(t *testing.T) {
		q := Queue{}
		q.Add(1)
		q.Add(2)

		assert.Equal(t, 1, q.Peek())
		assert.Equal(t, 1, q.Peek())
		assert.Equal(t, 1, q.Remove())
		assert.Equal(t, 2, q.Remove())
	})

	t.Run("weave can combine two queues", func(t *testing.T) {
		q1 := Queue{}
		q1.Add(1)
		q1.Add(2)
		q1.Add(3)
		q1.Add(4)
		q2 := Queue{}
		q2.Add("one")
		q2.Add("two")
		q2.Add("three")
		q2.Add("four")

		q3 := Weave(q1, q2)
		assert.Equal(t, 1, q3.Remove())
		assert.Equal(t, "one", q3.Remove())
		assert.Equal(t, 2, q3.Remove())
		assert.Equal(t, "two", q3.Remove())
		assert.Equal(t, 3, q3.Remove())
		assert.Equal(t, "three", q3.Remove())
		assert.Equal(t, 4, q3.Remove())
		assert.Equal(t, "four", q3.Remove())
		assert.Nil(t, q3.Remove())
	})
}
