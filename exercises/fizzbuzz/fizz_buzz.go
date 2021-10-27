package fizzbuzz

import "fmt"

var printFunc = fmt.Println

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		isMultipleOf3 := i%3 == 0
		isMultipleOf5 := i%5 == 0
		if isMultipleOf3 && isMultipleOf5 {
			printFunc("FizzBuzz")
		} else if isMultipleOf3 {
			printFunc("Fizz")
		} else if isMultipleOf5 {
			printFunc("Buzz")
		} else {
			printFunc(i)
		}
	}
}
