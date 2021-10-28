package chunk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunk(t *testing.T) {
	t.Run("chunk divides an array of 10 elements with chunk size 2", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		expected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
		assert.Equal(t, expected, Chunk(arr, 2))
	})

	t.Run("chunk divides an array of 3 elements with chunk size 1", func(t *testing.T) {
		arr := []int{1, 2, 3}

		expected := [][]int{{1}, {2}, {3}}
		assert.Equal(t, expected, Chunk(arr, 1))
	})

	t.Run("chunk divides an array of 5 elements with chunk size 3", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}

		expected := [][]int{{1, 2, 3}, {4, 5}}
		assert.Equal(t, expected, Chunk(arr, 3))
	})

	t.Run("chunk divides an array of 13 elements with chunk size 5", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

		expected := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}}
		assert.Equal(t, expected, Chunk(arr, 5))
	})
}
