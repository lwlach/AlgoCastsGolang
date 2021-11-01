package anagrams

import (
	"regexp"
	"strings"
)

func Anagram(s1, s2 string) bool {
	s1 = normalize(s1)
	s2 = normalize(s2)
	if len(s1) != len(s2) {
		return false
	}
	sMap1 := makeMap(s1)
	sMap2 := makeMap(s2)
	for r, i := range sMap1 {
		if i != sMap2[r] {
			return false
		}
	}
	return true
}

func normalize(s string) string {
	rgx := regexp.MustCompile("[^\\w]")
	return strings.ToLower(rgx.ReplaceAllString(s, ""))
}

func makeMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}
