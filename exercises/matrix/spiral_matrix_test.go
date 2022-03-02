package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralMatrix(t *testing.T) {

	t.Run("2x2 matrix", func(t *testing.T) {
		result := SpiralMatrix(2)

		assert.Len(t, result, 2)
		assert.Equal(t, []int{1, 2}, result[0])
		assert.Equal(t, []int{4, 3}, result[1])
	})

	t.Run("3x3 matrix", func(t *testing.T) {
		result := SpiralMatrix(3)

		assert.Len(t, result, 3)
		assert.Equal(t, []int{1, 2, 3}, result[0])
		assert.Equal(t, []int{8, 9, 4}, result[1])
		assert.Equal(t, []int{7, 6, 5}, result[2])
	})

	t.Run("4x4 matrix", func(t *testing.T) {
		result := SpiralMatrix(4)

		assert.Len(t, result, 4)
		assert.Equal(t, []int{1, 2, 3, 4}, result[0])
		assert.Equal(t, []int{12, 13, 14, 5}, result[1])
		assert.Equal(t, []int{11, 16, 15, 6}, result[2])
		assert.Equal(t, []int{10, 9, 8, 7}, result[3])
	})
}
