package stringproblem

func IsMatch(s string, p string) bool {
	if p == "" {
		return "" == s
	}
	curS, curP := 0, 0
	if len(p) == 1 || p[1] != '*' {
		if "" == s {
			return false
		}
		if equalChar(p[curP], s[curS]) {
			return IsMatch(s[curS+1:], p[curP+1:])
		}
		return false
	} else {
		for curS < len(s) && equalChar(p[curP], s[curS]) {
			if IsMatch(s[curS:], p[curP+2:]) {
				return true
			}
			curS++
		}
		return IsMatch(s[curS:], p[curP+2:])
	}
}

func equalChar(a, b byte) bool {
	if '.' == a || '.' == b {
		return true
	}
	return a == b
}
