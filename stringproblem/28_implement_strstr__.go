package stringproblem

func strStr(haystack string, needle string) int {
	if "" == needle {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if match(haystack[i:], needle) {
			return i
		}
	}
	return -1
}

func match(source, s string) bool {
	for i := 0; i < len(s); i++ {
		if source[i] != s[i] {
			return false
		}
	}
	return true
}
