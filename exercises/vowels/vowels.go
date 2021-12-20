package vowels

import "strings"

func Vowels(s string) int {
	var cont int
	for _, c := range strings.ToLower(s) {
		contains := strings.ContainsAny(string(c), "aeiou")
		if contains {
			cont++
		}
	}
	return cont
}
