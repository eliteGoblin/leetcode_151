package brute

func letterCombinations(digits string) []string {
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if digits == "" {
		return []string{}
	}
	res := make([]string, 0)
	var recur func(digitsLeft string, prefix string)
	recur = func(digitsLeft string, prefix string) {
		if len(digitsLeft) == 0 {
			res = append(res, prefix)
			return
		}
		set := mp[digitsLeft[0:1]]
		for _, v := range set {
			recur(digitsLeft[1:], prefix+string(v))
		}
	}
	recur(digits, "")
	return res
}
