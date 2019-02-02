package stringproblem

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range strs {
		s := sortString(v)
		if _, ok := mp[s]; ok {
			mp[s] = append(mp[s], v)
		} else {
			mp[s] = []string{v}
		}
	}
	res := make([][]string, 0)
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Sort(sortRunes(runes))
	return string(runes)
}
