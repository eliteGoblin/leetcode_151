package stringproblem

func IsMatchWildcard(s string, p string) bool {
	if "" == p {
		return "" == s
	}
	if p[0] != '*' {
		if "" != s {
			if wildcardEqual(s[0], p[0]) {
				return IsMatchWildcard(s[1:], p[1:])
			}
			return false
		}
		return false
	}
	for curS := 0; curS <= len(s); curS++ {
		if IsMatchWildcard(s[curS:], p[1:]) {
			return true
		}
	}
	return false
}

func wildcardEqual(a, b byte) bool {
	if a == '?' || b == '?' {
		return true
	}
	return a == b
}
