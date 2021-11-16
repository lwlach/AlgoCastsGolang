package capitalize

import "strings"

func CapitalizeAlternative(s string) string {
	var words []string

	for _, word := range strings.Split(s, " ") {
		firstChar := string(word[0])
		words = append(words, strings.ToUpper(firstChar)+word[1:])
	}

	return strings.Join(words, " ")
}
