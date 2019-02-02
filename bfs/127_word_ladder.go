package bfs

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	setOfNextRound := make([]string, 1)
	setOfNextRound[0] = beginWord
	isSearchBefore := make(map[string]bool)
	backup := make([]string, 0)
	step := 0
	for {
		step++
		isAllSearchedBefore := true
		for _, v := range setOfNextRound {
			if _, ok := isSearchBefore[v]; !ok {
				if v == endWord {
					return step
				}
				isSearchBefore[v] = ok
				backup = append(backup, get1StepSet(v, dict)...)
				isAllSearchedBefore = false
			}
		}
		if isAllSearchedBefore {
			return 0
		}
		setOfNextRound = backup
		backup = make([]string, 0)
	}
}

func get1StepSet(src string, isExist map[string]bool) []string {
	res := make([]string, 0)
	for i := 0; i < len(src); i++ {
		for ch := 'a'; ch <= 'z'; ch++ {
			tmp := src[0:i] + string(ch) + src[i+1:]
			if _, ok := isExist[tmp]; ok {
				res = append(res, tmp)
			}
		}
	}
	return res
}
