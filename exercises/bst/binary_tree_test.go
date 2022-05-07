package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	t.Run("Node can insert correctly", func(t *testing.T) {
		node := NewNode(10)
		node.Insert(5)
		node.Insert(15)
		node.Insert(17)

		assert.Equal(t, 5, node.left.data)
		assert.Equal(t, 15, node.right.data)
		assert.Equal(t, 17, node.right.right.data)
	})

	t.Run("Contains returns node with the same data", func(t *testing.T) {
		node := NewNode(10)
		node.Insert(5)
		node.Insert(15)
		node.Insert(20)
		node.Insert(0)
		node.Insert(-5)
		node.Insert(3)

		three := node.left.left.right
		assert.Equal(t, three, node.Contains(3))
	})

	t.Run("Contains returns null if value not found", func(t *testing.T) {
		node := NewNode(10)
		node.Insert(5)
		node.Insert(15)
		node.Insert(20)
		node.Insert(0)
		node.Insert(-5)
		node.Insert(3)

		assert.Nil(t, node.Contains(999))
	})

	t.Run("Validate recognizes a valid BST", func(t *testing.T) {
		node := NewNode(10)
		node.Insert(5)
		node.Insert(15)
		node.Insert(0)
		node.Insert(20)

		assert.True(t, Validate(node, nil, nil))
	})

	t.Run("Validate recognizes an invalid BST", func(t *testing.T) {
		node := NewNode(10)
		node.Insert(5)
		node.Insert(15)
		node.Insert(0)
		node.Insert(20)
		node.left.left.right = NewNode(999)

		assert.False(t, Validate(node, nil, nil))
	})
}
