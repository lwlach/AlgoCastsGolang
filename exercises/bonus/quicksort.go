package main

import (
	"fmt"
)

func main() {
	var arr = []int{-1, 0, 1, 4, 3, 5, 6, 8, 9, 11, 4, 11, -3, 2, 6, -2, 10, 5, 8, 0, -5, 3, 9, -4, 1, 7}
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}
	var (
		left, pivot = 0, len(arr) - 1
	)
	for i := range arr {
		if arr[i] < arr[pivot] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[pivot] = arr[pivot], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])
}
