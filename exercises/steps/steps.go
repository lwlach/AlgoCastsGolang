package steps

import "fmt"

var printFunc = fmt.Println

func Steps(n int) {
	for i := 1; i <= n; i++ {
		var line string
		for j := 1; j <= i; j++ {
			line += "#"
		}
		for j := i + 1; j <= n; j++ {
			line += " "
		}
		printFunc(line)
	}
}
