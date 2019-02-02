package array

import (
	"sort"
)

func FourSumUsingHash(nums []int, target int) [][]int {
	sort.Ints(nums)
	mp := make(map[[4]int]bool)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			start := j + 1
			end := len(nums) - 1
			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					mp[[4]int{nums[i], nums[j], nums[start], nums[end]}] = true
					start++
					end--
				} else if sum < target {
					start++
				} else {
					end--
				}
			}
		}
	}
	res := make([][]int, 0)
	for k := range mp {
		quadruplet := k
		res = append(res, quadruplet[:])
	}
	return res
}

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		for i > 0 && i < len(nums)-3 && nums[i] == nums[i-1] {
			i++
		}
		for j := i + 1; j < len(nums)-2; j++ {
			for j > i+1 && j < len(nums)-2 && nums[j] == nums[j-1] {
				j++
			}
			start := j + 1
			end := len(nums) - 1
			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					for start < end && nums[start] == nums[start-1] {
						start++
					}
				} else if sum < target {
					start++
					for start < end && nums[start] == nums[start-1] {
						start++
					}
				} else {
					end--
					for start < end && nums[end] == nums[end+1] {
						end--
					}
				}
			}
		}
	}
	return res
}
