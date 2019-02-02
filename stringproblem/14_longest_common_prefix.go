package stringproblem

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := ""
	pos := 0
	for {
		var char byte
		for _, v := range strs {
			if pos >= len(v) {
				return prefix
			}
			if 0 == char {
				char = v[pos]
			} else {
				if char != v[pos] {
					return prefix
				}
			}
		}
		prefix += string(char)
		pos++
	}
}
