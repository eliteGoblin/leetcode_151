package stringproblem

import "fmt"

func countAndSay(n int) string {
	if 1 == n {
		return "1"
	}
	s := "1"
	for i := 1; i < n; i++ {
		s = readStr(s)
	}
	return s
}

func readStr(s string) string {
	if "" == s {
		return ""
	}
	digit, count := findTilNotSame(s)
	return fmt.Sprintf("%d", count) + string(digit) + readStr(s[count:])
}

func findTilNotSame(s string) (digit byte, count int) {
	digit = s[0]
	var i int
	for i = 1; i < len(s); i++ {
		if digit != s[i] {
			return digit, i
		}
	}
	return digit, i
}
