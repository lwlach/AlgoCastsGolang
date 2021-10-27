package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	setupTest := func() {
		output = make(map[int]interface{})
		index = 0
	}

	t.Run("Calling fizzbuzz with `5` prints out 5 statements", func(t *testing.T) {
		setupTest()
		FizzBuzz(5)

		assert.Equal(t, 5, len(output))
	})

	t.Run("Calling fizzbuzz with 15 prints out the correct values", func(t *testing.T) {
		setupTest()
		FizzBuzz(15)

		assert.Equal(t, 1, output[0])
		assert.Equal(t, 2, output[1])
		assert.Equal(t, "Fizz", output[2])
		assert.Equal(t, 4, output[3])
		assert.Equal(t, "Buzz", output[4])
		assert.Equal(t, "Fizz", output[5])
		assert.Equal(t, 7, output[6])
		assert.Equal(t, 8, output[7])
		assert.Equal(t, "Fizz", output[8])
		assert.Equal(t, "Buzz", output[9])
		assert.Equal(t, 11, output[10])
		assert.Equal(t, "Fizz", output[11])
		assert.Equal(t, 13, output[12])
		assert.Equal(t, 14, output[13])
		assert.Equal(t, "FizzBuzz", output[14])
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
