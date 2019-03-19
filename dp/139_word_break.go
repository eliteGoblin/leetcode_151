package dp

func wordBreak(s string, wordDict []string) bool {
	mpDict := make(map[string]bool)
	for _, v := range wordDict {
		mpDict[v] = true
	}
	cache := make(map[string]bool)
	var searchDFS func(s string) (res []string, isSuccess bool)
	searchDFS = func(s string) (res []string, isSuccess bool) {
		if s == "" {
			return nil, true
		}
		if flag, ok := cache[s]; ok {
			return nil, flag
		}
		cache[s] = false
		for i := 1; i < len(s); i++ {
			prefix := s[:i]
			if _, ok := mpDict[prefix]; ok {
				if res, flag := searchDFS(s[i:]); flag {
					cache[s] = true
					break
				}
			}
		}
		return nil, cache[s]
	}
	return searchDFS(s)
}

func wordBreakDP(s string, wordDict []string) bool {
	mpDict := make(map[string]bool)
	for _, v := range wordDict {
		mpDict[v] = true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			second := s[j:i]
			if f[j] {
				if _, ok := mpDict[second]; ok {
					f[i] = true
					break
				}
			}
		}
	}
	return f[len(s)]
}
