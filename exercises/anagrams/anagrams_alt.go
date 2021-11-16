package anagrams

import (
	"sort"
	"strings"
)

func AnagramAlternative(s1, s2 string) bool {
	s1 = normalize(s1)
	s2 = normalize(s2)

	a := strings.Split(s1, "")
	sort.Strings(a)
	b := strings.Split(s2, "")
	sort.Strings(b)

	return strings.Join(a, "") == strings.Join(b, "")
}
