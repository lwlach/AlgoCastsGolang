package maxchar

func MaxChar(s string) rune {
	chars := make(map[rune]int)
	for _, c := range s {
		chars[c]++
	}
	var bigger int
	var result rune
	for char, count := range chars {
		if count > bigger {
			bigger = count
			result = char
		}
	}
	return result
}
