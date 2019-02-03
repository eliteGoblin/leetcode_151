package greedy

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	steps := 0
	for cur := 0; cur < len(nums); {
		cur = getNextLocation(nums, cur)
		steps++
		if cur >= len(nums) {
			break
		}
	}
	return steps
}

func getNextLocation(nums []int, cur int) int {
	max := -1
	maxLocation := -1
	for i := cur + 1; i <= cur+nums[cur]; i++ {
		if nums[i]+i-cur > max {
			max = nums[i] + i - cur
			maxLocation = i
		}
	}
	return maxLocation
}
