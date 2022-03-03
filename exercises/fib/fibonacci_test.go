package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Run("calculates correct fib value for 1", func(t *testing.T) {
		assert.Equal(t, 1, Fibonacci(1))
	})
	t.Run("calculates correct fib value for 2", func(t *testing.T) {
		assert.Equal(t, 1, Fibonacci(2))
	})
	t.Run("calculates correct fib value for 3", func(t *testing.T) {
		assert.Equal(t, 2, Fibonacci(3))
	})
	t.Run("calculates correct fib value for 4", func(t *testing.T) {
		assert.Equal(t, 3, Fibonacci(4))
	})
	t.Run("calculates correct fib value for 39", func(t *testing.T) {
		assert.Equal(t, 63245986, Fibonacci(39))
	})
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(100)
	}
}
