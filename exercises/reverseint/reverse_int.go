package reverseint

import (
	"strconv"

	"github.com/lwlach/AlgoCastsGolang/exercises/reservestring"
)

func ReverseInt(n int) int {
	reversedStr := reservestring.ReverseString(strconv.Itoa(n))
	lastIndex := len(reversedStr) - 1
	if reversedStr[lastIndex] == '-' {
		reversedStr = "-" + reversedStr[:lastIndex]
	}
	reversedInt, _ := strconv.Atoi(reversedStr)
	return reversedInt
}
