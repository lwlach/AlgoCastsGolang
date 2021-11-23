package steps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSteps(t *testing.T) {
	setupTest := func() {
		output = make(map[int]interface{})
		index = 0
	}

	t.Run("steps called with n = 1", func(t *testing.T) {
		setupTest()

		Steps(1)
		assert.Len(t, output, 1)
		assert.Equal(t, "#", output[0])
	})

	t.Run("steps called with n = 2", func(t *testing.T) {
		setupTest()

		Steps(2)
		assert.Len(t, output, 2)
		assert.Equal(t, "# ", output[0])
		assert.Equal(t, "##", output[1])
	})

	t.Run("steps called with n = 3", func(t *testing.T) {
		setupTest()

		Steps(3)
		assert.Len(t, output, 3)
		assert.Equal(t, "#  ", output[0])
		assert.Equal(t, "## ", output[1])
		assert.Equal(t, "###", output[2])
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
