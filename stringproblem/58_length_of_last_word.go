package stringproblem

func lengthOfLastWord(s string) int {
	if "" == s {
		return 0
	}
	resLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			resLen++
		} else if resLen > 0 {
			return resLen
		}
	}
	return resLen
}
