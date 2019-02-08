package greedy

// 最佳解法
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	last := make([]int, 255)
	for i := range last {
		last[i] = -1
	}
	start := 0
	maxLen := 0
	for i, v := range s {
		if last[v] > start {
			start = last[v] + 1
		} else {
			curLen := i - start + 1
			if curLen > maxLen {
				maxLen = curLen
			}
		}
		last[v] = i
	}
	return maxLen
}

func lengthOfLongestSubstringMy(s string) int {
	if len(s) == 0 {
		return 0
	}
	mp := make(map[rune]int)
	curLen := 0
	maxLen := 0
	for i, v := range s {
		if _, ok := mp[v]; !ok {
			curLen++
			if curLen > maxLen {
				maxLen = curLen
			}
		} else {
			lastPos := mp[v]
			for k := range mp {
				if mp[k] <= lastPos {
					delete(mp, k)
				}
			}
			curLen = i - mp[v]
		}
		mp[v] = i
	}
	return maxLen
}
