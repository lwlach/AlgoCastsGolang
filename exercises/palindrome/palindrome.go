package palindrome

func Palindrome(s string) bool {
	for i, r := range s {
		if i == len(s)/2 {
			return true
		}
		reversedIndex := len(s) - i - 1
		if r != rune(s[reversedIndex]) {
			return false
		}
	}
	return true //never gets here
}
