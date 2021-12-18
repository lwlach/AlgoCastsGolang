package pyramid

import "fmt"

var printFunc = fmt.Println

func Pyramid(n int) {
	cols := (n * 2) - 1
	middle := cols / 2
	for i := 0; i < n; i++ {
		var line string
		for col := 0; col < cols; col++ {
			if col >= middle-i && col <= middle+i {
				line += "#"
			} else {
				line += " "
			}
		}
		printFunc(line)
	}
}
