package array

import (
	"sort"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	minDiff := MaxInt
	minSum := MaxInt
	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if target == sum {
				return target
			}
			if absInt(sum, target) < minDiff {
				minDiff = absInt(sum, target)
				minSum = sum
			}
			if target < sum {
				end--
			} else {
				start++
			}
		}
	}
	return minSum
}

func absInt(a, b int) int {
	if b-a > 0 {
		return b - a
	}
	return a - b
}
