package stringproblem

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	if len(s) <= 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; {
		if !isAlphaNum(s, i) {
			i++
		} else if !isAlphaNum(s, j) {
			j--
		} else if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func isAlphaNum(s string, n int) bool {
	return (s[n] >= '0' && s[n] <= '9') ||
		(s[n] >= 'a' && s[n] <= 'z')
}
