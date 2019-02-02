package array

import (
	"sort"
)

func NextPermutation(nums []int) {
	i := findFirstDecreaseFromBack(nums)
	if -1 == i {
		sort.Ints(nums)
		return
	}
	j := findFirstLargerThanFromBack(nums, nums[i])
	nums[i], nums[j] = nums[j], nums[i]
	right := nums[i+1:]
	sort.Ints(right)
	nums = append(nums[:i+1], right...)
}

func findFirstDecreaseFromBack(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			return i
		}
	}
	return -1
}

func findFirstLargerThanFromBack(nums []int, target int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > target {
			return i
		}
	}
	return -1
}
