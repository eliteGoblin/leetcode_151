package dp

func wordBreak(s string, wordDict []string) bool {
	mpDict := make(map[string]bool)
	for _, v := range wordDict {
		mpDict[v] = true
	}
	cache := make(map[string]bool)
	var searchDFS func(s string) bool
	searchDFS = func(s string) bool {
		if s == "" {
			return true
		}
		if flag, ok := cache[s]; ok {
			return flag
		}
		cache[s] = false
		for i := 1; i < len(s); i++ {
			prefix := s[:i]
			if _, ok := mpDict[prefix]; ok {
				if searchDFS(s[i:]) {
					cache[s] = true
					break
				}
			}
		}
		return cache[s]
	}
	return searchDFS(s)
}
