package array

func LongestConsecutive(nums []int) int {
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = false
	}
	longest := 0
	for _, i := range nums {
		if mp[i] == true {
			continue
		}
		mp[i] = true
		length := 1
		for j := i + 1; ; j++ {
			if _, ok := mp[j]; ok {
				mp[j] = true
				length++
			} else {
				break
			}
		}
		for j := i - 1; ; j-- {
			if _, ok := mp[j]; ok {
				mp[j] = true
				length++
			} else {
				break
			}
		}
		if length > longest {
			longest = length
		}
	}
	return longest
}
