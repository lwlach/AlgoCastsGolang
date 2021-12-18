package pyramid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPyramid(t *testing.T) {
	setupTest := func() {
		output = make(map[int]interface{})
		index = 0
	}

	t.Run("prints a pyramid for n = 2", func(t *testing.T) {
		setupTest()

		Pyramid(2)
		assert.Len(t, output, 2)
		assert.Equal(t, " # ", output[0])
		assert.Equal(t, "###", output[1])
	})

	t.Run("steps called with n = 3", func(t *testing.T) {
		setupTest()

		Pyramid(3)
		assert.Len(t, output, 3)
		assert.Equal(t, "  #  ", output[0])
		assert.Equal(t, " ### ", output[1])
		assert.Equal(t, "#####", output[2])
	})

	t.Run("steps called with n = 4", func(t *testing.T) {
		setupTest()

		Pyramid(4)
		assert.Len(t, output, 4)
		assert.Equal(t, "   #   ", output[0])
		assert.Equal(t, "  ###  ", output[1])
		assert.Equal(t, " ##### ", output[2])
		assert.Equal(t, "#######", output[3])
	})
}

var output map[int]interface{}
var index int

func addOutput(a ...interface{}) (n int, err error) {
	for _, v := range a {
		output[index] = v
		index++
	}
	return len(a), nil
}

func init() {
	printFunc = addOutput
}
