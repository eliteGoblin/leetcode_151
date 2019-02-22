package array

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		first := i + 1
		last := len(nums) - 1
		for first < last {
			sum := nums[first] + nums[last] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[first], nums[last]})
				first++
				last--
				for first < last && nums[first] == nums[first-1] {
					first++
				}
			} else if sum < 0 {
				first++
			} else {
				last--
			}
		}
	}
	return res
}
