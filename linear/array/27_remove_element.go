package array

import (
	"sort"
)

func RemoveElement(nums []int, val int) int {
	sort.Ints(nums)
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
